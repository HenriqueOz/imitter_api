package repositories

import (
	"fmt"
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

	fmt.Printf("%v\n", result)

	if err != nil {
		fmt.Printf("error inserting user in the table: %v\n", err)
		if strings.Contains(err.Error(), "user.UC_email") {
			return apperrors.ErrEmailAlreadyInUse
		}

		if strings.Contains(err.Error(), "user.UC_name") {
			return apperrors.ErrNameAlreadyInUse
		}

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
		log.Printf("error sign in: %v\n", err)
		return nil, apperrors.ErrUnexpected
	}

	if !result.Next() {
		log.Printf("error sign in: %v\n", apperrors.ErrWrongLogin)
		return nil, apperrors.ErrWrongLogin
	}

	user := &models.UserModel{}
	result.Scan(&user.Uuid)

	return user, nil
}
