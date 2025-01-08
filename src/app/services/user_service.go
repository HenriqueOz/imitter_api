package services

import (
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

	if err := utils.ValidateUsername(user.Name); err != nil {
		return err
	}

	return repositories.CreateUser(user)
}

func Login(login string, password string) (*models.UserAuthModel, error) {
	if err := utils.ValidateEmail(login); err != nil {
		return nil, err
	}

	user, err := repositories.LoginWithEmail(login, password)
	if err != nil {
		return nil, err
	}

	userAuth, err := GetUserAuth(user.Uuid)
	if err != nil {
		// TODO handle error
		return nil, err
	}

	return userAuth, nil
}

func GetUserAuth(uuid string) (*models.UserAuthModel, error) {
	tokenString, err := utils.GenerateJwtToken(uuid)
	if err != nil {
		return nil, apperrors.ErrUnexpected
	}

	refreshTokenString, err := utils.GenerateRefreshJwtToken(tokenString)
	if err != nil {
		return nil, apperrors.ErrUnexpected
	}

	userAuth := &models.UserAuthModel{
		AccessToken:  "Bearer " + tokenString,
		RefreshToken: "Bearer " + refreshTokenString,
	}

	return userAuth, nil
}
