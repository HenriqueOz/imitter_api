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

func (s *PostService) GetRecentGlobal(limit int, offset int, myUUID string) ([]models.PostModel, error) {
	return s.PostRepository.GetRecentGlobal(limit, offset, myUUID)
}

func (s *PostService) GetRecentByPostOwnerUUID(limit int, offset int, myUUID string, ownerUUID string) ([]models.PostModel, error) {
	if len(ownerUUID) != int(constants.UUID_LENGTH) {
		return nil, apperrors.ErrInvalidUUIDFormat
	}

	return s.PostRepository.GetRecentByPostOwnerUUID(limit, offset, myUUID, ownerUUID)
}

func (s *PostService) GetMyRecent(limit int, offset int, userUUID string) ([]models.PostModel, error) {
	return s.PostRepository.GetRecentByPostOwnerUUID(limit, offset, userUUID, userUUID)
}

func (s *PostService) GetRecentFollowing(limit int, offset int, myUUID string) ([]models.PostModel, error) {
	return s.PostRepository.GetRecentFollowing(limit, offset, myUUID)
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
