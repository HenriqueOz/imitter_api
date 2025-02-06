package services

import (
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

func (s *PostService) GetRecent(limit int, offset int, userUUID string) ([]models.PostModel, error) {
	return s.PostRepository.GetRecent(limit, offset, userUUID)
}

func (s *PostService) GetRecentByPostUserUUID(limit int, offset int, userUUID string, postUserUUID string) ([]models.PostModel, error) {
	if len(postUserUUID) != int(constants.UUID_LENGTH) {
		return nil, apperrors.ErrInvalidRequest
	}

	return s.PostRepository.GetRecentByPostUserUUID(limit, offset, userUUID, postUserUUID)
}

func (s *PostService) GetMyRecent(limit int, offset int, userUUID string) ([]models.PostModel, error) {
	return s.PostRepository.GetRecentByPostUserUUID(limit, offset, userUUID, userUUID)
}

func (s *PostService) GetRecentFollowing(limit int, offset int, userUUID string) ([]models.PostModel, error) {
	return s.PostRepository.GetRecentFollowing(limit, offset, userUUID)
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
