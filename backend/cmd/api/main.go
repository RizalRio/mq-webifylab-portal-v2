package main

import (
	"fmt"
	"net/http"
	"os"

	"webifylab-backend/internal/config"
	"webifylab-backend/internal/middleware"
	"webifylab-backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file (skip if in production/Docker)
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load("../../.env")
	}

	// Load configuration
	config.LoadConfig()

	// Initialize logger
	logger.Init(config.AppConfig.AppEnv)

	logger.Info().
		Str("env", config.AppConfig.AppEnv).
		Str("port", config.AppConfig.AppPort).
		Msg("Starting WebifyLab API")

	// Setup database connection
	config.SetupDatabase()

	logger.Info().Msg("✅ Database connected")

	// Setup Gin Router
	router := gin.Default()

	// Apply request logger middleware
	router.Use(middleware.RequestLogger())

	// Health Check Endpoint
	router.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "WebifyLab API is running 🚀",
			"service": "backend",
			"env":     config.AppConfig.AppEnv,
		})
	})

	// Test error logging (hapus setelah test)
router.GET("/api/v1/test-error", func(c *gin.Context) {
	err := fmt.Errorf("this is a test error")
	logger.LogError(err, "Test error occurred")
	
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "test error",
	})
})

	// Determine port
	port := config.AppConfig.AppPort
	if port == "" {
		port = "8080"
	}

	logger.Info().Str("port", port).Msg("🌐 Server listening")
	if err := router.Run(":" + port); err != nil {
		logger.Fatal().Err(err).Msg("❌ Failed to start server")
	}
}