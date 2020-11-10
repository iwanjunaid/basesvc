package sql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/RoseRocket/xerrs"

	"github.com/jmoiron/sqlx"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
	uuid "github.com/satori/go.uuid"
)

const authorsTable = "authors"

type AuthorSQLRepositoryImpl struct {
	db *sqlx.DB
}

func (as *AuthorSQLRepositoryImpl) FindAll(ctx context.Context) ([]*model.Author, error) {
	panic("implement me")
}

func (as *AuthorSQLRepositoryImpl) Create(ctx context.Context, author *model.Author) (*model.Author, error) {
	id := uuid.NewV4()

	query := fmt.Sprintf(`INSERT INTO %s 
	(id, name, email, created_at, updated_at) 
	VALUES
	(:id, :name, :email, :created_at, :updated_at)`, authorsTable)
	params := map[string]interface{}{
		"id":         id,
		"name":       author.Name,
		"email":      author.Email,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	_, err := as.db.NamedExecContext(ctx, query, params)
	if err != nil {
		err = xerrs.Mask(err, errors.New("error query insert"))
		return author, err
	}
	return author, nil
}

func NewAuthorRepository(db *sqlx.DB) repository.AuthorSQLRepository {
	repo := &AuthorSQLRepositoryImpl{
		db: db,
	}

	return repo
}
