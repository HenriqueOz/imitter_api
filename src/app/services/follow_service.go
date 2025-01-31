package services

import (
	"log"

	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/database"
	"sm.com/m/src/app/repositories"
)

type FollowService struct {
	followRepository *repositories.FollowRepository
}

func NewFollowService() *FollowService {
	return &FollowService{
		followRepository: repositories.NewFollowRepository(
			database.Conn,
		),
	}
}

func (s *FollowService) ToogleFollow(userUUID string, toFollowUUID string) error {
	if userUUID == toFollowUUID {
		log.Printf("Someone tried to follow themselves")
		return apperrors.ErrUnexpected
	}

	return s.followRepository.ToogleLFollow(userUUID, toFollowUUID)
}
