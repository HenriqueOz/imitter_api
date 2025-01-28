package services

import (
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/database"
	"sm.com/m/src/app/repositories"
)

type IBlackListService interface {
	AddTokenToBlacklist(uuid string) error
}

type BlackListService struct {
	BlackListRepository *repositories.BlackListRepository
}

func NewBlackListService() *BlackListService {
	return &BlackListService{
		BlackListRepository: repositories.NewBlackListRepository(
			database.Conn,
		),
	}
}

func (s *BlackListService) AddTokenToBlacklist(uuid string) error {
	if len(uuid) != 36 {
		return apperrors.ErrInvalidClaims
	}

	return s.BlackListRepository.AddTokenToBlacklist(uuid)
}
