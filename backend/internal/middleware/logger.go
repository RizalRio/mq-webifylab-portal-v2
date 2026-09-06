package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// RequestLogger logs every HTTP request with method, path, status, duration, and client IP
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate duration
		latency := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		userAgent := c.Request.UserAgent()
		referer := c.Request.Referer()

		// Get error if any
		errors := c.Errors

		// Build log event
		event := log.Info()

		// Add fields
		event.
			Str("method", method).
			Str("path", path).
			Int("status", status).
			Dur("latency", latency).
			Str("client_ip", clientIP).
			Str("user_agent", userAgent)

		// Add query if exists
		if query != "" {
			event.Str("query", query)
		}

		// Add referer if exists
		if referer != "" {
			event.Str("referer", referer)
		}

		// Add body size
		event.Int("body_size", int(c.Request.ContentLength))

		// Log errors if any
		if len(errors) > 0 {
			event.Str("errors", errors.String())
		}

		// Log with appropriate level based on status code
		switch {
		case status >= 500:
			log.Error().
				Str("method", method).
				Str("path", path).
				Int("status", status).
				Dur("latency", latency).
				Str("client_ip", clientIP).
				Msg("Server error")
		case status >= 400:
			log.Warn().
				Str("method", method).
				Str("path", path).
				Int("status", status).
				Dur("latency", latency).
				Str("client_ip", clientIP).
				Msg("Client error")
		default:
			event.Msg("Request completed")
		}
	}
}