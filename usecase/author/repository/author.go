package repository

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
)

type AuthorSQLRepository interface {
	FindAll(ctx context.Context) ([]*model.Author, error)
	Create(ctx context.Context) error
}

type AuthorDocumentRepository interface {
	FindAll(ctx context.Context) ([]*model.Author, error)
	Create(ctx context.Context) error
}

type AuthorCacheRepository interface {
	Find(ctx context.Context) ([]*model.Author, error)
	Create(ctx context.Context) error
}

type AuthorEventRepository interface {
	Publish(ctx context.Context, topic string, key, message []byte) (err error)
}
