// Package server proporciona una implementación de un servidor HTTP
package server

import (
	"net/http"
	"time"
)

// Server representa un servidor HTTP
type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
}

// New crea una nueva instancia del servidor HTTP
func New(addr string) *Server {
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
	}
}

// Start inicia el servidor HTTP
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// RegisterRoutes registra las rutas y sus manejadores en el servidor
func (s *Server) RegisterRoutes(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

// ServeHTTP implementa la interfaz http.Handler para el servidor
func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.mux.ServeHTTP(writer, request)
}

func (server *Server) Use(middleware func(http.Handler) http.Handler) {
	server.httpServer.Handler = middleware(server.httpServer.Handler)
}
