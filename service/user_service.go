package service

import (
	"errors"

	"github.com/danangkonang/crud-rest/entity"
	"github.com/danangkonang/crud-rest/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func (service UserService) Get(id string) (*entity.User, error) {
	user := service.Repository.FindById(id)
	if user == nil {
		return nil, errors.New("not found")
	} else {
		return user, nil
	}
}
