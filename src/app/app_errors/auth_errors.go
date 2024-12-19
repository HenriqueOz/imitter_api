package apperrors

import "errors"

var (
	ErrInvalidToken error = errors.New("invalid token")
	ErrForbidden    error = errors.New("forbidden")
)
