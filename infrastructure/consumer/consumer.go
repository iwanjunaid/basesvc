package consumer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/iwanjunaid/basesvc/config"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/jmoiron/sqlx"

	"github.com/iwanjunaid/basesvc/adapter/controller"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/registry"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsumerImpl struct {
	kc            *kafka.Consumer
	appController *controller.AppController
}

func NewConsumer(kc *kafka.Consumer, db *sqlx.DB, mdb *mongo.Database) *ConsumerImpl {
	registry := registry.NewRegistry(db, registry.NewMongoConn(
		mdb.Collection(config.GetString("database.mongo.collection"))))
	appController := registry.NewAppController()

	return &ConsumerImpl{
		kc:            kc,
		appController: &appController,
	}
}

func (c *ConsumerImpl) Listen(topic []string) {
	err := c.kc.SubscribeTopics(topic, nil)
	if err != nil {
		panic(err)
	}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	run := true
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.kc.Poll(100)
			if ev == nil {
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				var author *model.Author
				if err := json.Unmarshal(e.Value, &author); err != nil {
					fmt.Println(err.Error())
				}
				if err := c.appController.Author.InsertDocument(author); err != nil {
					fmt.Println(err.Error())
				}
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				// Errors should generally be considered as informational, the client will try to automatically recover
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.kc.Close()
}
