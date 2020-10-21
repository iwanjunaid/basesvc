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
	//if dsn := config.(CfgSentryKey); len(dsn) > 0 {
	//	hook, err := logrus_sentry.NewSentryHook(dsn, []log.Level{
	//		log.PanicLevel,
	//		log.FatalLevel,
	//		log.ErrorLevel,
	//	})
	//	if err == nil {
	//		hook.StacktraceConfiguration.Enable = true
	//		l.Hooks.Add(hook)
	//	}
	//}
	return l
}
