package sql

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

const authorsTable = "authors"

type AuthorSQLRepositoryImpl struct {
	db *sqlx.DB
}

func (AuthorSQLRepositoryImpl) FindAll(ctx context.Context) ([]*model.Author, error) {
	panic("implement me")
}

func (AuthorSQLRepositoryImpl) Create(ctx context.Context) error {
	panic("implement me")
}

func NewAuthorRepository(db *sqlx.DB) repository.AuthorSQLRepository {
	repo := &AuthorSQLRepositoryImpl{
		db: db,
	}

	return repo
}
