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

func (r *FollowRepository) addFollow(userUUID string, toFollowUUID string) error {
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

	_, err = tx.ExecContext(ctx, `
		INSERT INTO follows
			(user_id, follower_id)
		VALUES
			(?, ?)
	`, toFollowUUID, userUUID)

	if err != nil {
		return apperrors.ErrUnexpected
	}

	_, err = tx.ExecContext(ctx, `
		UPDATE user
		SET follows_count = follows_count + 1
		WHERE uuid = ?
	`, toFollowUUID)

	if err != nil {
		return apperrors.ErrUnexpected
	}

	return nil
}

func (r *FollowRepository) removeFollow(userUUID string, toFollowUUID string) error {
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

	_, err = tx.ExecContext(ctx, `
		DELETE FROM follows
		WHERE
			user_id = ?
			AND follower_id  = ?
	`, toFollowUUID, userUUID)

	if err != nil {
		return apperrors.ErrUnexpected
	}

	_, err = tx.ExecContext(ctx, `
		UPDATE user
		SET follows_count = follows_count - 1
		WHERE uuid = ?
	`, toFollowUUID)

	if err != nil {
		return apperrors.ErrUnexpected
	}

	return nil
}

func (r *FollowRepository) ToogleLFollow(userUUID string, toFollowUUID string) error {
	ctx := context.Background()

	result := r.DB.QueryRowContext(ctx, `
		SELECT 1
		FROM follows
		WHERE
			user_id = ?
			AND follower_id = ?
	`, toFollowUUID, userUUID)

	var exists bool
	result.Scan(&exists)

	if exists {
		return r.removeFollow(userUUID, toFollowUUID)
	} else {
		return r.addFollow(userUUID, toFollowUUID)
	}
}
