package apperrors

import "errors"

var (
	ErrIsNotStruct error = errors.New("provided data is not a struct")
)
