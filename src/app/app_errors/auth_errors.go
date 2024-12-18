package apperrors

import "errors"

var (
	ErrMissingAuth error = errors.New("missing auth token")
)
