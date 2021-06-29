package repository

import (
	"github.com/danangkonang/crud-rest/entity"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) FindById(id string) *entity.User {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		user := arguments.Get(0).(entity.User)
		return &user
	}
}
