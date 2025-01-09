package services

import (
	"log"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

func CreateUser(user *models.UserModel) error {
	if err := utils.ValidateEmail(user.Email); err != nil {
		return err
	}

	if err := utils.ValidatePassword(user.Password); err != nil {
		return err
	}

	if err := utils.ValidateName(user.Name); err != nil {
		return err
	}

	return repositories.CreateUser(user)
}

func Login(method string, login string, password string) (*models.UserAuthModel, error) {
	var err error

	var user *models.UserModel
	if method == "email" {
		user, err = LoginWithEmail(login, password)
	} else if method == "name" {
		user, err = LoginWithName(login, password)
	} else {
		err = apperrors.ErrInvalidLoginMethod
	}

	if err != nil {
		return nil, err
	}

	userAuth, err := GetUserAuth(user.Uuid)
	if err != nil {
		return nil, err
	}

	return userAuth, nil
}

func LoginWithEmail(email string, password string) (*models.UserModel, error) {
	err := utils.ValidateEmail(email)
	if err != nil {
		return nil, err
	}

	user, err := repositories.LoginWithEmail(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func LoginWithName(name string, password string) (*models.UserModel, error) {
	err := utils.ValidateName(name)
	if err != nil {
		return nil, err
	}

	user, err := repositories.LoginWithName(name, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserAuth(uuid string) (*models.UserAuthModel, error) {
	tokenString, err := utils.GenerateJwtToken(uuid)
	if err != nil {
		log.Printf("Failed to generate jwt token: %v\n", err)
		return nil, apperrors.ErrUnexpected
	}

	refreshTokenString, err := utils.GenerateRefreshJwtToken(uuid, tokenString)
	if err != nil {
		log.Printf("Failed to generate jwt refresh token: %v\n", err)
		return nil, apperrors.ErrUnexpected
	}

	userAuth := &models.UserAuthModel{
		AccessToken:  "Bearer " + tokenString,
		RefreshToken: "Bearer " + refreshTokenString,
	}

	return userAuth, nil
}
