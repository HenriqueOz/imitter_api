package apperrors

import "errors"

var (
	ErrMissingFields  error = errors.New("missing required field(s)")
	ErrMissingHeaders error = errors.New("missing required header(s)")

	ErrInvalidRequest     error = errors.New("invalid request")
	ErrInvalidLoginMethod error = errors.New("login method should be 'email' or 'name'")
)
