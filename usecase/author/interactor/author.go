package interactor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/iwanjunaid/basesvc/config"

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

func (ai *AuthorInteractorImpl) GetAll(ctx context.Context) ([]*model.Author, error) {
	authors, err := ai.AuthorSQLRepository.FindAll(ctx)
	if err != nil {
		return nil, err
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
		err    error
		topic  string
		topics = config.GetStringSlice("kafka.topics")
	)
	author, err = ai.AuthorSQLRepository.Create(ctx, author)
	if err != nil {
		return author, err
	}
	dataByt, err := json.Marshal(author)
	if err != nil {
		return author, err
	}

	if len(topics) > 0 {
		topic = topics[0]
	}

	if err := ai.AuthorEventRepository.Publish(ctx, topic, nil, dataByt); err != nil {
		fmt.Println(err.Error())
		return author, err
	}
	return author, err
}
