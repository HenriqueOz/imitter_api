package servicestest

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"sm.com/m/src/app/services"
	"sm.com/m/test/mocks"
)

func Test(t *testing.T) {
	userRepositoryMock := mocks.NewUserRepositoryMock()
	service := services.NewUserService(
		userRepositoryMock,
	)

	uuid := "random"
	password := "random"
	e := errors.New("a want a error identic to myself")

	userRepositoryMock.On("DeleteUserAccount", uuid, password).Return(e)

	err := service.DeleteUserAccount(uuid, password)

	assert.Equal(t, e, err)
}
