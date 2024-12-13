package repositories

import (
	"fmt"
	"strings"

	db "sm.com/m/src/app/database"
	"sm.com/m/src/app/models"
)

func CreateUser(userSignIn *models.UserSignIn) (err error) {
	result, err := db.Conn.Exec(`
		INSERT INTO user(name, email, password) VALUES(?, ?, ?)
	`,
		userSignIn.Name,
		userSignIn.Email,
		userSignIn.Password,
	)

	if err != nil {
		fmt.Printf("error inserting user in the table: %v\n", err)
		if strings.Contains(err.Error(), "user.UC_email") {
			return ErrEmailAlreadyInUse
		}

		if strings.Contains(err.Error(), "user.UC_name") {
			return ErrNameAlreadyInUse
		}

		return ErrCreatingUser
	}

	fmt.Printf("result: %v\n", result)

	return nil
}
