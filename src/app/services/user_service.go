package services

import (
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(),
	}
}

func (service *UserService) UpdateUserPassword(uuid string, newPassword string, password string) error {
	err := utils.ValidatePassword(newPassword)
	if err != nil {
		return err
	}

	if password == newPassword {
		return apperrors.ErrNewAndOldPasswordEquals
	}

	err = service.userRepository.UpdateUserPassword(uuid, newPassword, password)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserService) UpdateUserName(uuid string, name string, newName string, password string) error {
	err := utils.ValidateName(name)
	if err != nil {
		return err
	}

	if name == newName {
		return apperrors.ErrNewAndOldNameEquals
	}

	err = service.userRepository.UpdateUserName(uuid, newName, password)
	if err != nil {
		return err
	}

	return nil
}
