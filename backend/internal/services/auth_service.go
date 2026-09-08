package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"webifylab-backend/internal/config"
	"webifylab-backend/internal/dto/request"
	authResponse "webifylab-backend/internal/dto/response"
	"webifylab-backend/internal/repositories"
	"webifylab-backend/pkg/hasher"
	"webifylab-backend/pkg/jwt"
	"webifylab-backend/pkg/logger"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo         *repositories.UserRepository
	refreshTokenRepo *repositories.RefreshTokenRepository
}

// NewAuthService membuat instance AuthService baru
func NewAuthService(
	userRepo *repositories.UserRepository,
	refreshTokenRepo *repositories.RefreshTokenRepository,
) *AuthService {
	return &AuthService{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
	}
}

// Login memverifikasi credentials dan mengembalikan tokens
// SECURITY: Selalu return error yang sama baik email salah atau password salah
// untuk mencegah email enumeration attack
func (s *AuthService) Login(req *request.LoginRequest) (*authResponse.LoginResponse, error) {
	// 1. Cari user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			// SECURITY: Jangan expose apakah email terdaftar atau tidak
			logger.Warn().Str("email", req.Email).Msg("⚠️  Login attempt with non-existent email")
			return nil, ErrInvalidCredentials
		}
		logger.Error().Err(err).Msg("❌ Database error during login")
		return nil, ErrInternalServer
	}

	// 2. Verify password dengan bcrypt
	if !hasher.CheckPassword(req.Password, user.PasswordHash) {
		logger.Warn().Str("email", req.Email).Msg("⚠️  Login attempt with wrong password")
		return nil, ErrInvalidCredentials
	}

	// 3. Generate JWT access token
	accessToken, err := jwt.GenerateToken(user.ID, user.Role)
	if err != nil {
		logger.Error().Err(err).Msg("❌ Failed to generate access token")
		return nil, ErrInternalServer
	}

	// 4. Generate refresh token (random string)
	refreshToken, err := s.generateRandomToken()
	if err != nil {
		logger.Error().Err(err).Msg("❌ Failed to generate refresh token")
		return nil, ErrInternalServer
	}

	// 5. Simpan refresh token ke database
	expiresAt := time.Now().Add(config.AppConfig.RefreshTokenExpiry)
	if err := s.refreshTokenRepo.Create(user.ID, refreshToken, expiresAt); err != nil {
		logger.Error().Err(err).Msg("❌ Failed to save refresh token")
		return nil, ErrInternalServer
	}

	// 6. Build response
	response := &authResponse.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(config.AppConfig.JWTExpiry.Seconds()),
		User: authResponse.UserResponse{
			ID:        user.ID.String(),
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	}

	logger.Info().
		Str("user_id", user.ID.String()).
		Str("email", user.Email).
		Msg("✅ User logged in successfully")

	return response, nil
}

// RefreshToken memvalidasi refresh token dan generate access token baru
func (s *AuthService) RefreshToken(refreshTokenStr string) (*authResponse.RefreshTokenResponse, error) {
	// 1. Cari refresh token di database
	rt, err := s.refreshTokenRepo.GetByToken(refreshTokenStr)
	if err != nil {
		if errors.Is(err, repositories.ErrRefreshTokenNotFound) {
			return nil, ErrInvalidRefreshToken
		}
		logger.Error().Err(err).Msg("❌ Database error during refresh token lookup")
		return nil, ErrInternalServer
	}

	// 2. Cek apakah refresh token sudah expired
	if rt.ExpiresAt.Before(time.Now()) {
		// Hapus refresh token yang expired dari DB untuk kebersihan
		_ = s.refreshTokenRepo.DeleteByToken(refreshTokenStr)
		return nil, ErrRefreshTokenExpired
	}

	// 3. Cari user pemilik refresh token
	user, err := s.userRepo.GetByID(rt.UserID)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			// User sudah dihapus, hapus juga refresh token-nya
			_ = s.refreshTokenRepo.DeleteByToken(refreshTokenStr)
			return nil, ErrInvalidRefreshToken
		}
		logger.Error().Err(err).Msg("❌ Database error fetching user during refresh")
		return nil, ErrInternalServer
	}

	// 4. Generate access token baru
	newAccessToken, err := jwt.GenerateToken(user.ID, user.Role)
	if err != nil {
		logger.Error().Err(err).Msg("❌ Failed to generate new access token")
		return nil, ErrInternalServer
	}

	logger.Info().Str("user_id", user.ID.String()).Msg("✅ Access token refreshed successfully")

	return &authResponse.RefreshTokenResponse{
		AccessToken: newAccessToken,
		ExpiresIn:   int64(config.AppConfig.JWTExpiry.Seconds()),
	}, nil
}

// Logout menghapus refresh token spesifik dari database
func (s *AuthService) Logout(refreshTokenStr string) error {
	if refreshTokenStr == "" {
		// Jika tidak ada refresh token yang dikirim, anggap sukses (client-side cleanup)
		return nil
	}
	
	err := s.refreshTokenRepo.DeleteByToken(refreshTokenStr)
	if err != nil {
		logger.Warn().Err(err).Msg("⚠️ Failed to delete refresh token during logout (might be already deleted)")
	}
	
	return nil
}

// GetCurrentUser mengambil data user berdasarkan ID dari JWT context
func (s *AuthService) GetCurrentUser(userID uuid.UUID) (*authResponse.UserResponse, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		logger.Error().Err(err).Msg("❌ Database error fetching current user")
		return nil, ErrInternalServer
	}

	return &authResponse.UserResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}, nil
}

// generateRandomToken membuat random string untuk refresh token
func (s *AuthService) generateRandomToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// Custom errors untuk Auth Service
var (
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrInvalidRefreshToken = errors.New("invalid or expired refresh token")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
	ErrUserNotFound        = errors.New("user not found")
	ErrInternalServer      = errors.New("internal server error")
)