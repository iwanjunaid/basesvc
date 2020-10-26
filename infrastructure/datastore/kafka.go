package datastore

import "github.com/confluentinc/confluent-kafka-go/kafka"

func NewKafkaConsumer(host, groupID string) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": host,
		"group.id":          groupID,
	})

	if err != nil {
		panic(err)
	}

	return c
}

func NewKafkaProducer(broker string) *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		panic(err)
	}
	return p
}
