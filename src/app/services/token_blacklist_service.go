package services

import (
	apperrors "sm.com/m/src/app/app_errors"
	"sm.com/m/src/app/repositories"
)

func AddTokenToBlacklist(uuid string) error {
	if len(uuid) != 36 {
		return apperrors.ErrInvalidClaims
	}

	return repositories.AddTokenToBlacklist(uuid)
}
