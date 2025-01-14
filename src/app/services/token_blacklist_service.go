package services

import (
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/repositories"
)

type BlackListService struct {
	BlackListRepository *repositories.BlackListRepository
}

func NewBlackListService() *BlackListService {
	return &BlackListService{
		BlackListRepository: repositories.NewBlackListRepository(),
	}
}

func (service *BlackListService) AddTokenToBlacklist(uuid string) error {
	if len(uuid) != 36 {
		return apperrors.ErrInvalidClaims
	}

	return service.BlackListRepository.AddTokenToBlacklist(uuid)
}
