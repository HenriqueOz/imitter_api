package userrepositorytest

import (
	"database/sql"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/repositories"
	"sm.com/m/src/app/utils"
)

func Test_DeleteUserAccount(t *testing.T) {
	setUp := func() (sqlmock.Sqlmock, *repositories.UserRepository, *sql.DB) {
		db, mock, err := sqlmock.New()
		if err != nil {
			log.Fatalf("not expected error: '%s' when opening database connection\n", err)
		}
		repository := repositories.NewUserRepository(db)
		return mock, repository, db
	}

	t.Run("Should return a ErrUnexpected when fails to open transaction", func(tt *testing.T) {
		mock, repository, db := setUp()
		defer db.Close()

		uuid := "test-uuid"
		password := "randompassword"
		err := errors.New("cool error")

		mock.ExpectBegin().
			WillReturnError(err)

		result := repository.DeleteUserAccount(uuid, password)

		assert.Equal(tt, apperrors.ErrUnexpected, result, "Result should match err")
	})

	t.Run("Should return a ErrUnexpected when find for user", func(tt *testing.T) {
		mock, repository, db := setUp()
		defer db.Close()

		uuid := "test-uuid"
		password := "randompassword"

		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT name FROM user WHERE uuid = \? AND password = \?`).
			WithArgs(uuid, utils.HashSha256(password)).
			WillReturnError(apperrors.ErrUnexpected)

		result := repository.DeleteUserAccount(uuid, password)

		assert.Equal(tt, apperrors.ErrUnexpected, result, "Error should match ErrUnexpected")
	})

	t.Run("Should return ErrWrongPassword when not find user", func(tt *testing.T) {
		mock, repository, db := setUp()
		defer db.Close()

		uuid := "test-uuid"
		password := "randompassword"

		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT name FROM user WHERE uuid = \? AND password = \?`).
			WithArgs(uuid, utils.HashSha256(password)).
			WillReturnRows(sqlmock.NewRows([]string{"name"}))

		result := repository.DeleteUserAccount(uuid, password)

		assert.Equal(tt, apperrors.ErrWrongPassword, result, "Error should match ErrWrongPassword")
	})

	t.Run("Should return ErrUnexpected when fail to delete user", func(tt *testing.T) {
		mock, repository, db := setUp()
		defer db.Close()

		uuid := "test-uuid"
		password := "randompassword"

		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT name FROM user WHERE uuid = \? AND password = \?`).
			WithArgs(uuid, utils.HashSha256(password)).
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow("user"))
		mock.ExpectExec(`DELETE FROM user WHERE uuid = \?`).
			WithArgs(uuid).
			WillReturnError(errors.New("cool error"))

		result := repository.DeleteUserAccount(uuid, password)

		assert.Equal(tt, apperrors.ErrUnexpected, result, "Error should match ErrUnexpected")
	})

	t.Run("Should return unexpected error when fail to get affecteds rows", func(tt *testing.T) {
		mock, repository, db := setUp()
		defer db.Close()

		uuid := "test-uuid"
		password := "randompassword"
		err := errors.New("cool error")

		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT name FROM user WHERE uuid = \? AND password = \?`).
			WithArgs(uuid, utils.HashSha256(password)).
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow("user"))
		mock.ExpectExec(`DELETE FROM user WHERE uuid = \?`).
			WithArgs(uuid).
			WillReturnResult(sqlmock.NewErrorResult(err))

		result := repository.DeleteUserAccount(uuid, password)

		assert.Equal(tt, apperrors.ErrUnexpected, result, "Error should match ErrUnexpected")
	})

	t.Run("Should return unexpected error when affected rows are not equals 1", func(tt *testing.T) {
		mock, repository, db := setUp()
		defer db.Close()

		uuid := "test-uuid"
		password := "randompassword"

		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT name FROM user WHERE uuid = \? AND password = \?`).
			WithArgs(uuid, utils.HashSha256(password)).
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow("user"))
		mock.ExpectExec(`DELETE FROM user WHERE uuid = \?`).
			WithArgs(uuid).
			WillReturnResult(sqlmock.NewResult(1, 100))

		result := repository.DeleteUserAccount(uuid, password)

		assert.Equal(tt, apperrors.ErrUnexpected, result, "Error should match ErrUnexpected")
	})

	t.Run("Should return nil error when delete user with success", func(tt *testing.T) {
		mock, repository, db := setUp()
		defer db.Close()

		uuid := "test-uuid"
		password := "randompassword"

		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT name FROM user WHERE uuid = \? AND password = \?`).
			WithArgs(uuid, utils.HashSha256(password)).
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow("user"))
		mock.ExpectExec(`DELETE FROM user WHERE uuid = \?`).
			WithArgs(uuid).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		result := repository.DeleteUserAccount(uuid, password)

		assert.Nil(tt, result, "Error should be nil")
	})
}
