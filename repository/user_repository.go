package repository

import "github.com/danangkonang/crud-rest/entity"

type UserRepository interface {
	FindById(id string) *entity.User
}
