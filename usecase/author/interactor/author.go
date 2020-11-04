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
	AuthorSQLRepository      repository.AuthorSQLRepository
	AuthorDocumentRepository repository.AuthorDocumentRepository
	AuthorCacheRepository    repository.AuthorCacheRepository
	AuthorEventRepository    repository.AuthorEventRepository
	AuthorPresenter          presenter.AuthorPresenter
}

type Option func(impl *AuthorInteractorImpl)

func NewAuthorInteractor(p presenter.AuthorPresenter, option ...Option) AuthorInteractor {
	interactor := &AuthorInteractorImpl{
		AuthorPresenter: p,
	}
	for _, opt := range option {
		opt(interactor)
	}
	return interactor
}

func AuthorSQLRepository(sql repository.AuthorSQLRepository) Option {
	return func(impl *AuthorInteractorImpl) {
		impl.AuthorSQLRepository = sql
	}
}

func AuthorDocumentRepository(document repository.AuthorDocumentRepository) Option {
	return func(impl *AuthorInteractorImpl) {
		impl.AuthorSQLRepository = document
	}
}

func AuthorCacheRepository(cache repository.AuthorCacheRepository) Option {
	return func(impl *AuthorInteractorImpl) {
		impl.AuthorCacheRepository = cache
	}
}

func AuthorEventRepository(event repository.AuthorEventRepository) Option {
	return func(impl *AuthorInteractorImpl) {
		impl.AuthorEventRepository = event
	}
}

func (ai *AuthorInteractorImpl) GetAll(ctx context.Context) ([]*model.Author, error) {
	authors, err := ai.AuthorSQLRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return ai.AuthorPresenter.ResponseUsers(ctx, authors)
}

func (ai *AuthorInteractorImpl) InsertDocument(ctx context.Context, author *model.Author) (err error) {
	return
}
