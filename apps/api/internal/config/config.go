package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	Environment   string `mapstructure:"ENVIRONMENT"`
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
	TelegramToken string `mapstructure:"TELEGRAM_TOKEN"`
	TelegramChatID string `mapstructure:"TELEGRAM_CHAT_ID"`
	ResendAPIKey  string `mapstructure:"RESEND_API_KEY"`
	AdminEmail    string `mapstructure:"ADMIN_EMAIL"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found or couldn't be loaded: %v", err)
	}

	viper.SetDefault("PORT", "8080")
	viper.SetDefault("ENVIRONMENT", "development")

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
