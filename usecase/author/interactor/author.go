package interactor

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/presenter"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

type AuthorInteractor interface {
	GetAll(ctx context.Context) ([]*model.Author, error)
}

type AuthorInteractorImpl struct {
	AuthorRepository repository.AuthorRepository
	AuthorPresenter  presenter.AuthorPresenter
}

func NewAuthorInteractor(r repository.AuthorRepository, p presenter.AuthorPresenter) AuthorInteractor {
	return &AuthorInteractorImpl{r, p}
}

func (ai *AuthorInteractorImpl) GetAll(ctx context.Context) ([]*model.Author, error) {
	authors, err := ai.AuthorRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return ai.AuthorPresenter.ResponseUsers(ctx, authors)
}
