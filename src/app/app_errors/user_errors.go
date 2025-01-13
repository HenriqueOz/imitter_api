package apperrors

import "errors"

var (
	ErrUserNotFound error = errors.New("Could not find the user")
)
