package apperrors

import "errors"

var (
	ErrInvalidToken  error = errors.New("invalid Token")
	ErrTokenFormat   error = errors.New("wrong token format")
	ErrInvalidClaims error = errors.New("token claims are invalid")

	ErrMissingAuthorization error = errors.New("authorization header required")

	ErrTokenAlreadyClaimed error = errors.New("token already claimed")
)
