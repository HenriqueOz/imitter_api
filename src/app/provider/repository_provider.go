package provider

import (
	"sm.com/m/src/app/database"
	"sm.com/m/src/app/repositories"
)

type IRepositoryProvider interface {
	NewUserRepository(db database.Database) *repositories.IUserRepository
}

type RepositoryProvider struct{}

func (p *RepositoryProvider) NewUserRepository() *repositories.UserRepository {
	return repositories.NewUserRepository(
		database.Conn,
	)
}
