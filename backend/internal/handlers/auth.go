package handlers

import (
	"webifylab-backend/internal/dto/request"
	"webifylab-backend/internal/services"
	"webifylab-backend/pkg/logger"
	"webifylab-backend/pkg/response"
	"webifylab-backend/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login menangani POST /api/v1/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	// 1. Bind request body ke DTO
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", []response.Error{
			{Message: "Request body must be valid JSON"},
		})
		return
	}

	// 2. Validasi request
	if errors := validator.Validate(req); len(errors) > 0 {
		response.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	// 3. Panggil service
	loginResp, err := h.authService.Login(&req)
	if err != nil {
		// Handle error berdasarkan tipe
		switch err {
		case services.ErrInvalidCredentials:
			// SECURITY: Jangan expose detail error (email salah atau password salah)
			response.Unauthorized(c, "Invalid email or password")
			return
		default:
			logger.Error().Err(err).Msg("❌ Login failed")
			response.InternalServerError(c, "An error occurred during login")
			return
		}
	}

	// 4. Return success response
	response.Success(c, "Login successful", loginResp)
}

// RefreshToken menangani POST /api/v1/auth/refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req request.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", []response.Error{{Message: "Request body must be valid JSON"}})
		return
	}

	if errors := validator.Validate(req); len(errors) > 0 {
		response.UnprocessableEntity(c, "Validation failed", errors)
		return
	}

	refreshResp, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		switch err {
		case services.ErrInvalidRefreshToken, services.ErrRefreshTokenExpired:
			response.Unauthorized(c, "Invalid or expired refresh token")
			return
		default:
			logger.Error().Err(err).Msg("❌ Refresh token failed")
			response.InternalServerError(c, "An error occurred during token refresh")
			return
		}
	}

	response.Success(c, "Token refreshed successfully", refreshResp)
}

// Logout menangani POST /api/v1/auth/logout
func (h *AuthHandler) Logout(c *gin.Context) {
	var req request.RefreshTokenRequest // Kita pakai DTO yang sama untuk ambil refresh_token
	_ = c.ShouldBindJSON(&req)          // Ignore error, karena refresh_token opsional di body

	err := h.authService.Logout(req.RefreshToken)
	if err != nil {
		logger.Error().Err(err).Msg("❌ Logout failed")
		response.InternalServerError(c, "An error occurred during logout")
		return
	}

	logger.Info().Msg("✅ User logged out successfully")
	response.Success(c, "Logout successful", nil)
}

// GetCurrentUser menangani GET /api/v1/auth/me
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	// Ambil user_id dari context (sudah di-set oleh middleware.AuthRequired)
	userIDStr, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "User not authenticated")
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		logger.Error().Err(err).Str("user_id_str", userIDStr.(string)).Msg("❌ Invalid UUID in context")
		response.BadRequest(c, "Invalid user ID format", nil)
		return
	}

	userResp, err := h.authService.GetCurrentUser(userID)
	if err != nil {
		if err == services.ErrUserNotFound {
			response.NotFound(c, "User not found")
			return
		}
		logger.Error().Err(err).Msg("❌ Get current user failed")
		response.InternalServerError(c, "An error occurred while fetching user profile")
		return
	}

	response.Success(c, "User profile retrieved", userResp)
}