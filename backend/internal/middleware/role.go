package middleware

import (
	"net/http"

	"webifylab-backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

// RequireRole memastikan user memiliki salah satu role yang diizinkan
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil role dari context (harus sudah di-set oleh AuthRequired)
		userRole, exists := c.Get("role")
		if !exists {
			logger.Error().Str("path", c.Request.URL.Path).Msg("❌ Role not found in context. AuthRequired middleware missing?")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Internal server error",
			})
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Internal server error",
			})
			return
		}

		// Cek apakah role user ada di daftar allowedRoles
		for _, allowed := range allowedRoles {
			if roleStr == allowed {
				c.Next() // Role valid, lanjut
				return
			}
		}

		// Jika tidak ada yang cocok, tolak akses
		logger.Warn().
			Str("path", c.Request.URL.Path).
			Str("user_role", roleStr).
			Strs("allowed_roles", allowedRoles).
			Msg("🚫 Forbidden: Insufficient permissions")

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Forbidden: Insufficient permissions",
			"errors":  []string{"You do not have permission to access this resource"},
		})
	}
}