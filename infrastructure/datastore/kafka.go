package datastore

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaConsumer(host, groupID string, protocol string, mechanisms string, key string, secret string) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  host,
		"sasl.mechanisms":    mechanisms,
		"security.protocol":  protocol,
		"sasl.username":      key,
		"sasl.password":      secret,
		"group.id":           groupID,
		"session.timeout.ms": 6000,
	})
	if err != nil {
		panic(err)
	}

	return c
}

func NewKafkaProducer(broker string, protocol string, mechanisms string, key string, secret string) *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"sasl.mechanisms":   mechanisms,
		"security.protocol": protocol,
		"sasl.username":     key,
		"sasl.password":     secret,
	})
	if err != nil {
		panic(err)
	}
	return p
}
