package cmd

import (
	"database/sql"

	"github.com/iwanjunaid/basesvc/config"
	log "github.com/sirupsen/logrus"

	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
)

const (
	CfgMySql   = "database.mysql"
	CfgMongoDB = "database.mysql"
	CfgRedis   = "database.redis"
)

var (
	logger *log.Logger
	db     *sql.DB
)

func init() {
	config.Configure()
	db = InitDB()
	logger = InitLogger()
}

func InitDB() (db *sql.DB) {
	db = datastore.NewDB("mysql", config.GetString(CfgMySql))
	return
}

func InitLogger() *log.Logger {
	log.SetFormatter(&log.JSONFormatter{})
	l := log.StandardLogger()
	return l
}
