package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func SetupDatabase() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBName,
		AppConfig.DBSSLMode,
	)
	
	// Set log mode based on environment
	logMode := logger.Info
	if AppConfig.AppEnv == "production" {
		logMode = logger.Warn
	}
	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	
	// SetMaxOpenConns sets the maximum number of open connections
	sqlDB.SetMaxOpenConns(20)
	
	// SetMaxIdleConns sets the maximum number of idle connections
	sqlDB.SetMaxIdleConns(10)
	
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused
	sqlDB.SetConnMaxLifetime(time.Hour)
	
	log.Println("Database connected successfully")
}

func GetDB() *gorm.DB {
	return DB
}