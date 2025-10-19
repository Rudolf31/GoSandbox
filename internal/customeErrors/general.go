package customeerrors

import "errors"

var (
	ErrBadRequest = errors.New("Toooo bad request")
	ErrNotFound   = errors.New("Resource not found")
)
