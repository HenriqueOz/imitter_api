package repositories

import (
	"log"
	"strings"

	apperrors "sm.com/m/src/app/app_errors"
	db "sm.com/m/src/app/database"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/utils"
)

func CreateUser(user *models.UserModel) (err error) {
	result, err := db.Conn.Exec(`
		INSERT INTO user(uuid, name, email, password)
			VALUES (UUID(), ?, ?, ?)
	`,
		user.Name,
		user.Email,
		utils.HashPassword(user.Password),
	)

	if err != nil {
		if strings.Contains(err.Error(), "user.UC_email") {
			return apperrors.ErrEmailAlreadyInUse
		}

		if strings.Contains(err.Error(), "user.UC_name") {
			return apperrors.ErrNameAlreadyInUse
		}

		log.Printf("Failed to create user: %v result: %v\n", err, result)
		return apperrors.ErrUnexpected
	}

	return nil
}

func LoginWithEmail(email string, password string) (*models.UserModel, error) {
	result, err := db.Conn.Query(`
		SELECT uuid
		FROM user
		WHERE email = ? AND password = ?
	`, email, utils.HashPassword(password))

	if err != nil {
		log.Printf("Failed login with name: %v\n", err)
		return nil, apperrors.ErrUnexpected
	}

	if !result.Next() {
		return nil, apperrors.ErrWrongLogin
	}

	user := &models.UserModel{}
	result.Scan(&user.Uuid)

	return user, nil
}

func LoginWithName(name string, password string) (*models.UserModel, error) {
	result, err := db.Conn.Query(`
		SELECT uuid
		FROM user
		WHERE name = ? AND password = ?
		LIMIT 1
	`, name, utils.HashPassword(password))

	if err != nil {
		log.Printf("Failed login with name: %v\n", err)
		return nil, apperrors.ErrUnexpected
	}

	if !result.Next() {
		return nil, apperrors.ErrWrongLogin
	}

	user := &models.UserModel{}
	result.Scan(&user.Uuid)

	return user, nil
}
