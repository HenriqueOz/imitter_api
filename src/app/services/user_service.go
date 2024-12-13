package services

import (
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

func CreateUser(userSignIn *models.UserSignIn) error {
	if !utils.ValidateEmail(userSignIn.Email) {
		return apperrors.ErrInvalidEmail
	}

	// if !utils.ValidatePassword(userSignIn.Password) {
	// 	return apperrors.ErrIvalidPassword
	// }

	if !utils.ValidateUsername(userSignIn.Name) {
		return apperrors.ErrInvalidName
	}

	userSignIn.Password = utils.HashPassword(userSignIn.Password)

	return repositories.CreateUser(userSignIn)
}
