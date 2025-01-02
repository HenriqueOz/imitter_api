package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	apperrors "sm.com/m/src/app/app_errors"
	db "sm.com/m/src/app/database"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/utils"
)

func CreateUser(UserSignUp *models.UserSignUp) (err error) {
	result, err := db.Conn.Exec(`
		INSERT INTO user(uuid, name, email, password)
			VALUES (UUID(), ?, ?, ?)
	`,
		UserSignUp.Name,
		UserSignUp.Email,
		utils.HashPassword(UserSignUp.Password),
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

		return apperrors.ErrCreatingUser
	}

	return nil
}

func SignInWithEmail(email string, password string) (*models.UserSignIn, error) {
	result, err := db.Conn.Query(`
		SELECT uuid, name
		FROM user
		WHERE email = ? AND password = ?
	`, email, utils.HashPassword(password))

	return verifySignIn(result, err)
}

func SignInWithName(name string, password string) (*models.UserSignIn, error) {
	result, err := db.Conn.Query(`
		SELECT uuid, name
		FROM user
		WHERE name = ? AND password = ?
	`, name, utils.HashPassword(password))

	return verifySignIn(result, err)
}

func verifySignIn(result *sql.Rows, err error) (*models.UserSignIn, error) {
	if err != nil {
		fmt.Printf("error sign in: %v", err)
		return nil, apperrors.ErrSignIn
	}

	if !result.Next() {
		fmt.Printf("error sign in: %v", apperrors.ErrWrongLogin)
		return nil, apperrors.ErrSignIn
	}

	user := &models.UserSignIn{}
	result.Scan(&user.Uuid, &user.Name)

	return user, nil
}
