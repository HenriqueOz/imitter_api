package repositories

import (
	"context"
	"log"
	"strings"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/database"
	"sm.com/m/src/app/utils"
)

type IUserRepository interface {
	FindUserByUUIDAndPassword(uuid string, password string) (bool, error)
	UpdateUserPassword(uuid string, newPassword string, password string) error
	UpdateUserName(uuid string, name string, password string) error
	DeleteUserAccount(uuid string, password string) error
}

type UserRepository struct {
	DB database.Database
}

func NewUserRepository(db database.Database) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindUserByUUIDAndPassword(uuid string, password string) (bool, error) {
	rows, err := r.DB.Query(`
		SELECT name
		FROM user
		WHERE uuid = ? AND password = ?
	`, uuid, utils.HashSha256(password))

	if err != nil {
		return false, apperrors.ErrUnexpected
	}

	defer rows.Close()

	if !rows.Next() {
		return false, nil
	}

	return true, nil
}

func (r *UserRepository) UpdateUserPassword(uuid string, newPassword string, password string) error {
	tx, err := r.DB.BeginTx(context.Background(), nil)
	if err != nil {
		log.Printf("Failed to open transaction: %v\n", err)
		return apperrors.ErrUnexpected
	}

	defer func() {
		if err != nil {
			log.Printf("Transaction rollback due to error: %v\n", err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	userExists, err := r.FindUserByUUIDAndPassword(uuid, password)
	if err != nil {
		return err
	}
	if !userExists {
		return apperrors.ErrWrongPassword
	}

	_, err = tx.Exec(`
		UPDATE user
		SET password = ?
		WHERE uuid = ?
	`, utils.HashSha256(newPassword), uuid)
	if err != nil {
		return apperrors.ErrUnexpected
	}

	return nil
}

func (r *UserRepository) UpdateUserName(uuid string, name string, password string) error {
	tx, err := r.DB.BeginTx(context.Background(), nil)
	if err != nil {
		log.Printf("Failed to open transaction: %v\n", err)
		return apperrors.ErrUnexpected
	}

	defer func() {
		if err != nil {
			log.Printf("Transaction rollback due to error: %v\n", err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	userExists, err := r.FindUserByUUIDAndPassword(uuid, password)
	if err != nil {
		return err
	}
	if !userExists {
		return apperrors.ErrWrongPassword
	}

	_, err = tx.Exec(`
		UPDATE user
		SET name = ?
		WHERE uuid = ?
	`, name, uuid)

	if err != nil {
		if strings.Contains(err.Error(), "user.UC_name") {
			return apperrors.ErrNameAlreadyInUse
		}

		return apperrors.ErrUnexpected
	}

	return nil
}

func (r *UserRepository) DeleteUserAccount(uuid string, password string) error {
	tx, err := r.DB.BeginTx(context.Background(), nil)
	if err != nil {
		log.Printf("Failed to open transaction: %v\n", err)
		return apperrors.ErrUnexpected
	}

	defer func() {
		if err != nil {
			log.Printf("Transaction rollback due to error: %v\n", err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	userExists, err := r.FindUserByUUIDAndPassword(uuid, password)
	if err != nil {
		return apperrors.ErrUnexpected
	}
	if !userExists {
		return apperrors.ErrWrongPassword
	}

	result, err := tx.Exec(`
		DELETE FROM user
		WHERE uuid = ?
	`, uuid)

	if err != nil {
		return apperrors.ErrUnexpected
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return apperrors.ErrUnexpected
	}

	return nil
}
