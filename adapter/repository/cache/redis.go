package cache

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

type AuthorCacheRepositoryImpl struct {
}

func (AuthorCacheRepositoryImpl) Find(ctx context.Context) ([]*model.Author, error) {
	panic("implement me")
}

func (AuthorCacheRepositoryImpl) Create(ctx context.Context) error {
	panic("implement me")
}

func NewAuthorCacheRepository() repository.AuthorCacheRepository {
	return &AuthorCacheRepositoryImpl{}
}
