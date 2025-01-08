package apperrors

import "errors"

var (
	ErrIvalidPassword    error = errors.New("invalid password")
	ErrShortPassword     error = errors.New("password must be at least 8 length")
	ErrLongPassword      error = errors.New("password too long")
	ErrInvalidEmail      error = errors.New("invalid email")
	ErrInvalidName       error = errors.New("invalid name")
	ErrShortName         error = errors.New("name too short")
	ErrLongName          error = errors.New("name too long")
	ErrEmailAlreadyInUse error = errors.New("email already in use")
	ErrNameAlreadyInUse  error = errors.New("name already in use")
	ErrWrongLogin        error = errors.New("wrong login or password")
	ErrCreatingUser      error = errors.New("failed to create user")
	ErrLogin             error = errors.New("failed to login")
)
