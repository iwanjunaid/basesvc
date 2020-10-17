package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"

	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
	"github.com/iwanjunaid/basesvc/infrastructure/router"
	"github.com/iwanjunaid/basesvc/registry"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
