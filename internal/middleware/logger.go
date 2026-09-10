// responseRecorder es un envoltorio alrededor de http.ResponseWriter que permite capturar el código de estado HTTP y si se ha escrito una respuesta.
package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// responseRecorder es un envoltorio alrededor de http.ResponseWriter que permite capturar el código de estado HTTP y si se ha escrito una respuesta.
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

// newResponseRecorder crea un nuevo responseRecorder que envuelve un http.ResponseWriter.
func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

// WriteHeader implementa la interfaz http.ResponseWriter y asegura que se llame a WriteHeader solo una vez.
func (rr *responseRecorder) WriteHeader(code int) {
	if !rr.written {
		rr.statusCode = code
		rr.written = true
		rr.ResponseWriter.WriteHeader(code)
	}
}

// Write implementa la interfaz http.ResponseWriter y asegura que se llame a WriteHeader si no se ha llamado antes.
func (rr *responseRecorder) Write(b []byte) (int, error) {
	if !rr.written {
		rr.WriteHeader(http.StatusOK)
		rr.written = true
	}
	return rr.ResponseWriter.Write(b)
}

// Logger es un middleware que registra información sobre cada solicitud HTTP, incluyendo el método, la ruta, el código de estado y la duración de la solicitud.
func Logger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rec := newResponseRecorder(w)

			// llamar al siguiente handler
			next.ServeHTTP(rec, r)

			logger.InfoContext(r.Context(), "Request completado",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", rec.statusCode),
				slog.Duration("duration", time.Since(start)),
				slog.String("remote_addr", r.RemoteAddr),
			)
		})
	}
}
