package main

import (
	"net/http"
	"os"

	"webifylab-backend/internal/config"
	"webifylab-backend/internal/middleware"
	"webifylab-backend/internal/migration"
	"webifylab-backend/pkg/logger"
	"webifylab-backend/pkg/validator"

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

	// Initialize validator
	validator.Init()
	logger.Info().Msg("✅ Validator initialized")

	// Handle CLI commands
	if len(os.Args) > 1 {
		command := os.Args[1]
		handleCommand(command)
		return
	}

	// Default: start server
	startServer()
}

func handleCommand(command string) {
	switch command {
	case "migrate":
		logger.Info().Msg("🔄 Running migrations...")
		if err := migration.RunMigrations(); err != nil {
			logger.Fatal().Err(err).Msg("❌ Migration failed")
		}
		logger.Info().Msg("✅ Migration completed")

	case "rollback":
		logger.Info().Msg("🔄 Rolling back last migration...")
		if err := migration.RollbackMigrations(); err != nil {
			logger.Fatal().Err(err).Msg("❌ Rollback failed")
		}
		logger.Info().Msg("✅ Rollback completed")

	case "reset":
		logger.Warn().Msg("⚠️  Resetting database - ALL DATA WILL BE LOST!")
		if err := migration.ResetMigrations(); err != nil {
			logger.Fatal().Err(err).Msg("❌ Reset failed")
		}
		logger.Info().Msg("✅ Database reset completed")

	case "version":
		logger.Info().Msg("WebifyLab Backend v1.0.0")

	default:
		logger.Error().Str("command", command).Msg("Unknown command")
		os.Exit(1)
	}
}

func startServer() {
	logger.Info().
		Str("env", config.AppConfig.AppEnv).
		Str("port", config.AppConfig.AppPort).
		Msg("🚀 Starting WebifyLab API")

	// Setup database connection
	config.SetupDatabase()
	logger.Info().Msg("✅ Database connected")

	// Auto-migrate in development only
	if config.AppConfig.AppEnv == "development" {
		logger.Info().Msg("🔄 Running auto-migration (development mode)...")
		if err := migration.RunMigrations(); err != nil {
			logger.Warn().Err(err).Msg("⚠️  Auto-migration failed, continuing...")
		}
	}

	// Setup Gin Router
	router := gin.Default()

	// 1. CORS Middleware
	router.Use(middleware.SetupCORS())

	// 2. Request Logger Middleware
	router.Use(middleware.RequestLogger())

	// ============================================
	// PUBLIC ROUTES
	// ============================================
	router.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "WebifyLab API is running 🚀",
			"service": "backend",
			"env":     config.AppConfig.AppEnv,
		})
	})

	// CORS test endpoint
	router.GET("/api/v1/cors-test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "CORS is working correctly",
			"data": map[string]interface{}{
				"origin":       c.GetHeader("Origin"),
				"env":          config.AppConfig.AppEnv,
				"cors_enabled": true,
			},
		})
	})

	// ============================================
	// PROTECTED ROUTES (untuk testing)
	// ============================================
	protected := router.Group("/api/v1/protected")
	protected.Use(middleware.AuthRequired())
	protected.Use(middleware.RateLimit("100-M"))

	{
		protected.GET("/me", func(c *gin.Context) {
			userID := c.GetString("user_id")
			role := c.GetString("role")
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Welcome to protected area",
				"data": gin.H{
					"user_id": userID,
					"role":    role,
				},
			})
		})

		adminOnly := protected.Group("/admin")
		adminOnly.Use(middleware.RequireRole("super_admin", "admin"))
		adminOnly.Use(middleware.RateLimit("5-M"))

		{
			adminOnly.GET("/dashboard", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"success": true,
					"message": "Welcome to Admin Dashboard",
					"data":    "Secret admin data here",
				})
			})
		}
	}

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