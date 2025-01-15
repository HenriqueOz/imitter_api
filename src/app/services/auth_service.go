package services

import (
	"log"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/database"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

type AuthService struct {
	AuthRepository *repositories.AuthRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		AuthRepository: repositories.NewAuthRepository(
			database.Conn,
		),
	}
}

func (s *AuthService) CreateUser(user *models.UserModel) error {
	if err := utils.ValidateEmail(user.Email); err != nil {
		return err
	}

	if err := utils.ValidatePassword(user.Password); err != nil {
		return err
	}

	if err := utils.ValidateName(user.Name); err != nil {
		return err
	}

	return s.AuthRepository.CreateUser(user)
}

func (s *AuthService) Login(method string, login string, password string) (*models.UserAuthModel, error) {
	var err error

	var user *models.UserModel

	switch method {
	default:
		err = apperrors.ErrInvalidLoginMethod
	case "email":
		user, err = s.LoginWithEmail(login, password)
	case "name":
		user, err = s.LoginWithName(login, password)
	}

	if err != nil {
		return nil, err
	}

	userAuth, err := s.GetUserAuth(user.Uuid)
	if err != nil {
		return nil, err
	}

	return userAuth, nil
}

func (s *AuthService) LoginWithEmail(email string, password string) (*models.UserModel, error) {
	err := utils.ValidateEmail(email)
	if err != nil {
		return nil, err
	}

	user, err := s.AuthRepository.LoginWithEmail(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) LoginWithName(name string, password string) (*models.UserModel, error) {
	err := utils.ValidateName(name)
	if err != nil {
		return nil, err
	}

	user, err := s.AuthRepository.LoginWithName(name, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) GetUserAuth(uuid string) (*models.UserAuthModel, error) {
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
