package services

import (
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

func CreateUser(userSignIn *models.UserSignIn) error {
	if err := utils.ValidateEmail(userSignIn.Email); err != nil {
		return err
	}

	if err := utils.ValidatePassword(userSignIn.Password); err != nil {
		return err
	}

	if err := utils.ValidateUsername(userSignIn.Name); err != nil {
		return err
	}

	return repositories.CreateUser(userSignIn)
}
