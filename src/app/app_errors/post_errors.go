package apperrors

import "errors"

var (
	ErrPostTooLong  error = errors.New("post must not be greater than 500 characteres")
	ErrUUIDNotMatch error = errors.New("post's uuid doesn't match user's uuid")
)
