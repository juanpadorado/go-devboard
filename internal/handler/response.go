// Package handler defines the structures and functions related to handling API responses, including error responses. It provides a standardized way to represent errors and validation issues in the API responses.
package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/juanpadorado/devboard/internal/domain"
	"github.com/juanpadorado/devboard/internal/validator"
)

// errorResponse represents the structure of an error response that can be returned by the API. It includes an error message and an optional slice of validation errors for detailed information about validation failures.
type errorResponse struct {
	Error   string                      `json:"error"`
	Details []validator.ValidationError `json:"details,omitempty"`
}

// ProblemDetails represents the structure of a problem details response that can be returned by the API. It includes information about the type of problem, title, status code, detailed message, instance, and an optional slice of validation errors for detailed information about validation failures.
type ProblemDetails struct {
	Type     string                      `json:"type"`
	Title    string                      `json:"title"`
	Status   int                         `json:"status"`
	Detail   string                      `json:"detail"`
	Instance string                      `json:"instance"`
	Errors   []validator.ValidationError `json:"errors,omitempty"`
}

const problemBaseUrl = "https://example.com/problems"

// RespondError handles the response for different types of errors that may occur during API processing. It takes an HTTP response writer, the incoming request, a logger, and the error to be handled. Based on the type of error, it sets the appropriate HTTP status code and error message in the response. If the error is not classified, it logs the error details for further investigation.
func RespondError(writer http.ResponseWriter, request *http.Request, logger *slog.Logger, err error) {
	var problem ProblemDetails
	problem.Instance = request.URL.Path

	switch {
	case errors.Is(err, domain.ErrNotFound):
		problem.Type = problemBaseUrl + "/not-found"
		problem.Title = "Not Found"
		problem.Status = http.StatusNotFound
		problem.Detail = err.Error()
	case errors.Is(err, domain.ErrAlreadyExists):
		problem.Type = problemBaseUrl + "/already-exists"
		problem.Title = "Already Exists"
		problem.Status = http.StatusConflict
		problem.Detail = err.Error()
	case errors.Is(err, domain.ErrUnauthorized):
		problem.Type = problemBaseUrl + "/unauthorized"
		problem.Title = "Unauthorized"
		problem.Status = http.StatusUnauthorized
		problem.Detail = err.Error()
	case errors.Is(err, domain.ErrForbidden):
		problem.Type = problemBaseUrl + "/forbidden"
		problem.Title = "Forbidden"
		problem.Status = http.StatusForbidden
		problem.Detail = err.Error()
	case errors.Is(err, domain.ErrInvalidInput):
		problem.Type = problemBaseUrl + "/invalid-input"
		problem.Title = "Invalid Input"
		problem.Status = http.StatusBadRequest
		problem.Detail = err.Error()
	default:
		problem.Type = problemBaseUrl + "/internal-server-error"
		problem.Title = "Internal Server Error"
		problem.Status = http.StatusInternalServerError
		problem.Detail = "internal server error"
		logger.ErrorContext(request.Context(), "Error no clasificado",
			slog.Any("error", err),
			slog.String("path", request.URL.Path),
		)
	}
	writer.Header().Set("Content-Type", "application/problem+json")
	writer.WriteHeader(problem.Status)

	if encodeErr := json.NewEncoder(writer).Encode(problem); encodeErr != nil {
		logger.Error("Error al escribir problem details",
			slog.String("Error", encodeErr.Error()))
	}
}

// RespondValidationError handles the response for validation errors that may occur during API processing. It takes an HTTP response writer and a slice of validation errors. It sets the HTTP status code to Bad Request (400) and includes the validation error details in the response body.
func RespondValidationError(writer http.ResponseWriter, errs []validator.ValidationError) error {
	return respondJSON(writer, http.StatusBadRequest, errorResponse{
		Error:   "Datos de entrada invalidos",
		Details: errs,
	})
}

// RespondJSON sends an error response with the specified HTTP status code and error message. It creates an errorResponse struct and encodes it as JSON in the response body.
func RespondJSON(write http.ResponseWriter, status int, data any) error {
	return respondJSON(write, status, data)
}

// respondJSON is a helper function that sets the appropriate headers and status code for the HTTP response, and encodes the provided data as JSON in the response body. It is used internally by RespondError to send error responses in a standardized format.
func respondJSON(write http.ResponseWriter, status int, data any) error {
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(status)
	return json.NewEncoder(write).Encode(data)
}
