package cmd

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
)

var (
	logger *log.Logger
	db     *sql.DB
)

func init() {
	db = InitDB()
	logger = InitLogger()
}

func InitDB() (db *sql.DB) {
	db = datastore.NewDB()

	return
}

func InitLogger() *log.Logger {
	log.SetFormatter(&log.JSONFormatter{})
	l := log.StandardLogger()
	return l
}
