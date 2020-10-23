package cmd

import (
	"database/sql"

	"github.com/iwanjunaid/basesvc/config"
	log "github.com/sirupsen/logrus"

	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
)

var (
	logger *log.Logger
	db     *sql.DB
)

func init() {
	config.ReadConfig()
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
