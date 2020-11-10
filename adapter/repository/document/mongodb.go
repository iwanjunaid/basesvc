package document

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorDocumentRepository struct {
	db *mongo.Collection
}

func (aD *AuthorDocumentRepository) FindAll(ctx context.Context) ([]*model.Author, error) {
	panic("implement me")
}

func (aD *AuthorDocumentRepository) Create(ctx context.Context, author *model.Author) error {
	_, err := aD.db.InsertOne(ctx, author)
	if err != nil {
		return err
	}

	return nil
}

func NewAuthorDocumentRepository(mdb *mongo.Collection) repository.AuthorDocumentRepository {
	impl := &AuthorDocumentRepository{
		db: mdb,
	}
	return impl
}
