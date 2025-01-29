package repositories

import (
	"context"
	"log"
	"time"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/database"
	"sm.com/m/src/app/models"
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

	date := time.Now().UTC()

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

func (r *PostRepository) GetRecentByStartDate(startDate time.Time, userUUID string) ([]models.PostModel, error) {
	ctx := context.Background()
	result, err := r.DB.QueryContext(ctx, `
		SELECT user_id, user.name, content, date, likes_count
		FROM post
		INNER JOIN user
		ON user.uuid = user_id
		WHERE user.uuid != ? AND date < ?
		ORDER BY date DESC
		LIMIT 20;
	`, userUUID, startDate)

	if err != nil {
		log.Printf("Failed to create post: %v result: %v\n", err, result)
		return nil, apperrors.ErrUnexpected
	}

	posts := []models.PostModel{}
	for result.Next() {
		p := models.PostModel{}
		err := result.Scan(&p.UserUUID, &p.Author, &p.Content, &p.Date, &p.Likes)

		if err != nil {
			return posts, apperrors.ErrUnexpected
		}

		posts = append(posts, p)
	}

	return posts, nil
}
