package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migratePqCmd = &cobra.Command{
	Use: "migratepq",
	Run: func(cmd *cobra.Command, args []string) {
		mig, err := migrate.New(
			"./db/migrations/postgres",
			"postgres://postgres:123456@127.0.0.1:5432/postgres?sslmode=disable&search_path=public")

		if err != nil {
			logger.Fatal(err)
		}

		if len(args) > 1 {
			switch args[0] {
			case "up":
				if err := mig.Up(); err != nil {
					logger.Fatal(err)
				}
			case "down":
				if err := mig.Down(); err != nil {
					logger.Fatal(err)
				}
			}
		}
	},
}
