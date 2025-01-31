package repositories

import (
	"context"
	"log"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/database"
)

type FollowRepository struct {
	DB database.Database
}

func NewFollowRepository(db database.Database) *FollowRepository {
	return &FollowRepository{
		DB: db,
	}
}

func (r *FollowRepository) ToogleLFollow(userUUID string, userToFollowUUID string) error {
	ctx := context.Background()

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Failed to start transaction: %v\n", err)
		return apperrors.ErrUnexpected
	}

	defer func() {
		if err != nil {
			log.Printf("Transaction rollback due to an error: %v\n", err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	return nil
}
