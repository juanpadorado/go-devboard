// Package server proporciona una implementación de un servidor HTTP
package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server representa un servidor HTTP
type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
	logger     *slog.Logger
}

// New crea una nueva instancia del servidor HTTP
func New(addr string, logger *slog.Logger) *Server {
	mux := http.NewServeMux() //Enrutador HTTP

	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux,
		// Timeouts
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
		mux:        mux,
		logger:     logger,
	}
}

// Start inicia el servidor HTTP
func (s *Server) Start() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	defer signal.Stop(quit)

	serveErr := make(chan error, 1)

	go func() {
		s.logger.Info("servidor iniciado", slog.String("addr", s.httpServer.Addr))
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serveErr <- err
		}
	}()

	select {
	case err := <-serveErr:
		return err
	case sig := <-quit:
		s.logger.Info("recibida señal de terminación", slog.String("signal", sig.String()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error("error al cerrar el servidor", slog.Any("error", err))
		return err
	}

	s.logger.Info("servidor cerrado correctamente")

	return nil
}

// RegisterRoutes registra las rutas y sus manejadores en el servidor
func (s *Server) RegisterRoutes(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

// ServeHTTP implementa la interfaz http.Handler para el servidor
func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.mux.ServeHTTP(writer, request)
}

// Use agrega un middleware al servidor
func (s *Server) Use(middleware func(http.Handler) http.Handler) {
	s.httpServer.Handler = middleware(s.httpServer.Handler)
}
