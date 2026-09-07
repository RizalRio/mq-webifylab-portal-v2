package middleware

import (
	"net/http"

	"webifylab-backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	limitergin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimit membuat middleware pembatas request berdasarkan IP
// Format limit: "5-M" (5 per menit), "100-H" (100 per jam)
func RateLimit(limitFormatted string) gin.HandlerFunc {
	// Parse format limit (misal: "5-M")
	rate, err := limiter.NewRateFromFormatted(limitFormatted)
	if err != nil {
		logger.Error().Err(err).Str("limit", limitFormatted).Msg("⚠️ Failed to parse rate limit, falling back to default")
		rate, _ = limiter.NewRateFromFormatted("100-M") // Fallback aman
	}

	// Gunakan in-memory store (sangat efisien untuk VPS 1GB RAM)
	store := memory.NewStore()

	// Inisialisasi limiter instance
	instance := limiter.New(store, rate)

	// Opsi kustom untuk response JSON sesuai API Specification
	options := []limitergin.Option{
		limitergin.WithLimitReachedHandler(func(c *gin.Context) {
			c.Header("Content-Type", "application/json")
			// Format sesuai API Spec Section 6.3
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"success":     false,
				"message":     "Too many requests. Please try again later.",
				"errors":      []interface{}{},
				"retry_after": 60,
			})
		}),
	}

	// Buat middleware Gin dan langsung return karena itu adalah gin.HandlerFunc
	return limitergin.NewMiddleware(instance, options...)
}