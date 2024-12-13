package services

import "errors"

var (
	ErrIvalidPassword error = errors.New("invalid password")
	ErrInvalidEmail   error = errors.New("invalid email")
	ErrInvalidName    error = errors.New("invalid name")
)
