package rest

import (
	"database/sql"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	mlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iwanjunaid/basesvc/adapter/controller"
	"github.com/iwanjunaid/basesvc/infrastructure/rest/group"
	"github.com/iwanjunaid/basesvc/registry"
	logger "github.com/sirupsen/logrus"
)

type RestImpl struct {
	port          string
	db            *sql.DB
	router        *fiber.App
	appController *controller.AppController
	log           *logger.Logger
}

func NewRest(port string, logg *logger.Logger, db *sql.DB) *RestImpl {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(mlogger.New())
	//app.Use(log.RequestLogger(logg))

	// add graceful shutdown when interrupt signal detected
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		logg.Infof("server gracefully shutting down...")
		_ = app.Shutdown()
	}()

	registry := registry.NewRegistry(db)
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
	if err := r.router.Listen(r.port); err != nil {
		panic(err)
	}
}

func (r *RestImpl) GetRouter() *fiber.App {
	return r.router
}

func (r *RestImpl) GetAppController() *controller.AppController {
	return r.appController
}
