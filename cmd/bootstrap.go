package cmd

import (
	"database/sql"

	newrelic "github.com/newrelic/go-agent"
	"github.com/pkg/errors"

	"github.com/evalphobia/logrus_sentry"
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
	"github.com/newrelic/go-agent/_integrations/nrlogrus"
	log "github.com/sirupsen/logrus"
)

const (
	CfgMySql         = "database.mysql"
	CfgMongoDB       = "database.mysql"
	CfgRedis         = "database.redis"
	CfgSentryKey     = "sentry.key"
	CfgNewRelicKey   = "newrelic.key"
	CfgNewRelicDebug = "newrelic.debug"
	TelemteryID      = "newrelic.id"
)

var (
	db        *sql.DB
	logger    *log.Logger
	telemetry newrelic.Application
)

func init() {
	config.Configure()
	db = InitDB()
	logger = InitLogger()
	telemetry = NewTelemetry(logger)
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

func NewTelemetry(l *log.Logger) newrelic.Application {
	key := config.GetString(CfgNewRelicKey)
	e := l.WithField("component", "newrelic")
	if len(key) == 0 {
		e.Warnf("configuration %s is not defined", CfgNewRelicKey)
		return nil
	}
	conf := newrelic.NewConfig(config.GetString(TelemteryID), key)
	conf.Logger = nrlogrus.StandardLogger()
	if isDebug := config.GetBool(CfgNewRelicDebug); isDebug {
		l.SetLevel(log.DebugLevel)
	}
	app, err := newrelic.NewApplication(conf)
	if err != nil {
		e.Info(errors.Cause(err))
		return nil
	}
	return app
}
