package interactor

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/presenter"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

type AuthorInteractor interface {
	GetAll(ctx context.Context) ([]*model.Author, error)
	InsertDocument(ctx context.Context, author *model.Author) error
}

type AuthorInteractorImpl struct {
	AuthorRepositoryReader         repository.AuthorRepository
	AuthorRepositoryWriter         repository.AuthorRepository
	AuthorRepositoryEventPublisher repository.AuthorRepository
	AuthorPresenter                presenter.AuthorPresenter
}

type Option func(impl *AuthorInteractorImpl)

func NewAuthorInteractor(r repository.AuthorRepositoryWriter, p presenter.AuthorPresenter) AuthorInteractor {
	return &AuthorInteractorImpl{r, p}
}

func (ai *AuthorInteractorImpl) GetAll(ctx context.Context) ([]*model.Author, error) {
	authors, err := ai.AuthorRepositoryReader.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return ai.AuthorPresenter.ResponseUsers(ctx, authors)
}

func (ai *AuthorInteractorImpl) InsertDocument(ctx context.Context, author *model.Author) (err error) {
	return
}
