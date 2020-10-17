package presenter

import "github.com/iwanjunaid/basesvc/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
