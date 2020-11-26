package rest

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/iwanjunaid/basesvc/internal/telemetry"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	newrelic "github.com/newrelic/go-agent"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/iwanjunaid/basesvc/config"

	"github.com/jmoiron/sqlx"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iwanjunaid/basesvc/adapter/controller"
	// _ "github.com/iwanjunaid/basesvc/docs"
	"github.com/iwanjunaid/basesvc/infrastructure/rest/group"
	logInternal "github.com/iwanjunaid/basesvc/internal/logger"
	"github.com/iwanjunaid/basesvc/registry"
	logger "github.com/sirupsen/logrus"
)

type RestImpl struct {
	port          int
	db            *sqlx.DB
	router        *fiber.App
	appController *controller.AppController
	log           *logger.Logger
}

// @title BaseSVC API
// @version 1.0
// @description This is a sample basesvc server api.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func NewRest(port int, logg *logger.Logger, db *sqlx.DB, mdb *mongo.Database, kp *kafka.Producer, rdb *redis.Ring, nra newrelic.Application) *RestImpl {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use("/swagger", swagger.Handler)
	app.Use(logInternal.RequestLogger(logg))
	app.Use(requestid.New())

	app.Use(telemetry.NewrelicMiddleware(nra, nil))
	// add graceful shutdown when interrupt signal detected
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		logg.Infof("server gracefully shutting down...")
		_ = app.Shutdown()
	}()

	registry := registry.NewRegistry(db, registry.NewMongoConn(mdb.Collection(config.GetString("database.collection"))), registry.NewKafkaProducer(kp), registry.NewRedisClient(rdb))
	appController := registry.NewAppController()

	r := &RestImpl{
		db:     db,
		port:   port,
		router: app,

		appController: &appController,
	}

	root := group.InitRoot(r)
	group.InitHealthCheck(r, root)
	v1 := group.InitV1(r, root)
	v2 := group.InitV2(r, root)

	group.InitAuthorV1(r, v1)
	group.InitAuthorV2(r, v2)

	return r
}

func (r *RestImpl) Serve() {
	if err := r.router.Listen(fmt.Sprintf(":%d", r.port)); err != nil {
		panic(err)
	}
}

func (r *RestImpl) GetRouter() *fiber.App {
	return r.router
}

func (r *RestImpl) GetAppController() *controller.AppController {
	return r.appController
}
