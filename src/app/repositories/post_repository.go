package repositories

import (
	"context"
	"database/sql"
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

func (r *PostRepository) GetRecent(startDate time.Time, userUUID string) ([]models.PostModel, error) {
	ctx := context.Background()
	query := `
		SELECT
			post.id,
			post.user_id,
			user.name,
			post.content,
			post.date,
			post.likes_count,
			EXISTS (
				SELECT 1
				FROM likes
				WHERE likes.post_id = post.id
				AND likes.user_id = ?
			) AS is_liked
		FROM
			post
		INNER JOIN
			user ON user.uuid = post.user_id
		WHERE
			user.uuid != ?
			AND post.date < ?
		ORDER BY
			post.date DESC
		LIMIT 20;
	`
	stmt, err := r.DB.PrepareContext(ctx, query)
	defer func() {
		if err != nil {
			log.Printf("Query failed: %v\n", err)
		}
	}()

	if err != nil {
		return nil, apperrors.ErrUnexpected
	}
	defer stmt.Close()

	result, err := stmt.QueryContext(ctx, userUUID, userUUID, startDate)
	if err != nil {
		return nil, apperrors.ErrUnexpected
	}
	return fetchPosts(result)
}

func (r *PostRepository) GetRecentByPostUserUUID(startDate time.Time, userUUID string, postUserUUID string) ([]models.PostModel, error) {
	ctx := context.Background()
	query := `
		SELECT
			post.id,
			post.user_id,
			user.name,
			post.content,
			post.date,
			post.likes_count,
			EXISTS (
				SELECT 1
				FROM likes
				WHERE likes.post_id = post.id
				AND likes.user_id = ?
			) AS is_liked
		FROM
			post
		INNER JOIN
			user ON user.uuid = post.user_id
		WHERE
			user.uuid = ?
			AND post.date < ?
		ORDER BY
			post.date DESC
		LIMIT 20;
	`

	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Failed to open statement: %v\n", err)
		return nil, apperrors.ErrUnexpected
	}
	defer stmt.Close()

	result, err := stmt.QueryContext(ctx, userUUID, postUserUUID, startDate)
	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
		return nil, apperrors.ErrUnexpected
	}
	return fetchPosts(result)
}

func fetchPosts(result *sql.Rows) ([]models.PostModel, error) {
	posts := []models.PostModel{}
	for result.Next() {
		p := models.PostModel{}
		err := result.Scan(&p.Id, &p.UserUUID, &p.Author, &p.Content, &p.Date, &p.Likes, &p.IsLiked)

		if err != nil {
			return posts, apperrors.ErrUnexpected
		}

		posts = append(posts, p)
	}
	return posts, nil
}

func (r *PostRepository) addLike(userUUID string, postId uint64) error {
	var err error

	ctx := context.Background()
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		log.Printf("Faield to start transaction: %v", err)
		return apperrors.ErrUnexpected
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.ExecContext(ctx, `
		INSERT INTO likes
			(post_id, user_id)
		VALUES
			(?, ?)
	`, postId, userUUID)

	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return apperrors.ErrUnexpected
	}

	_, err = tx.ExecContext(ctx, `
		UPDATE post
		SET likes_count = likes_count + 1
		WHERE id = ?
	`, postId)

	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return apperrors.ErrUnexpected
	}

	return nil
}

func (r *PostRepository) removeLike(userUUID string, postId uint64) error {
	var err error

	ctx := context.Background()
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		log.Printf("Faield to start transaction: %v", err)
		return apperrors.ErrUnexpected
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.ExecContext(ctx, `
		DELETE FROM likes
		WHERE post_id = ? AND user_id = ?
	`, postId, userUUID)

	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return apperrors.ErrUnexpected
	}

	_, err = tx.ExecContext(ctx, `
		UPDATE post
		SET likes_count = likes_count - 1
		WHERE id = ?
	`, postId)

	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return apperrors.ErrUnexpected
	}

	return nil
}

func (r *PostRepository) ToogleLike(userUUID string, postId uint64) error {
	ctx := context.Background()

	var count uint64

	row := r.DB.QueryRowContext(ctx, `
		SELECT COUNT(post_id)
		FROM likes
		WHERE
			user_id = ? AND
			post_id = ?;
	`, userUUID, postId)

	row.Scan(&count)

	if count == 1 {
		return r.removeLike(userUUID, postId)
	} else if count == 0 {
		return r.addLike(userUUID, postId)
	}

	return apperrors.ErrUnexpected
}

func (r *PostRepository) DeletePost(postId uint64, userUUID string) error {
	ctx := context.Background()

	_, err := r.DB.ExecContext(ctx, `
		DELETE FROM post WHERE id = ? AND user_id = ?
	`, postId, userUUID)

	if err != nil {
		log.Printf("Error executing delete: %v\n", err)
		return apperrors.ErrUnexpected
	}

	return nil
}
