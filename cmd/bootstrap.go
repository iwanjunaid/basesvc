package cmd

import (
	"database/sql"

	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
	log "github.com/sirupsen/logrus"
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
