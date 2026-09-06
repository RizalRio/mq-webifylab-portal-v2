package logger

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/rs/zerolog/log"
)

// LogError logs an error with file and line information
func LogError(err error, message string) {
	if err == nil {
		return
	}

	// Get caller information
	pc, file, line, ok := runtime.Caller(1)
	caller := "unknown"
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			// Extract function name
			funcName := fn.Name()
			parts := strings.Split(funcName, ".")
			if len(parts) > 0 {
				caller = parts[len(parts)-1]
			}
		}
		// Extract just filename
		fileParts := strings.Split(file, "/")
		if len(fileParts) > 0 {
			file = fileParts[len(fileParts)-1]
		}
	}

	log.Error().
		Err(err).
		Str("caller", caller).
		Str("file", fmt.Sprintf("%s:%d", file, line)).
		Msg(message)
}

// LogErrorWithFields logs an error with additional context fields
func LogErrorWithFields(err error, message string, fields map[string]interface{}) {
	if err == nil {
		return
	}

	event := log.Error().Err(err)

	// Add custom fields
	for key, value := range fields {
		event.Interface(key, value)
	}

	event.Msg(message)
}

// LogPanic logs a panic with stack trace
func LogPanic(err interface{}) {
	// Capture stack trace
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	stackTrace := string(buf[:n])

	log.Error().
		Interface("panic", err).
		Str("stack_trace", stackTrace).
		Msg("PANIC occurred")
}