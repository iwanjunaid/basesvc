package kafka

import (
	"context"
	"errors"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/basesvc/config"
)

type InternalKafkaImpl struct {
	kp *kafka.Producer
}

type InternalKafka interface {
	Publish(ctx context.Context, key, message []byte) (err error)
}

func (author *InternalKafkaImpl) Publish(ctx context.Context, key, message []byte) (err error) {
	var (
		topics       = config.GetStringSlice("kafka.topics")
		topic        string
		deliveryChan = make(chan kafka.Event)
	)

	if len(topics) > 0 {
		topic = topics[0]
	}
	err = author.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
		Key:            key,
	}, deliveryChan)

	e := <-deliveryChan

	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return errors.New(fmt.Sprintf("Delivery failed: %v\n", m.TopicPartition.Error))
	}
	close(deliveryChan)
	return
}

func NewInternalKafkaImpl(kp *kafka.Producer) InternalKafka {
	return &InternalKafkaImpl{
		kp: kp,
	}
}
