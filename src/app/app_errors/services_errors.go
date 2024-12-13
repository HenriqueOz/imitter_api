package apperrors

import "errors"

var (
	ErrIvalidPassword error = errors.New("invalid password")
	ErrShortPassword  error = errors.New("password too short")
	ErrLongPassword   error = errors.New("password too long")
	ErrInvalidEmail   error = errors.New("invalid email")
	ErrInvalidName    error = errors.New("invalid name")
	ErrShortName      error = errors.New("name too short")
	ErrLongName       error = errors.New("name too long")
)
