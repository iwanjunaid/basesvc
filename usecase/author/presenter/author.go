package presenter

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
)

// AuthorPresenter :
type AuthorPresenter interface {
	ResponseUsers(context.Context, []*model.Author) ([]*model.Author, error)
	ResponseUser(context.Context, *model.Author) error
}
