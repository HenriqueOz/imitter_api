package repositories

import (
	"fmt"
	"strings"

	apperrors "sm.com/m/src/app/app_errors"
	db "sm.com/m/src/app/database"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/utils"
)

func CreateUser(userSignIn *models.UserSignIn) (err error) {
	result, err := db.Conn.Exec(`
		INSERT INTO user(name, email, password) VALUES(?, ?, ?)
	`,
		userSignIn.Name,
		userSignIn.Email,
		utils.HashPassword(userSignIn.Password),
	)

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

	fmt.Printf("result: %v\n", result)

	return nil
}
