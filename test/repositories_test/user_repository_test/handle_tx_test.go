package userrepositorytest

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"sm.com/m/src/app/repositories"
)

func Test_HandleTx(t *testing.T) {
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

	t.Run("Should commit when error is nil", func(tt *testing.T) {
		mock, repository, db := setUp()

		mock.ExpectBegin()
		tx, _ := db.BeginTx(context.Background(), nil)

		mock.ExpectCommit()

		result := repository.HandleTx(tx, nil)

		assert.Equal(tt, nil, result, "Result should be nil")
	})

	t.Run("Should not Rollback when tx is nil and return the given error", func(tt *testing.T) {
		_, repository, _ := setUp()

		err := errors.New("cool error")

		result := repository.HandleTx(nil, err)

		assert.Equal(tt, err, result, "Result should match the err")
	})

	t.Run("Should Rollback when error is not nil and tx is not nil", func(tt *testing.T) {
		mock, repository, db := setUp()

		err := errors.New("cool error")

		mock.ExpectBegin()
		tx, _ := db.BeginTx(context.Background(), nil)

		mock.ExpectRollback()

		result := repository.HandleTx(tx, err)

		assert.Equal(tt, err, result, "Result should match the err")
	})
}
