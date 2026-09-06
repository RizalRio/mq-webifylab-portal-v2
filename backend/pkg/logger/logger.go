package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Init initializes the global logger based on environment
func Init(env string) {
	// Set global log level based on environment
	switch env {
	case "production":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		// JSON format for production (machine-readable)
		log.Logger = zerolog.New(os.Stdout).
			With().
			Timestamp().
			Str("service", "webifylab-backend").
			Str("env", env).
			Logger()
	default:
		// Development: pretty print for human readability
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		log.Logger = zerolog.New(output).
			With().
			Timestamp().
			Str("service", "webifylab-backend").
			Str("env", env).
			Logger()
	}

	log.Info().Msg("Logger initialized")
}

// SetOutput allows changing the output destination (useful for testing)
func SetOutput(w io.Writer) {
	log.Logger = log.Output(w)
}

// Info logs info level message
func Info() *zerolog.Event {
	return log.Info()
}

// Error logs error level message
func Error() *zerolog.Event {
	return log.Error()
}

// Warn logs warn level message
func Warn() *zerolog.Event {
	return log.Warn()
}

// Debug logs debug level message
func Debug() *zerolog.Event {
	return log.Debug()
}

// Fatal logs fatal level message and exits
func Fatal() *zerolog.Event {
	return log.Fatal()
}

// WithError adds error to the log context
func WithError(err error) *zerolog.Event {
	return log.Error().Err(err)
}