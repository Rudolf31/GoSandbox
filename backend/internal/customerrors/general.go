package customerrors

import "errors"

var (
	ErrBadRequest   = errors.New("toooo bad request")
	ErrNotFound     = errors.New("resource not found")
	ErrServerError  = errors.New("server error")
	ErrConflict     = errors.New("resource conflict")
	ErrUnauthorized = errors.New("unauthorized")
)
