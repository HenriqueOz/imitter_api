package repositories

import (
	"context"
	"database/sql"
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
	HandleTx(tx *sql.Tx, err error) error
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

func (r *UserRepository) HandleTx(tx *sql.Tx, err error) error {
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		return err
	}
	return tx.Commit()
}

func (r *UserRepository) UpdateUserPassword(uuid string, newPassword string, password string) error {
	tx, err := r.DB.BeginTx(context.Background(), nil)
	if err != nil {
		log.Printf("Failed to open transaction: %v\n", err)
		return apperrors.ErrUnexpected
	}

	userExists, err := r.FindUserByUUIDAndPassword(uuid, password)
	if err != nil {
		return r.HandleTx(tx, err)
	}
	if !userExists {
		return r.HandleTx(tx, apperrors.ErrWrongPassword)
	}

	result, err := tx.Exec(`
		UPDATE user
		SET password = ?
		WHERE uuid = ?
	`, utils.HashSha256(newPassword), uuid)

	if err != nil {
		return r.HandleTx(tx, apperrors.ErrUnexpected)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return r.HandleTx(tx, apperrors.ErrUnexpected)
	}

	return r.HandleTx(tx, nil)
}

func (r *UserRepository) UpdateUserName(uuid string, name string, password string) error {
	tx, err := r.DB.BeginTx(context.Background(), nil)
	if err != nil {
		log.Printf("Failed to open transaction: %v\n", err)
		return apperrors.ErrUnexpected
	}

	userExists, err := r.FindUserByUUIDAndPassword(uuid, password)
	if err != nil {
		return r.HandleTx(tx, err)
	}
	if !userExists {
		return r.HandleTx(tx, apperrors.ErrWrongPassword)
	}

	result, err := tx.Exec(`
		UPDATE user
		SET name = ?
		WHERE uuid = ?
	`, name, uuid)

	if err != nil {
		if strings.Contains(err.Error(), "user.UC_name") {
			return r.HandleTx(tx, apperrors.ErrNameAlreadyInUse)
		}

		return r.HandleTx(tx, apperrors.ErrUnexpected)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return r.HandleTx(tx, apperrors.ErrUnexpected)
	}

	return r.HandleTx(tx, nil)
}

func (r *UserRepository) DeleteUserAccount(uuid string, password string) error {
	tx, err := r.DB.BeginTx(context.Background(), nil)
	if err != nil {
		log.Printf("Failed to open transaction: %v\n", err)
		return apperrors.ErrUnexpected
	}

	userExists, err := r.FindUserByUUIDAndPassword(uuid, password)
	if err != nil {
		return r.HandleTx(tx, apperrors.ErrUnexpected)
	}
	if !userExists {
		return r.HandleTx(tx, apperrors.ErrWrongPassword)
	}

	result, err := tx.Exec(`
		DELETE FROM user
		WHERE uuid = ?
	`, uuid)

	if err != nil {
		return r.HandleTx(tx, apperrors.ErrUnexpected)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return r.HandleTx(tx, apperrors.ErrUnexpected)
	}

	return r.HandleTx(tx, nil)
}
