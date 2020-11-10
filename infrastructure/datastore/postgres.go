package datastore

import "github.com/jmoiron/sqlx"

func PostgresConn(dsn string) *sqlx.DB {
	db := sqlx.MustConnect("postgres", dsn)
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
