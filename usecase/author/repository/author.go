package repository

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
)

// AuthorRepository :
type AuthorRepositoryReader interface {
	FindAll(ctx context.Context) ([]*model.Author, error)
	InsertDocument(ctx context.Context) error
	Publish(ctx context.Context, topic string, message []byte) (err error)
}

type AuthorRepositoryWriter interface {
	FindAll(ctx context.Context) ([]*model.Author, error)
	InsertDocument(ctx context.Context) error
	Publish(ctx context.Context, topic string, message []byte) (err error)
}

type AuthoRepositoryEventPublisher interface {
}
