package interactor

import (
	"context"
	"encoding/json"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/presenter"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

type AuthorInteractor interface {
	GetAll(ctx context.Context) ([]*model.Author, error)
	CreateDocument(ctx context.Context, author *model.Author) error
	Create(ctx context.Context, author *model.Author) (*model.Author, error)
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
		impl.AuthorDocumentRepository = document
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

func (ai *AuthorInteractorImpl) GetAll(ctx context.Context) (authors []*model.Author, err error) {
	// Construct key for redis cache
	// Key for this example GetAll is `all_authors`
	key := "all_authors"

	// Get value from redis based on the key
	authors, err = ai.AuthorCacheRepository.FindAll(ctx, key)

	if authors != nil {
		return ai.AuthorPresenter.ResponseUsers(ctx, authors)
	}

	// If not available then get from sql repo
	authors, err = ai.AuthorSQLRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// Set value for the key
	if err = ai.AuthorCacheRepository.Create(ctx, key, authors); err != nil {
		return authors, err
	}
	return ai.AuthorPresenter.ResponseUsers(ctx, authors)
}

func (ai *AuthorInteractorImpl) CreateDocument(ctx context.Context, author *model.Author) error {
	err := ai.AuthorDocumentRepository.Create(ctx, author)
	if err != nil {
		return err
	}
	return nil
}

func (ai *AuthorInteractorImpl) Create(ctx context.Context, author *model.Author) (*model.Author, error) {
	var (
		err error
	)
	author, err = ai.AuthorSQLRepository.Create(ctx, author)
	if err != nil {
		return author, err
	}
	dataByt, err := json.Marshal(author)
	if err != nil {
		return author, err
	}

	if err := ai.AuthorEventRepository.Publish(ctx, nil, dataByt); err != nil {
		return author, err
	}
	return author, err
}
