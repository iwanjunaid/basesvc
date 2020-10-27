package datastore

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(driverName, dsn string) *sql.DB {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
