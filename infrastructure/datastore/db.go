package datastore

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iwanjunaid/basesvc/config"
)

func NewDB() *sql.DB {
	DBMS := `mysql`
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.C.Database.User, config.C.Database.Password, config.C.Database.Host, config.C.Database.Port, config.C.Database.DBName)
	db, err := sql.Open(DBMS, conn)

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
