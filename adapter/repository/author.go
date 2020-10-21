package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

const authorsTable = "authors"

type AuthorRepositoryImpl struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) repository.AuthorRepository {
	return &AuthorRepositoryImpl{
		db: db,
	}
}

func (author *AuthorRepositoryImpl) FindAll(ctx context.Context) ([]*model.Author, error) {
	query := fmt.Sprintf("SELECT id, name, email FROM %s", authorsTable)
	rows, err := author.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var authors []*model.Author

	for rows.Next() {
		var (
			id    sql.NullInt32
			name  sql.NullString
			email sql.NullString
		)

		err := rows.Scan(&id, &name, &email)

		if err != nil {
			return nil, err
		}

		authors = append(authors, &model.Author{ID: uint(id.Int32), Name: name.String, Email: email.String})
	}

	return authors, nil
}
