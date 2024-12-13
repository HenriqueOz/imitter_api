package services

import (
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

func CreateUser(userSignIn *models.UserSignIn) error {

	if !utils.ValidateEmail(userSignIn.Email) {
		return ErrInvalidEmail
	}

	if !utils.ValidatePassword(userSignIn.Password) {
		return ErrIvalidPassword
	}

	if !utils.ValidateUsername(userSignIn.Name) {
		return ErrInvalidName
	}

	userSignIn.Password = utils.HashPassword(userSignIn.Password)

	return repositories.CreateUser(userSignIn)
}
