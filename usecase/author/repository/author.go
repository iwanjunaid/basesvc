package repository

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
)

type AuthorSQLRepository interface {
	FindAll(ctx context.Context) ([]*model.Author, error)
	Create(ctx context.Context, author *model.Author) (*model.Author, error)
}

type AuthorDocumentRepository interface {
	FindAll(ctx context.Context) ([]*model.Author, error)
	Create(ctx context.Context, author *model.Author) error
}

type AuthorCacheRepository interface {
	Find(ctx context.Context, key string) (*model.Author, error)
	FindAll(ctx context.Context, key string) ([]*model.Author, error)
	Create(ctx context.Context, key string, value interface{}) error
}

type AuthorEventRepository interface {
	Publish(ctx context.Context, key, message []byte) (err error)
}
