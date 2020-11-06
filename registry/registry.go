package registry

import (
	"github.com/jmoiron/sqlx"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/basesvc/adapter/controller"
)

type registry struct {
	db  *sqlx.DB
	mdb *mongo.Database
	kP  *kafka.Producer
}

type Registry interface {
	NewAppController() controller.AppController
}

type Option func(*registry)

func NewRegistry(db *sqlx.DB, option ...Option) Registry {
	r := &registry{db: db}
	for _, o := range option {
		o(r)
	}
	return r
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Author: r.NewAuthorController(),
	}
}

func NewMongoConn(mdb *mongo.Database) Option {
	return func(i *registry) {
		i.mdb = mdb
	}
}

func NewKafkaProducer(kp *kafka.Producer) Option {
	return func(i *registry) {
		i.kP = kp
	}
}
