package rest

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/mongo"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iwanjunaid/basesvc/adapter/controller"
	_ "github.com/iwanjunaid/basesvc/docs"
	"github.com/iwanjunaid/basesvc/infrastructure/rest/group"
	logInternal "github.com/iwanjunaid/basesvc/internal/logger"
	"github.com/iwanjunaid/basesvc/registry"
	logger "github.com/sirupsen/logrus"
)

type RestImpl struct {
	port          int
	db            *sql.DB
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
func NewRest(port int, logg *logger.Logger, db *sql.DB, mdb *mongo.Database, kp *kafka.Producer) *RestImpl {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use("/swagger", swagger.Handler)
	app.Use(logInternal.RequestLogger(logg))

	// add graceful shutdown when interrupt signal detected
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		logg.Infof("server gracefully shutting down...")
		_ = app.Shutdown()
	}()

	registry := registry.NewRegistry(db, registry.NewMongoConn(mdb), registry.NewKafkaProducer(kp))
	appController := registry.NewAppController()

	r := &RestImpl{
		db:            db,
		port:          port,
		router:        app,
		appController: &appController,
	}

	group.InitRoot(r)
	group.InitV1(r)
	group.InitV2(r)
	group.InitAuthorV1(r)
	group.InitAuthorV2(r)

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
