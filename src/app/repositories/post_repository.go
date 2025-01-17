package repositories

import (
	"sm.com/m/src/app/database"
)

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
	// ctx := context.Background()
	// result, err := r.DB.ExecContext(ctx, `
	// 	INSERT INTO
	// 		post(content, user_id, date)
	// 	VALUES
	// 		(?, ?, ?)
	// `)

	return nil
}
