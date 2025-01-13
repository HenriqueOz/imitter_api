package apperrors

import "errors"

var (
	ErrUnauthourized       error = errors.New("unathourized")
	ErrInternalServerError error = errors.New("internal server error")
	ErrBadRequest          error = errors.New("bad request")
	ErrUnexpected          error = errors.New("an unexpected error has occured")
)
