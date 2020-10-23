package datastore

import (
	"fmt"
	"github.com/iwanjunaid/basesvc/config"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(config *config.Config) *sql.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Database.Master.Mysql.User, config.Database.Master.Mysql.Password, config.Database.Master.Mysql.Host, config.Database.Master.Mysql.Port, config.Database.Master.Mysql.DBName)
	fmt.Println(conn)

	db, err := sql.Open(config.Database.Master.Mysql.Adapter, conn)

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
