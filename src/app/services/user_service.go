package services

import (
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/database"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

type IUserService interface {
	UpdateUserPassword(uuid string, newPassword string, password string) error
	UpdateUserName(uuid string, name string, newName string, password string) error
	DeleteUserAccount(uuid string, password string) error
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(
			database.Conn,
		),
	}
}

func (s *UserService) UpdateUserPassword(uuid string, newPassword string, password string) error {
	err := utils.ValidatePassword(newPassword)
	if err != nil {
		return err
	}

	if password == newPassword {
		return apperrors.ErrNewAndOldPasswordEquals
	}

	err = s.userRepository.UpdateUserPassword(uuid, newPassword, password)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateUserName(uuid string, newName string, password string) error {
	err := utils.ValidateName(newName)
	if err != nil {
		return err
	}

	err = s.userRepository.UpdateUserName(uuid, newName, password)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUserAccount(uuid string, password string) error {
	return s.userRepository.DeleteUserAccount(uuid, password)
}
