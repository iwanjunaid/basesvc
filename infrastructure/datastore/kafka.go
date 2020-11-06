package datastore

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//c, err := c.k.NewConsumer(&kafka.ConfigMap{
//	"bootstrap.servers": broker,
//	// Avoid connecting to IPv6 brokers:
//	// This is needed for the ErrAllBrokersDown show-case below
//	// when using localhost brokers on OSX, since the OSX resolver
//	// will return the IPv6 addresses first.
//	// You typically don't need to specify this configuration property.
//	"broker.address.family": "v4",
//	"group.id":              group,
//	"session.timeout.ms":    6000,
//	"auto.offset.reset":     "earliest"})

func NewKafkaConsumer(host, groupID string) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     host,
		"broker.address.family": "v4",
		"group.id":              groupID,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest",
	})
	fmt.Println(host)
	fmt.Println(groupID)

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
