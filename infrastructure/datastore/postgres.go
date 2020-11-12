package datastore

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/newrelic/go-agent/_integrations/nrpq"
)

func PostgresConn(dsn string) *sqlx.DB {
	db, err := sqlx.Open("nrpostgres", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
