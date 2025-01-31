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

func (s *PostService) GetRecent(startDate time.Time, userUUID string) ([]models.PostModel, error) {
	return s.PostRepository.GetRecent(startDate, userUUID)
}

func (s *PostService) GetRecentByPostUserUUID(startDate time.Time, userUUID string, postUserUUID string) ([]models.PostModel, error) {
	if len(postUserUUID) != 36 {
		return nil, apperrors.ErrInvalidRequest
	}

	return s.PostRepository.GetRecentByPostUserUUID(startDate, userUUID, postUserUUID)
}

func (s *PostService) GetMyRecent(startDate time.Time, userUUID string) ([]models.PostModel, error) {
	return s.PostRepository.GetRecentByPostUserUUID(startDate, userUUID, userUUID)
}

func (s *PostService) GetRecentFollowing(startDate time.Time, userUUID string) ([]models.PostModel, error) {
	return s.PostRepository.GetRecentFollowing(startDate, userUUID)
}

func (s *PostService) ToogleLike(userUUID string, postId uint64) error {
	return s.PostRepository.ToogleLike(userUUID, postId)
}

func (s *PostService) DeletePost(postUserUUID string, userUUID string, postId uint64) error {
	if postUserUUID != userUUID {
		return apperrors.ErrUUIDNotMatch
	}

	return s.PostRepository.DeletePost(postId, userUUID)
}
