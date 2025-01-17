package mocks

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (m *UserRepositoryMock) FindUserByUUIDAndPassword(uuid string, password string) (bool, error) {
	args := m.Called(uuid, password)
	return args.Bool(0), args.Error(1)
}

func (m *UserRepositoryMock) UpdateUserPassword(uuid string, newPassword string, password string) error {
	args := m.Called(uuid, newPassword, password)
	return args.Error(0)
}

func (m *UserRepositoryMock) UpdateUserName(uuid string, name string, password string) error {
	args := m.Called(uuid, name, password)
	return args.Error(0)
}

func (m *UserRepositoryMock) DeleteUserAccount(uuid string, password string) error {
	args := m.Called(uuid, password)
	return args.Error(0)
}

func (m *UserRepositoryMock) HandleTx(tx *sql.Tx, err error) error {
	args := m.Called(tx, err)
	return args.Error(0)
}
