package errors

import "errors"

var (
	ErrNotFound       = errors.New("the requested resource was not found")
	ErrInternalServer = errors.New("an unexpected error occurred on the server")
)
