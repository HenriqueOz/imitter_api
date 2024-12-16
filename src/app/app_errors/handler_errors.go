package apperrors

import "errors"

var (
	ErrInternalServerError   error = errors.New("internal server error")
	ErrBadRequest            error = errors.New("bad request")
	ErrUnexpectedError       error = errors.New("unexpected error")
	ErrMissingRequiredFields error = errors.New("missing required fields")
	ErrEmptyPayload          error = errors.New("empty payload")
)
