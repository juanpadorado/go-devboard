// Package handler proporciona manejadores HTTP para la aplicación
package handler

import (
	"encoding/json"
	"net/http"
)

// HealthHandler proporciona un manejador para la ruta de salud del servidor
type HealthHandler struct{}

// NewHealthHandler crea una nueva instancia de HealthHandler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// ServeHTTP implementa la interfaz http.Handler para HealthHandler
type healthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// HealthCheck verifica que el servidor este funcionando
//
// @Summary Health check
// @Description Verifica que el servidor esta corriendo y respondiendo
// @Tags system
// @Produce json
// @Success 200 {object} healthResponse
// @Router /health [get]
func (handler *HealthHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	resp := healthResponse{
		Status:  "ok",
		Version: "1.0.0", // Aquí puedes poner la versión de tu aplicación
	}

	writer.Header().Set("Content-Type", "application/json")
	// writer.WriteHeader(http.StatusOK) // opcional

	if err := json.NewEncoder(writer).Encode(resp); err != nil {
		return
	}
}
