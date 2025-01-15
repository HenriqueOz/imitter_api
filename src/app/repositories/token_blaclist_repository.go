package repositories

import (
	"log"
	"strings"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/database"
)

type BlackListRepository struct {
	DB database.Database
}

func NewBlackListRepository(db database.Database) *BlackListRepository {
	return &BlackListRepository{
		DB: db,
	}
}

func (r *BlackListRepository) AddTokenToBlacklist(uuid string) error {
	result, err := r.DB.Exec(
		`INSERT INTO token_blacklist(token_uuid)
			VALUES(?)`,
		uuid,
	)

	if err != nil {
		if strings.Contains(err.Error(), "token_blacklist.UC_token_blacklist_token_uuid") {
			return apperrors.ErrTokenAlreadyClaimed
		}

		log.Printf("Failed to insert token uuid into token_blacklist table: %v result: %v", err, result)
		return apperrors.ErrUnexpected
	}

	return nil
}
