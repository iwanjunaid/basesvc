package datastore

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/newrelic/go-agent/_integrations/nrpq"
)

func PostgresConn(dsn string) *sqlx.DB {
	db := sqlx.MustConnect("nrpostgres", dsn)
	return db
}
