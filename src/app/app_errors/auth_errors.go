package apperrors

import "errors"

var (
	ErrInvalidToken         error = errors.New("invalid Token")
	ErrMissingAuthorization error = errors.New("authorization header required")
)
