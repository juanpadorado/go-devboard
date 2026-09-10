// Package main
package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/juanpadorado/devboard/internal/handler"
	"github.com/juanpadorado/devboard/internal/middleware"
	"github.com/juanpadorado/devboard/internal/server"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	// Aqui arranca el servidor
	server := server.New(":8080")

	server.Use(middleware.Recovery(logger))
	server.Use(middleware.Logger(logger))

	healthHandler := handler.NewHealthHandler()

	server.RegisterRoutes("GET /health", healthHandler)

	server.RegisterRoutes("GET /panic", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("error probocado")
	}))

	logger.Info("Servidor iniciado en :8080")

	if err := server.Start(); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
