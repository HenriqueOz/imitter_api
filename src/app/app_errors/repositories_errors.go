package apperrors

import "errors"

var (
	ErrCreatingUser      error = errors.New("error creating user")
	ErrEmailAlreadyInUse error = errors.New("email already in use")
	ErrNameAlreadyInUse  error = errors.New("name already in use")
	ErrEmailNotFound     error = errors.New("email not found")
	ErrNameNotFound      error = errors.New("name not found")
	ErrWrongLogin        error = errors.New("wrong email or password")
	ErrSignIn            error = errors.New("sign in error")
)
