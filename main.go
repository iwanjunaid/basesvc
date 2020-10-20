package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/labstack/echo"

	"github.com/gofiber/fiber/v2"

	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
	"github.com/iwanjunaid/basesvc/infrastructure/router"
	"github.com/iwanjunaid/basesvc/registry"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()

	defer db.Close()

	r := registry.NewRegistry(db)

	app := fiber.New()

	router.NewRouter(app, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := app.Listen(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
