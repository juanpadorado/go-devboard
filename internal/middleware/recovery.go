// Package middleware proporciona middlewares para el servidor HTTP.
package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

// Recovery es un middleware que recupera de pánicos en los manejadores HTTP y registra el error y la pila de llamadas.
func Recovery(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					stack := debug.Stack()
					logger.ErrorContext(r.Context(), "panic  recuperado",
						slog.Any("error", err),
						slog.String("stack", string(stack)),
						slog.String("path", r.URL.Path),
					)

					http.Error(w, "error interno del servidor", http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
