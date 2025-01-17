package userrepositorytest

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

func Test_FindUserByUUIDAndPassword(t *testing.T) {
	setUp := func() (sqlmock.Sqlmock, *repositories.UserRepository, *sql.DB) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("not expected error: '%s' when opening database connection\n", err)
		}
		repository := repositories.NewUserRepository(
			db,
		)
		return mock, repository, db
	}

	t.Run("Should fail to find and return a ErrUnexpected", func(tt *testing.T) {
		mock, repository, db := setUp()
		defer db.Close()

		uuid := "test-uuid"
		password := "password"

		mock.ExpectQuery(`SELECT name FROM user WHERE uuid = \? AND password = \?`).
			WithArgs(uuid, utils.HashSha256(password)).
			WillReturnError(errors.New("cool error"))

		find, err := repository.FindUserByUUIDAndPassword(uuid, password)

		assert.Equal(tt, apperrors.ErrUnexpected, err, "Err should match ErrUnexpected")
		assert.False(tt, find, "Find result shoul be False")
	})

	t.Run("Should not find a user without errors", func(tt *testing.T) {
		mock, repository, db := setUp()
		defer db.Close()

		uuid := "test-uuid"
		password := "password"

		mock.ExpectQuery(`SELECT name FROM user WHERE uuid = \? AND password = \?`).
			WithArgs(uuid, utils.HashSha256(password)).
			WillReturnRows(sqlmock.NewRows([]string{"name"}))

		find, err := repository.FindUserByUUIDAndPassword(uuid, password)

		assert.Nil(tt, err, "Err should be nil")
		assert.False(tt, find, "Find result shoul be False")
	})

	t.Run("Should find a user with success", func(tt *testing.T) {
		mock, repository, db := setUp()
		defer db.Close()

		uuid := "test-uuid"
		password := "password"

		mock.ExpectQuery(`SELECT name FROM user WHERE uuid = \? AND password = \?`).
			WithArgs(uuid, utils.HashSha256(password)).
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow("cool_name"))

		find, err := repository.FindUserByUUIDAndPassword(uuid, password)

		assert.Nil(tt, err, "Err should be nil")
		assert.True(tt, find, "Find should be true")
	})

}
