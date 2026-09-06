package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv      string
	AppPort     string
	AppHost     string
	
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBSSLMode   string
	
	JWTSecret   string
	JWTExpiry   time.Duration
	RefreshTokenExpiry time.Duration
	
	ResendAPIKey string
	AdminEmail   string
	
	FrontendURL string
	CORSOrigin  string
	
	UploadDir   string
	MaxUploadSize int64
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigFile("../.env") // Path dari backend/ ke root monorepo
	viper.AutomaticEnv()
	
	// Set defaults
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("APP_HOST", "0.0.0.0")
	
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "webifylab_dev")
	viper.SetDefault("DB_PASSWORD", "dev")
	viper.SetDefault("DB_NAME", "webifylab_dev")
	viper.SetDefault("DB_SSLMODE", "disable")
	
	viper.SetDefault("JWT_EXPIRY", "24h")
	viper.SetDefault("REFRESH_TOKEN_EXPIRY", "168h")
	
	viper.SetDefault("FRONTEND_URL", "http://localhost:3000")
	viper.SetDefault("CORS_ORIGIN", "http://localhost:3000")
	
	viper.SetDefault("UPLOAD_DIR", "./uploads")
	viper.SetDefault("MAX_UPLOAD_SIZE", 5242880) // 5MB
	
	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Error reading config file: %v", err)
	}
	
	AppConfig = &Config{
		AppEnv:      viper.GetString("APP_ENV"),
		AppPort:     viper.GetString("APP_PORT"),
		AppHost:     viper.GetString("APP_HOST"),
		
		DBHost:      viper.GetString("DB_HOST"),
		DBPort:      viper.GetString("DB_PORT"),
		DBUser:      viper.GetString("DB_USER"),
		DBPassword:  viper.GetString("DB_PASSWORD"),
		DBName:      viper.GetString("DB_NAME"),
		DBSSLMode:   viper.GetString("DB_SSLMODE"),
		
		JWTSecret:   viper.GetString("JWT_SECRET"),
		JWTExpiry:   viper.GetDuration("JWT_EXPIRY"),
		RefreshTokenExpiry: viper.GetDuration("REFRESH_TOKEN_EXPIRY"),
		
		ResendAPIKey: viper.GetString("RESEND_API_KEY"),
		AdminEmail:   viper.GetString("ADMIN_EMAIL"),
		
		FrontendURL: viper.GetString("FRONTEND_URL"),
		CORSOrigin:  viper.GetString("CORS_ORIGIN"),
		
		UploadDir:   viper.GetString("UPLOAD_DIR"),
		MaxUploadSize: viper.GetInt64("MAX_UPLOAD_SIZE"),
	}
	
	log.Printf("Config loaded successfully (env: %s)", AppConfig.AppEnv)
}