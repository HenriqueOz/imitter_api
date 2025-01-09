package services

import "sm.com/m/src/app/repositories"

func StoreClaimUuid(uuid string) error {
	return repositories.StoreTokenUuid(uuid)
}
