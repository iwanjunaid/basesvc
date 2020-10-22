package router

import (
	"database/sql"
	"log"

	"github.com/iwanjunaid/basesvc/adapter/controller"

	"github.com/iwanjunaid/basesvc/config"

	"github.com/iwanjunaid/basesvc/registry"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Rest struct {
	port   string
	db     *sql.DB
	router *fiber.App
}

func NewRest(port string, db *sql.DB) *Rest {
	r := &Rest{
		db:   db,
		port: port,
	}
	return r
}

func (r *Rest) Serve() {
	r.setup()
	if err := r.router.Listen(config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}

func (r *Rest) setup() {
	r.router = r.InitRouter()
}

func (r *Rest) InitRouter() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	registry := registry.NewRegistry(r.db)

	c := registry.NewAppController()
	v1 := app.Group("/v1")
	author := v1.Group("/author")
	r.AuthorRoutes(author, c)

	return app
}

func (r *Rest) AuthorRoutes(author fiber.Router, c controller.AppController) {
	author.Get("/", func(ctx *fiber.Ctx) error {
		return c.Author.GetAuthors(ctx)
	})
	author.Get("/:id", func(ctx *fiber.Ctx) error {
		return c.Author.GetAuthors(ctx)
	})
	author.Post("/", func(ctx *fiber.Ctx) error {
		return c.Author.GetAuthors(ctx)
	})
	author.Put("/:id", func(ctx *fiber.Ctx) error {
		return c.Author.GetAuthors(ctx)
	})
}
