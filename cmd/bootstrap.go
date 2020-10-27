package cmd

import (
	"database/sql"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
	log "github.com/iwanjunaid/basesvc/internal/logger"
)

const (
	CfgMySql      = "database.mysql"
	CfgRedis      = "database.redis"
	CfgKafkaGroup = "kafka.groupID"
	CfgKafkaHost  = "kafka.host"
	CfgKafkaTopic = "kafka.topic"
	CfgMongoURI   = "database.mongo.uri"
	CfgMongoDB    = "database.mongo.db"
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

func InitKafkaConsumer() *kafka.Consumer {
	return datastore.NewKafkaConsumer(config.GetString(CfgKafkaHost), config.GetString(CfgKafkaGroup))
}

func InitKafkaProducer() *kafka.Producer {
	return datastore.NewKafkaProducer(config.GetString(CfgKafkaHost))
}

func InitMongoConnect() *mongo.Database {
	return datastore.MongoMustConnect(config.GetString(CfgMongoURI), config.GetString(CfgMongoDB))
}
