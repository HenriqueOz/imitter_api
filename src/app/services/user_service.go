package services

import (
	"fmt"

	apperrors "sm.com/m/src/app/app_errors"
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

func SignInWithEmail(email string, password string) (*models.UserAuth, error) {
	if err := utils.ValidateEmail(email); err != nil {
		return nil, err
	}

	user, err := repositories.SignInWithEmail(email, password)
	if err != nil {
		return nil, err
	}
	return GetUserAuth(user)
}

func SignInWithName(name string, password string) (*models.UserAuth, error) {
	if err := utils.ValidateUsername(name); err != nil {
		return nil, err
	}

	user, err := repositories.SignInWithName(name, password)
	if err != nil {
		return nil, err
	}
	return GetUserAuth(user)
}

func GetUserAuth(user *models.UserSignIn) (*models.UserAuth, error) {
	tokenString, err := utils.GenerateJwtToken(user)
	if err != nil {
		fmt.Printf("error signing jwt token: %v\n", err)
		return nil, apperrors.ErrSignIn
	}

	refreshTokenString, err := utils.GenerateRefreshJwtToken(tokenString)
	if err != nil {
		fmt.Printf("error signing jwt token: %v\n", err)
		return nil, apperrors.ErrSignIn
	}

	return &models.UserAuth{
		AccessToken:  "Bearer " + tokenString,
		RefreshToken: "Bearer " + refreshTokenString,
	}, nil
}
