package sql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/iwanjunaid/basesvc/internal/telemetry"

	newrelic "github.com/newrelic/go-agent"

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

	var authors []*model.Author

	if nr, ok := ctx.Value("newRelicTransaction").(newrelic.Transaction); ok {
		ctx = newrelic.NewContext(ctx, nr)
	}
	query := fmt.Sprintf(`SELECT id, name, email, created_at, updated_at FROM %s`, authorsTable)
	ds := telemetry.StartDataSegment(ctx, map[string]interface{}{
		"collection":   authorsTable,
		"operation":    "READ",
		"query":        query,
		"query_params": map[string]interface{}{},
	})
	rows, err := as.db.QueryContext(ctx, query)
	telemetry.StopDataSegment(ds)
	if err != nil {
		err = xerrs.Mask(err, errors.New("error query select"))
		return authors, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID                   uuid.UUID
			name, email          string
			createdAt, updatedAt time.Time
		)
		err := rows.Scan(&ID, &name, &email, &createdAt, &updatedAt)
		if err != nil {
			err = xerrs.Mask(err, errors.New("error query select"))
			return authors, err
		}
		authors = append(authors, &model.Author{
			ID:        ID,
			Name:      name,
			Email:     email,
			CreatedAt: createdAt.Unix(),
			UpdatedAt: updatedAt.Unix(),
		})
	}

	return authors, nil
}

func (as *AuthorSQLRepositoryImpl) Create(ctx context.Context, author *model.Author) (*model.Author, error) {
	var (
		id        = uuid.NewV4()
		createdAt = time.Now()
		updatedAt = time.Now()
	)

	query := fmt.Sprintf(`INSERT INTO %s 
		(id, name, email, created_at, updated_at) 
		VALUES 
		($1, $2, $3, $4, $5)`, authorsTable)
	_, err := as.db.ExecContext(ctx, query, id, author.Name, author.Email, createdAt, updatedAt)
	if err != nil {
		err = xerrs.Mask(err, errors.New("error query insert"))
		return author, err
	}
	author.ID = id
	author.CreatedAt = createdAt.Unix()
	author.UpdatedAt = updatedAt.Unix()
	return author, nil
}

func NewAuthorRepository(db *sqlx.DB) repository.AuthorSQLRepository {
	repo := &AuthorSQLRepositoryImpl{
		db: db,
	}

	return repo
}
