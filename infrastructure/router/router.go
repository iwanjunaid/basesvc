package router

import (
	"database/sql"
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/iwanjunaid/basesvc/config"
	_ "github.com/iwanjunaid/basesvc/docs"

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

// @title BaseSVC API
// @version 1.0
// @description This is a sample basesvc server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func (r *Rest) InitRouter() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Use("/swagger", swagger.Handler)
	app.Use(recover.New())
  
	registry := registry.NewRegistry(r.db)

	c := registry.NewAppController()
	v1 := app.Group("/v1")
	author := v1.Group("/author")
	author.Get("/", func(ctx *fiber.Ctx) error {
		return c.Author.GetAuthors(ctx)
	})
	app.Get("/:id", func(ctx *fiber.Ctx) error {
		return c.Author.GetAuthors(ctx)
	})
	app.Post("/", func(ctx *fiber.Ctx) error {
		return c.Author.GetAuthors(ctx)
	})
	app.Put("/:id", func(ctx *fiber.Ctx) error {
		return c.Author.GetAuthors(ctx)
	})

	return app
}
