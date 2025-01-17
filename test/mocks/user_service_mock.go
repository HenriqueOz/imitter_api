package mocks

type UserServiceMock struct {
	UpdateUserPasswordFunc func(uuid string, newPassword string, password string) error
	UpdateUserNameFunc     func(uuid string, name string, newName string, password string) error
	DeleteUserAccountFunc  func(uuid string, password string) error
}

func (m *UserServiceMock) UpdateUserPassword(uuid string, newPassword string, password string) error {
	if m.UpdateUserPasswordFunc != nil {
		return m.UpdateUserPasswordFunc(uuid, newPassword, password)
	}
	return nil
}
func (m *UserServiceMock) UpdateUserName(uuid string, name string, newName string, password string) error {
	if m.UpdateUserNameFunc != nil {
		return m.UpdateUserNameFunc(uuid, name, newName, password)
	}
	return nil
}
func (m *UserServiceMock) DeleteUserAccount(uuid string, password string) error {
	if m.DeleteUserAccountFunc != nil {
		return m.DeleteUserAccountFunc(uuid, password)
	}
	return nil
}
