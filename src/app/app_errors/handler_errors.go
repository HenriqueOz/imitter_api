package apperrors

import "errors"

var (
	ErrInvalidRequest error = errors.New("missing required field(s)")
	ErrMissingHeaders error = errors.New("missing required header(s)")
)
