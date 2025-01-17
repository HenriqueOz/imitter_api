package provider

import "sm.com/m/src/app/services"

type IServiceProvider interface {
	NewUserService() *services.UserService
}

type ServiceProvider struct {
	RepositoryProvider
}

func (p *ServiceProvider) NewUserService() *services.UserService {
	return services.NewUserService(
		p.NewUserRepository(),
	)
}
