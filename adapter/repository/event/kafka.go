package event

import (
	"context"
	"errors"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
)

type AuthorEventRepositoryImpl struct {
	kp *kafka.Producer
}

func (author *AuthorEventRepositoryImpl) Publish(ctx context.Context, topic string, key, message []byte) (err error) {
	deliveryChan := make(chan kafka.Event)

	err = author.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
		Key:            key,
	}, deliveryChan)

	e := <-deliveryChan

	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return errors.New(fmt.Sprintf("Delivery failed: %v\n", m.TopicPartition.Error))
	} else {
		return errors.New(fmt.Sprintf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset))
	}

	close(deliveryChan)
	return
}

func NewAuthorEventRepository(kp *kafka.Producer) repository.AuthorEventRepository {
	return &AuthorEventRepositoryImpl{
		kp: kp,
	}
}
