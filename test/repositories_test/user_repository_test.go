package repositoriestest

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/repositories"
)

type MockDB interface {
	*sql.DB
}

func Test_DeleteUserAccount(t *testing.T) {
	var repository *repositories.UserRepository

	setUp := func() {
		repository = &repositories.UserRepository{
			DB: MockDB,
		}
	}

	t.Run("Should return a ErrWrongLogin when password do not match the database registry", func(tt *testing.T) {
		setUp()

		uuid := "test-uuid"
		password := "randompassword"

		err := repository.DeleteUserAccount(uuid, password)

		assert.Equal(tt, apperrors.ErrWrongLogin, err, "Error should match ErrWrongLogin")
	})

	t.Run("Should return a unexpected error when a unexpected error occur", func(tt *testing.T) {
		setUp()

		uuid := "test-uuid"
		password := "randompassword"

		err := repository.DeleteUserAccount(uuid, password)

		assert.Equal(tt, apperrors.ErrWrongLogin, err, "Error should match ErrWrongLogin")
	})

	t.Run("Should delete the user with succes", func(tt *testing.T) {})
}
