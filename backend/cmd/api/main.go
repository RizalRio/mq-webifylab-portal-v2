package main

import (
	"net/http"
	"os"

	"webifylab-backend/internal/config"
	"webifylab-backend/internal/middleware"
	"webifylab-backend/internal/models"
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

	logger.Info().Msg("Database connected")

	// Auto-migrate database (ONLY in development!)
	if config.AppConfig.AppEnv == "development" {
		logger.Info().Msg("Running auto-migration...")
		if err := config.DB.AutoMigrate(models.AllModels()...); err != nil {
			logger.Fatal().Err(err).Msg("❌ Failed to migrate database")
		}
		logger.Info().Msg("Database migrated successfully")
	} else {
		logger.Info().Msg("⏭Skipping auto-migration (production mode)")
	}

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

	// Determine port
	port := config.AppConfig.AppPort
	if port == "" {
		port = "8080"
	}

	logger.Info().Str("port", port).Msg("Server listening")
	if err := router.Run(":" + port); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}