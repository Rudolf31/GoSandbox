package customeerrors

import "errors"

var (
	ErrBadRequest  = errors.New("Toooo bad request")
	ErrNotFound    = errors.New("Resource not found")
	ErrServerError = errors.New("Server error")
	ErrConflict    = errors.New("Resourse conflicte")
)
