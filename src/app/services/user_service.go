package services

import (
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

func UpdateUserPassword(uuid string, newPassword string, password string) error {
	err := utils.ValidatePassword(newPassword)
	if err != nil {
		return err
	}

	if password == newPassword {
		return apperrors.ErrNewAndOldPasswordEquals
	}

	err = repositories.UpdateUserPassword(uuid, newPassword, password)
	if err != nil {
		return err
	}

	return nil
}
