package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/iwanjunaid/basesvc/cmd"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/labstack/echo"
)

func main() {
	cmd.Run()
}
