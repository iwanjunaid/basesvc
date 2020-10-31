package cmd

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/evalphobia/logrus_sentry"
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
	log "github.com/sirupsen/logrus"
)

const (
	CfgMySql      = "database.mysql"
	CfgRedis      = "database.redis"
	CfgKafkaGroup = "kafka.groupID"
	CfgKafkaHost  = "kafka.host"
	CfgKafkaTopic = "kafka.topic"
	CfgMongoURI   = "database.mongo.uri"
	CfgMongoDB    = "database.mongo.db"
	CfgSentryKey  = "sentry.key"
)

var (
	logger *log.Logger
	db     *sql.DB
	kc     *kafka.Consumer
	kp     *kafka.Producer
	mdb    *mongo.Database
)

func init() {
	config.Configure()
	db = InitDB()
	logger = InitLogger()
	kc = InitKafkaConsumer()
	kp = InitKafkaProducer()
	mdb = InitMongoConnect()
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

func InitKafkaConsumer() *kafka.Consumer {
	return datastore.NewKafkaConsumer(config.GetString(CfgKafkaHost), config.GetString(CfgKafkaGroup))
}

func InitKafkaProducer() *kafka.Producer {
	return datastore.NewKafkaProducer(config.GetString(CfgKafkaHost))
}

func InitMongoConnect() *mongo.Database {
	return datastore.MongoMustConnect(config.GetString(CfgMongoURI), config.GetString(CfgMongoDB))
}
