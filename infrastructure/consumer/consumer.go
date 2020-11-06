package consumer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"

	"github.com/iwanjunaid/basesvc/adapter/controller"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/registry"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsumerImpl struct {
	kc            *kafka.Consumer
	appController controller.AppController
}

func NewConsumer(kc *kafka.Consumer, db *sqlx.DB) *ConsumerImpl {
	registry := registry.NewRegistry(db)
	appController := registry.NewAppController()

	return &ConsumerImpl{
		kc:            kc,
		appController: appController,
	}
}

func (c *ConsumerImpl) Listen(topic []string) {
	//fmt.Println(c.kc.)
	//fmt.Println(topic)
	err := c.kc.SubscribeTopics(topic, nil)
	if err != nil {
		panic(err)
	}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	run := true
	fmt.Println("blah")
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
			c.kc.Events()
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.kc.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.kc.Unassign()
			case *kafka.Message:
				fmt.Println("test")
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				var author *model.Author
				if err := json.Unmarshal(e.Value, &author); err != nil {
					fmt.Println("invalid")
				}
				c.appController.Author.InsertAuthor(author)
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				// Errors should generally be considered as informational, the client will try to automatically recover
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
		}
		//case ev := <-c.kc.Events():
		//	fmt.Println("Test")
		//
		//
		//}
	}

	fmt.Printf("Closing consumer\n")
	c.kc.Close()
}
