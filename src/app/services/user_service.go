package services

import (
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

func CreateUser(UserSignUp *models.UserSignUp) error {
	if err := utils.ValidateEmail(UserSignUp.Email); err != nil {
		return err
	}

	if err := utils.ValidatePassword(UserSignUp.Password); err != nil {
		return err
	}

	if err := utils.ValidateUsername(UserSignUp.Name); err != nil {
		return err
	}

	return repositories.CreateUser(UserSignUp)
}
