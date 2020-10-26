package cmd

import (
	"database/sql"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/basesvc/config"
	log "github.com/sirupsen/logrus"

	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
)

const (
	CfgMySql      = "database.mysql"
	CfgMongoDB    = "database.mysql"
	CfgRedis      = "database.redis"
	CfgKafkaGroup = "kafka.groupID"
	CfgKafkaHost  = "kafka.host"
	CfgKafkaTopic = "kafka.topic"
)

var (
	logger *log.Logger
	db     *sql.DB
	kc     *kafka.Consumer
	kp     *kafka.Producer
)

func init() {
	config.Configure()
	db = InitDB()
	logger = InitLogger()
	kc = InitKafkaConsumer()
	kp = InitKafkaProducer()
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

func InitKafkaConsumer() *kafka.Consumer {
	return datastore.NewKafkaConsumer(config.GetString(CfgKafkaHost), config.GetString(CfgKafkaGroup))
}

func InitKafkaProducer() *kafka.Producer {
	return datastore.NewKafkaProducer(config.GetString(CfgKafkaHost))
}
