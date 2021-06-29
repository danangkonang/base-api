package service

import (
	"testing"

	"github.com/danangkonang/crud-rest/entity"
	"github.com/danangkonang/crud-rest/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserService{Repository: userRepository}

func TestUserService_GetNotFound(t *testing.T) {
	userRepository.Mock.On("FindById", "1").Return(nil)
	user, err := userService.Get("1")
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUserService_GetSuccess(t *testing.T) {
	user := entity.User{
		Id:   "2",
		Name: "danang",
	}
	userRepository.Mock.On("FindById", "2").Return(user)
	res, err := userService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, user.Id, res.Id)
	assert.Equal(t, user.Id, res.Id)
}
