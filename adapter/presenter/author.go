package presenter

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/presenter"
)

type AuthorPresenterImpl struct{}

func NewAuthorPresenter() presenter.AuthorPresenter {
	return &AuthorPresenterImpl{}
}

func (presenter *AuthorPresenterImpl) ResponseUsers(ctx context.Context, authors []*model.Author) ([]*model.Author, error) {
	for _, u := range authors {
		u.Name = "Mr." + u.Name
	}

	return authors, nil
}
