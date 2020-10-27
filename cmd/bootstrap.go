package cmd

import (
	"database/sql"
	"fmt"

	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
	log "github.com/iwanjunaid/basesvc/internal/logger"
)

const (
	CfgMySql   = "database.mysql"
	CfgMongoDB = "database.mysql"
	CfgRedis   = "database.redis"
)

var db     *sql.DB

func init() {
	config.Configure()
	db = InitDB()
	InitLogger()
}

func InitDB() (db *sql.DB) {
	db = datastore.NewDB("mysql", config.GetString(CfgMySql))
	return
}

func InitLogger() {
	logger := log.Configuration{
		EnableConsole:     config.GetBool("logger.console.enable"),
		ConsoleJSONFormat: config.GetBool("logger.console.json"),
		ConsoleLevel:      config.GetString("logger.console.level"),
		EnableFile:        config.GetBool("logger.file.enable"),
		FileJSONFormat:    config.GetBool("logger.file.json"),
		FileLevel:         config.GetString("logger.file.level"),
		FileLocation:      config.GetString("logger.file.path"),
	}

	if err := log.NewLogger(logger, log.InstanceZapLogger); err != nil {
		panic(fmt.Sprintf("could not instantiate log %v", err))
	}
}
