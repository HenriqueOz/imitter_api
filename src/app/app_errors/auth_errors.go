package apperrors

import "errors"

var (
	ErrInvalidToken         error = errors.New("invalid Token")
	ErrTokenFormat          error = errors.New("wrong token format")
	ErrMissingAuthorization error = errors.New("authorization header required")
)
