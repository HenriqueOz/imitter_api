package repositories

import (
	"log"
	"strings"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/database"
	"sm.com/m/src/app/utils"
)

type UserRepository struct {
	DB database.Database
}

func NewUserRepository(db database.Database) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) UpdateUserPassword(uuid string, newPassword string, password string) error {

	rows, err := r.DB.Query(`
		SELECT name
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
		return apperrors.ErrWrongPassword
	}

	result, err := r.DB.Exec(`
		UPDATE user
		SET password = ?
		WHERE uuid = ?
	`, utils.HashSha256(newPassword), uuid)

	if err != nil {
		log.Printf("Failed to find user by uuid: %v\n result: %v", err, result)
		return apperrors.ErrUnexpected
	}

	return nil
}

func (r *UserRepository) UpdateUserName(uuid string, name string, password string) error {
	hash := utils.HashSha256(password)
	rows, err := r.DB.Query(`
		SELECT name
		FROM user
		WHERE
			uuid = ? AND
			password = ?
	`, uuid, hash)

	if err != nil {
		log.Printf("Failed to find user by uuid: %v\n", err)
		return apperrors.ErrUnexpected
	}

	if !rows.Next() {
		return apperrors.ErrWrongPassword
	}

	result, err := r.DB.Exec(`
		UPDATE user
		SET name = ?
		WHERE uuid = ?
	`, name, uuid)

	if err != nil {
		log.Printf("Failed to find user by uuid: %v\n result: %v", err, result)

		if strings.Contains(err.Error(), "user.UC_name") {
			return apperrors.ErrNameAlreadyInUse
		}

		return apperrors.ErrUnexpected
	}

	return nil
}

func (r *UserRepository) DeleteUserAccount(uuid string, password string) error {
	return nil
}
