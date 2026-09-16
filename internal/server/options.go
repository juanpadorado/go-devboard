package server

import (
	"log/slog"
	"time"
)

type config struct {
	readTimeout       time.Duration
	writeTimeout      time.Duration
	idleTimeout       time.Duration
	readHeaderTimeout time.Duration
	logger            *slog.Logger
}

func defaultConfig() config {
	return config{
		readTimeout:       10 * time.Second,
		writeTimeout:      30 * time.Second,
		idleTimeout:       60 * time.Second,
		readHeaderTimeout: 5 * time.Second,
		logger:            slog.Default(),
	}
}

// Option es una funcion que modifica la configuracion
type Option func(*config)

// WithReadTimeout configura el timeout de lectura
func WithReadTimeout(duration time.Duration) Option {
	return func(c *config) {
		c.readTimeout = duration
	}
}

// WithWriteTimeout configura el timeout de escritura
func WithWriteTimeout(duration time.Duration) Option {
	return func(c *config) {
		c.writeTimeout = duration
	}
}

// WithLogger configura el logger del servidor
func WithLogger(logger *slog.Logger) Option {
	return func(c *config) {
		c.logger = logger
	}
}

// WithIdleTimeout configura el logger del servidor
func WithIdleTimeout(duration time.Duration) Option {
	return func(c *config) {
		c.idleTimeout = duration
	}
} // WithReadHeaderTimeout configura el logger del servidor
func WithReadHeaderTimeout(duration time.Duration) Option {
	return func(c *config) {
		c.readHeaderTimeout = duration
	}
}
