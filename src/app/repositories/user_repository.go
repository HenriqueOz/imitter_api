package repositories

import (
	"log"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/database"
	"sm.com/m/src/app/utils"
)

func UpdateUserPassword(uuid string, newPassword string, password string) error {
	rows, err := database.Conn.Query(`
		SELECT password
		FROM user
		WHERE
			uuid = ? AND
			password = ?
	`, uuid, utils.HashSha256(password))

	if err != nil {
		log.Printf("Failed to find user by uuid: %v\n", err)
		return apperrors.ErrUnexpected
	}

	if !rows.Next() {
		return apperrors.ErrUserNotFound
	}

	result, err := database.Conn.Exec(`
		UPDATE WHERE
			uuid = ?
		SET
			password = ?
	`, uuid, newPassword)

	if err != nil {
		log.Printf("Failed to find user by uuid: %v\n result: %v", err, result)
		return apperrors.ErrUnexpected
	}

	return nil
}

func UpdateUserName(uuid string, name string, password string) error {
	return nil
}

func DeleteUserAccount(uuid string, password string) error {
	return nil
}
