package services

import (
	"time"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/constants"
	"sm.com/m/src/app/database"
	"sm.com/m/src/app/models"
	"sm.com/m/src/app/repositories"
)

type PostService struct {
	PostRepository *repositories.PostRepository
}

func NewPostService() *PostService {
	return &PostService{
		repositories.NewPostRepository(
			database.Conn,
		),
	}
}

func (s *PostService) CreatePost(userUUID string, content string) error {
	if len(content) > int(constants.MAX_POST_SIZE) {
		return apperrors.ErrPostTooLong
	}

	return s.PostRepository.CreatePost(userUUID, content)
}

func (s *PostService) GetRecentByStartDate(startDate time.Time, userUUID string) ([]models.PostModel, error) {
	return s.PostRepository.GetRecentByStartDate(startDate, userUUID)
}
