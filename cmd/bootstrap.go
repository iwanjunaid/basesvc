package cmd

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/evalphobia/logrus_sentry"
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	CfgMySql      = "database.mysql"
	CfgRedis      = "database.redis"
	CfgKafkaGroup = "kafka.group_id"
	CfgKafkaHost  = "kafka.host"
	CfgKafkaTopic = "kafka.topic"
	CfgMongoURI   = "database.mongo.uri"
	CfgMongoDB    = "database.mongo.db"
	CfgSentryKey  = "sentry.key"
)

var (
	logger *log.Logger
	db     *sqlx.DB
	kc     *kafka.Consumer
	kp     *kafka.Producer
	mdb    *mongo.Database
)

func init() {
	config.Configure()
	db = InitPostgresDB()
	logger = InitLogger()
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

func InitKafkaConsumer() *kafka.Consumer {
	return datastore.NewKafkaConsumer(config.GetString(CfgKafkaHost), config.GetString(CfgKafkaGroup))
}

func InitKafkaProducer() *kafka.Producer {
	return datastore.NewKafkaProducer(config.GetString(CfgKafkaHost))
}

func InitMongoConnect() *mongo.Database {
	return datastore.MongoMustConnect(config.GetString(CfgMongoURI), config.GetString(CfgMongoDB))
}
