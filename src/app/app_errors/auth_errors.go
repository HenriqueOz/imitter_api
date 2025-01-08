package apperrors

import "errors"

var (
	ErrInvalidToken  error = errors.New("invalid token")
	ErrUnauthourized error = errors.New("unathourized")
)
