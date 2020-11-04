package document

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorDocumentRepository struct {
	mdb *mongo.Database
}

func (AuthorDocumentRepository) FindAll(ctx context.Context) ([]*model.Author, error) {
	panic("implement me")
}

func (AuthorDocumentRepository) Create(ctx context.Context) error {
	panic("implement me")
}

func NewAuthorDocumentRepository(mdb *mongo.Database) repository.AuthorDocumentRepository {
	impl := &AuthorDocumentRepository{
		mdb: mdb,
	}
	return impl
}
