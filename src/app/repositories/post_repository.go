package repositories

import (
	"context"
	"log"
	"time"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/database"
)

type IPostRepository interface {
	CreatePost(userUUID string, content string) error
}

type PostRepository struct {
	DB database.Database
}

func NewPostRepository(db database.Database) *PostRepository {
	return &PostRepository{
		DB: db,
	}
}

func GetUserIdByUUID(uuid string) {}

func (r *PostRepository) CreatePost(userUUID string, content string) error {
	ctx := context.Background()

	date := time.Now()

	result, err := r.DB.ExecContext(ctx, `
		INSERT INTO
			post(content, user_id, date)
		VALUES
			(?, ?, ?)
	`, content, userUUID, date)

	if err != nil {
		log.Printf("Failed to create post: %v result: %v\n", err, result)
		return apperrors.ErrUnexpected
	}

	return nil
}
