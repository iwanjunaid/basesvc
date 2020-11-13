package cmd

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	newrelic "github.com/newrelic/go-agent"
	"github.com/pkg/errors"

	"github.com/evalphobia/logrus_sentry"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
	"github.com/newrelic/go-agent/_integrations/nrlogrus"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/basesvc/config"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var (
	CfgMySql         = "database.mysql"
	CfgRedis         = "database.redis"
	CfgKafkaGroup    = "kafka.group_id"
	CfgKafkaHost     = "kafka.host"
	CfgKafkaTopic    = "kafka.topics"
	CfgNewRelicKey   = "newrelic.key"
	CfgNewRelicDebug = "newrelic.debug"
	CfgMongoURI      = "database.mongo.uri"
	CfgMongoDB       = "database.mongo.db"
	CfgSentryKey     = "sentry.key"
	TelemetryID      = "newrelic.id"
)

var (
	logger    *log.Logger
	db        *sqlx.DB
	kc        *kafka.Consumer
	kp        *kafka.Producer
	mdb       *mongo.Database
	telemetry newrelic.Application
)

func init() {
	c := config.Configure()

	// hot reload on config change...
	go func() {
		c.WatchConfig()
		c.OnConfigChange(func(e fsnotify.Event) {
			log.Printf("config file changed %v", e.Name)
		})
	}()

	db = InitPostgresDB()
	logger = InitLogger()
	telemetry = NewTelemetry(logger)
	kc = InitKafkaConsumer()
	kp = InitKafkaProducer()
	mdb = InitMongoConnect()
}

func InitPostgresDB() (db *sqlx.DB) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.GetString("database.postgres.write.host"), config.GetInt("database.postgres.write.port"),
		config.GetString("database.postgres.write.user"), config.GetString("database.postgres.write.pass"),
		config.GetString("database.postgres.write.db"), config.GetString("database.postgres.write.sslmode"))
	db = datastore.PostgresConn(dsn)
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
	conf := newrelic.NewConfig(config.GetString(TelemetryID), key)
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

func InitKafkaConsumer() *kafka.Consumer {
	return datastore.NewKafkaConsumer(config.GetString(CfgKafkaHost), config.GetString(CfgKafkaGroup))
}

func InitKafkaProducer() *kafka.Producer {
	return datastore.NewKafkaProducer(config.GetString(CfgKafkaHost))
}

func InitMongoConnect() *mongo.Database {
	return datastore.MongoMustConnect(config.GetString(CfgMongoURI), config.GetString(CfgMongoDB))
}
