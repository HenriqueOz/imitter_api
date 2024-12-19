package repositories

import (
	"fmt"
	"strings"

	apperrors "sm.com/m/src/app/app_errors"
	db "sm.com/m/src/app/database"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/utils"
)

func CreateUser(UserSignUp *models.UserSignUp) (err error) {
	result, err := db.Conn.Exec(`
		INSERT INTO user(name, email, password) VALUES(?, ?, ?)
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
		SELECT * FROM user WHERE name = ? AND password = ?
	`, email, password)

	if err != nil {
		fmt.Printf("error signin with name: %v", err)
		return nil, apperrors.ErrSignIn
	}

	if !result.Next() {
		fmt.Printf("error signin with name: %v", err)
		return nil, apperrors.ErrSignIn
	}

	user := &models.UserSignIn{}
	result.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	return user, nil
}

func SignInWithName(name string, password string) (*models.UserSignIn, error) {
	result, err := db.Conn.Query(`
		SELECT * FROM user WHERE name = ? AND password = ?
	`, name, password)

	if err != nil {
		fmt.Printf("error signin with name: %v", err)
		return nil, apperrors.ErrSignIn
	}

	user := &models.UserSignIn{}
	for result.Next() {
		result.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	}

	return user, nil
}
