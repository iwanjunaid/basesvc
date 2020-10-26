package rest

import (
	"database/sql"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iwanjunaid/basesvc/adapter/controller"
	"github.com/iwanjunaid/basesvc/infrastructure/rest/group"
	log "github.com/iwanjunaid/basesvc/internal/logger"
	"github.com/iwanjunaid/basesvc/registry"
)

type RestImpl struct {
	port          string
	db            *sql.DB
	router        *fiber.App
	appController *controller.AppController
}

func NewRest(port string, db *sql.DB) *RestImpl {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	// add graceful shutdown when interrupt signal detected
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Infof ("server gracefully shutting down...")
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
