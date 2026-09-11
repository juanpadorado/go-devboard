// Package main
package main

import (
	"os"

	"github.com/juanpadorado/devboard/internal/handler"
	"github.com/juanpadorado/devboard/internal/logger"
	"github.com/juanpadorado/devboard/internal/middleware"
	"github.com/juanpadorado/devboard/internal/server"
)

func main() {
	/* logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})) */
	logger := logger.New(logger.DefaultConfig())

	// Aqui arranca el servidor
	server := server.New(":8080", logger)

	server.Use(middleware.Recovery(logger))
	server.Use(middleware.Logger(logger))

	healthHandler := handler.NewHealthHandler()

	server.RegisterRoutes("GET /health", healthHandler)

	/* server.RegisterRoutes("GET /panic", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("error probocado")
	})) */

	if err := server.Start(); err != nil {
		logger.Error("Error Fatal", "error", err)
		os.Exit(1)
		//log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
