package repositories

import "errors"

var (
	ErrCreatingUser      error = errors.New("error creating user")
	ErrEmailAlreadyInUse error = errors.New("email already in use")
	ErrNameAlreadyInUse  error = errors.New("name already in use")
)
