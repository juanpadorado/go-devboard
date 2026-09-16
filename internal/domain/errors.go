package domain

import "errors"

// Common domain errors that can be used throughout the application to represent specific error conditions. These errors can be returned by functions or methods to indicate various failure scenarios, such as not found, already exists, unauthorized access, forbidden actions, invalid input, and internal server errors.
var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrForbidden     = errors.New("forbidden")
	ErrInvalidInput  = errors.New("invalid input")
	ErrInternal      = errors.New("internal error")
)
