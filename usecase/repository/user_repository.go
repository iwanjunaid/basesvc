package repository

import "github.com/iwanjunaid/basesvc/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
