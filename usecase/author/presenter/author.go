package presenter

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
)

// AuthorRepository :
type AuthorPresenter interface {
	ResponseUsers(context.Context, []*model.Author) ([]*model.Author, error)
}
