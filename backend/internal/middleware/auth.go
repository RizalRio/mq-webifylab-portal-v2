package middleware

import (
	"net/http"
	"strings"

	"webifylab-backend/pkg/jwt"
	"webifylab-backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

// AuthRequired memvalidasi JWT token dari header Authorization
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warn().Str("path", c.Request.URL.Path).Msg("⚠️  Missing Authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Missing authorization header",
				"errors":  []string{"Authorization header is required"},
			})
			return
		}

		// 2. Cek format "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			logger.Warn().Str("path", c.Request.URL.Path).Msg("⚠️  Invalid Authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid authorization header format. Use 'Bearer <token>'",
				"errors":  []string{"Invalid format"},
			})
			return
		}

		tokenString := parts[1]

		// 3. Validasi token
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			logger.Warn().
				Str("path", c.Request.URL.Path).
				Err(err).
				Msg("⚠️  Invalid or expired token")
			
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid or expired token",
				"errors":  []string{err.Error()},
			})
			return
		}

		// 4. Set user info ke Gin context untuk digunakan di handler
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		// Lanjut ke handler berikutnya
		c.Next()
	}
}