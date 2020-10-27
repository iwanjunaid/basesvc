package repository

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
)

// AuthorRepository :
type AuthorRepository interface {
	FindAll(ctx context.Context) ([]*model.Author, error)
	InsertDocument(ctx context.Context) error
	Publish(ctx context.Context, topic string, message []byte) (err error)
}
