package cmd

import (
	"database/sql"

	"github.com/evalphobia/logrus_sentry"
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
	log "github.com/sirupsen/logrus"
)

const (
	CfgMySql     = "database.mysql"
	CfgMongoDB   = "database.mysql"
	CfgRedis     = "database.redis"
	CfgSentryKey = "sentry.key"
)

var (
	db     *sql.DB
	logger *log.Logger
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
	if dsn := config.GetString(CfgSentryKey); len(dsn) > 0 {
		hook, err := logrus_sentry.NewSentryHook(dsn, []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
		})
		if err == nil {
			hook.StacktraceConfiguration.Enable = true
			l.Hooks.Add(hook)
		}
	}
	return l
}
