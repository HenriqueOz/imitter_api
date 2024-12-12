package handlers

import "errors"

type RequestError struct {
	StatusCode int
	Err        error
	Message    string
	Details    map[string]any
}

var (
	ErrInternalServerError   error = errors.New("internal server error")
	ErrBadRequest            error = errors.New("bad request")
	ErrUnexpectedError       error = errors.New("unexpected error")
	ErrMissingRequiredFields error = errors.New("missing required fields")
)
