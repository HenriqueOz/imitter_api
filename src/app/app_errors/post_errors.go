package apperrors

import "errors"

var (
	ErrPostTooLong error = errors.New("post must not be greater than 500 characteres")
)
