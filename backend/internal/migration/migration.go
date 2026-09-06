package migration

import (
	"fmt"

	"webifylab-backend/internal/config"
	"webifylab-backend/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations runs all pending migrations
func RunMigrations() error {
	// Build database URL
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
		config.AppConfig.DBSSLMode,
	)

	// Path ke folder migrations (relative dari binary location)
	migrationsPath := "file://migrations"

	logger.Info().
		Str("db_url", maskPassword(dbURL)).
		Str("migrations_path", migrationsPath).
		Msg("🔄 Starting migrations")

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// Run migrations
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			logger.Info().Msg("✅ No migration changes - database is up to date")
			return nil
		}
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	logger.Info().Msg("✅ All migrations applied successfully")
	return nil
}

// RollbackMigrations rolls back the last migration
func RollbackMigrations() error {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
		config.AppConfig.DBSSLMode,
	)

	migrationsPath := "file://migrations"

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Steps(-1); err != nil {
		if err == migrate.ErrNoChange {
			logger.Info().Msg("⚠️  No migrations to rollback")
			return nil
		}
		return fmt.Errorf("failed to rollback migration: %w", err)
	}

	logger.Info().Msg("✅ Migration rolled back successfully")
	return nil
}

// ResetMigrations drops all tables and re-runs all migrations
func ResetMigrations() error {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
		config.AppConfig.DBSSLMode,
	)

	migrationsPath := "file://migrations"

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// Drop all
	if err := m.Drop(); err != nil {
		return fmt.Errorf("failed to drop database: %w", err)
	}

	logger.Info().Msg("🗑️  Database dropped")

	// Re-create migrate instance (karena Drop menutup connection)
	m2, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		return fmt.Errorf("failed to re-create migrate instance: %w", err)
	}

	// Run all migrations again
	if err := m2.Up(); err != nil {
		return fmt.Errorf("failed to re-run migrations: %w", err)
	}

	logger.Info().Msg("✅ Database reset and migrated successfully")
	return nil
}

// maskPassword hides password in URL for logging
func maskPassword(url string) string {
	// Simple masking - replace password with ***
	// Format: postgres://user:password@host:port/db
	return "postgres://***:***@***"
}