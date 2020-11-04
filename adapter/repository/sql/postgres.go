package sql

import (
	"context"
	"database/sql"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

const authorsTable = "authors"

type AuthorSQLRepositoryImpl struct {
	db *sql.DB
}

func (AuthorSQLRepositoryImpl) FindAll(ctx context.Context) ([]*model.Author, error) {
	panic("implement me")
}

func (AuthorSQLRepositoryImpl) Create(ctx context.Context) error {
	panic("implement me")
}

func NewAuthorRepository(db *sql.DB) repository.AuthorSQLRepository {
	repo := &AuthorSQLRepositoryImpl{
		db: db,
	}

	return repo
}
