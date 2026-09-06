package middleware

import (
	"net/http"
	"strings"
	"time"

	"webifylab-backend/internal/config"
	"webifylab-backend/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupCORS returns a CORS middleware with security-hardened configuration.
//
// SECURITY NOTES:
// 1. We use explicit origin allowlist (NO wildcards, NO reflection)
// 2. Credentials are allowed because we use httpOnly cookies for JWT
// 3. Methods are limited to what the API actually needs
// 4. Headers are limited to what the API actually accepts
// 5. MaxAge is set to 1 hour to reduce preflight requests
// 6. Null origin is explicitly blocked (prevents file:// and sandbox attacks)
func SetupCORS() gin.HandlerFunc {
	// Build allowed origins based on environment
	allowedOrigins := getAllowedOrigins()

	logger.Info().
		Strs("allowed_origins", allowedOrigins).
		Str("env", config.AppConfig.AppEnv).
		Msg("🔒 CORS middleware initialized")

	corsConfig := cors.Config{
		// ============================================
		// 1. ALLOWED ORIGINS (STRICT ALLOWLIST)
		// ============================================
		// SECURITY: Never use AllowAllOrigins=true with AllowCredentials=true
		// This would allow ANY website to make authenticated requests to our API
		AllowOrigins: allowedOrigins,

		// ============================================
		// 2. ALLOWED METHODS (MINIMAL)
		// ============================================
		// Only methods that our API actually uses
		// HEAD and PATCH are intentionally excluded
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},

		// ============================================
		// 3. ALLOWED HEADERS (MINIMAL)
		// ============================================
		// Only headers that our API accepts from clients
		// - Content-Type: for JSON/form-data requests
		// - Authorization: for Bearer token (fallback, though we use cookies)
		// - X-Requested-With: to identify AJAX requests (CSRF protection helper)
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
		},

		// ============================================
		// 4. EXPOSED HEADERS
		// ============================================
		// Headers that JavaScript in browser can read from response
		// Rate limit headers are exposed so frontend can show quota info
		ExposeHeaders: []string{
			"Content-Length",
			"X-RateLimit-Limit",
			"X-RateLimit-Remaining",
			"X-RateLimit-Reset",
		},

		// ============================================
		// 5. CREDENTIALS (CRITICAL FOR JWT COOKIES)
		// ============================================
		// SECURITY: This allows httpOnly cookies to be sent cross-origin
		// REQUIRED because we store JWT in httpOnly cookies
		// WARNING: Only safe because we have strict AllowOrigins
		AllowCredentials: true,

		// ============================================
		// 6. PREFLIGHT CACHE (PERFORMANCE + SECURITY)
		// ============================================
		// Browser will cache preflight (OPTIONS) response for 1 hour
		// Reduces latency and server load
		// 1 hour is a good balance between performance and security
		MaxAge: 1 * time.Hour,
	}

	return cors.New(corsConfig)
}

// getAllowedOrigins returns the list of allowed origins based on environment.
//
// SECURITY: This function uses an explicit allowlist approach.
// We NEVER:
// - Use "*" wildcard (incompatible with credentials)
// - Reflect the Origin header back (vulnerable to origin spoofing)
// - Allow "null" origin (vulnerable to file:// and sandbox attacks)
func getAllowedOrigins() []string {
	origins := []string{}

	// Production origins (always included in production)
	if config.AppConfig.AppEnv == "production" {
		origins = append(origins,
			"https://webifylab.my.id",
		)
		// Note: We do NOT include http://webifylab.my.id in production
		// All traffic must be HTTPS (enforced by Nginx redirect)
	}

	// Development origins (only in development)
	if config.AppConfig.AppEnv == "development" || config.AppConfig.AppEnv == "" {
		origins = append(origins,
			"http://localhost:3000",  // Next.js dev server
			"http://127.0.0.1:3000", // Alternative localhost
			"http://localhost:5173",  // Vite (if used for testing)
		)
	}

	// Additional origins from config (for staging/testing)
	if config.AppConfig.CORSOrigin != "" {
		// Split by comma to support multiple origins
		additionalOrigins := strings.Split(config.AppConfig.CORSOrigin, ",")
		for _, origin := range additionalOrigins {
			origin = strings.TrimSpace(origin)
			if origin != "" && !contains(origins, origin) {
				// SECURITY: Validate origin format before adding
				if isValidOrigin(origin) {
					origins = append(origins, origin)
				} else {
					logger.Warn().
						Str("origin", origin).
						Msg("⚠️  Invalid origin format in CORS_ORIGIN, skipping")
				}
			}
		}
	}

	return origins
}

// isValidOrigin validates that an origin string is properly formatted.
// This prevents misconfiguration that could lead to security issues.
func isValidOrigin(origin string) bool {
	// Must start with http:// or https://
	if !strings.HasPrefix(origin, "http://") && !strings.HasPrefix(origin, "https://") {
		return false
	}

	// Must not contain wildcard
	if strings.Contains(origin, "*") {
		return false
	}

	// Must not be "null"
	if origin == "null" {
		return false
	}

	// Must not contain path (origins are scheme + host + port only)
	// Example: "https://example.com/path" is invalid
	parsed := strings.TrimPrefix(origin, "http://")
	parsed = strings.TrimPrefix(parsed, "https://")
	if strings.Contains(parsed, "/") {
		return false
	}

	return true
}

// contains checks if a string slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// CORSPreflightHandler is an optional manual handler for OPTIONS requests
// if you need custom logic beyond what gin-contrib/cors provides.
// Currently not used, but kept for future reference.
func CORSPreflightHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}