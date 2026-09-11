// Package logger proporciona una implementación de un logger estructurado y configurable
package logger

import (
	"io"
	"log/slog"
	"os"
)

// Config representa la configuración del logger
type Config struct {
	Level  slog.Level
	JSON   bool
	Output io.Writer
}

// DefaultConfig devuelve una configuración de logger por defecto
func DefaultConfig() Config {
	return Config{
		Level:  slog.LevelDebug,
		JSON:   false,
		Output: os.Stdout,
	}
}

// ProductionConfig devuelve una configuración de logger optimizada para producción
func ProductionConfig() Config {
	return Config{
		Level:  slog.LevelInfo,
		JSON:   true,
		Output: os.Stdout,
	}
}

// New crea un nuevo logger basado en la configuración proporcionada
func New(config Config) *slog.Logger {
	out := config.Output
	if out == nil {
		out = os.Stdout
	}
	opts := &slog.HandlerOptions{
		Level:     config.Level,
		AddSource: config.Level == slog.LevelDebug,
	}

	var handler slog.Handler
	if config.JSON {
		handler = slog.NewJSONHandler(out, opts)
	} else {
		handler = slog.NewTextHandler(out, opts)
	}
	return slog.New(handler)
}
