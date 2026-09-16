// Package validator provides a simple wrapper around the go-playground/validator package for validating structs and translating validation errors into user-friendly messages.
package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator is a wrapper around the go-playground/validator package that provides methods for validating structs and translating validation errors into user-friendly messages.
type Validator struct {
	validate *validator.Validate
}

// New creates a new instance of Validator with the default configuration.
func New() *Validator {
	return &Validator{
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

// ValidationError represents a validation error for a specific field, including the field name and an error message.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Validate validates the provided struct and returns a slice of ValidationError for any validation errors encountered. If there are no validation errors, it returns nil.
func (v *Validator) Validate(s interface{}) []ValidationError {
	err := v.validate.Struct(s)
	if err == nil {
		return nil
	}

	var errors []ValidationError

	validationErrors, ok := err.(validator.ValidationErrors)

	if !ok {
		return nil
	}

	for _, err := range validationErrors {
		errors = append(errors, ValidationError{
			Field:   strings.ToLower(err.Field()),
			Message: translateTag(err),
		})
	}
	return errors
}

// translateTag translates a validation tag into a user-friendly error message. It takes a validator.FieldError as input and returns a string message describing the validation error.
func translateTag(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "este campo es obligatorio"
	case "email":
		return "debe ser un correo electrónico válido"
	case "min":
		return fmt.Sprintf("debe tener al menos %s caracteres", e.Param())
	case "max":
		return fmt.Sprintf("no debe tener más de %s caracteres", e.Param())
	default:
		return "valor inválido"
	}
}
