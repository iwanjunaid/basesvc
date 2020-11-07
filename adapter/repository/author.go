package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"github.com/RoseRocket/xerrs"

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
		return nil, xerrs.Mask(err, errors.New("error query"))
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
			return nil, xerrs.Mask(err, errors.New("error query"))
		}

		authors = append(authors, &model.Author{ID: uint(id.Int32), Name: name.String, Email: email.String})
	}

	return authors, nil
}

func (author *AuthorRepositoryImpl) Create(ctx context.Context, entry *model.Author) (err error) {
	q := fmt.Sprintf("INSERT %s SET Name=$1, Email=$2, CreatedAt=$3, UpdatedAt=$4", authorsTable)

	stmt, err := author.db.PrepareContext(ctx, q)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, entry.Name, entry.Email, entry.CreatedAt, entry.UpdatedAt)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	entry.ID = uint(lastID)
	return

}
