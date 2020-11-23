package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/iwanjunaid/basesvc/config"
	"github.com/spf13/cobra"
)

var migrateMongoCmd = &cobra.Command{
	Use: "migratemongo",
	Run: func(cmd *cobra.Command, args []string) {
		mongoURI := config.GetString("database.mongo.uri")
		sourceURL := config.GetString("migration.mongo.dir")

		mig, err := migrate.New(
			sourceURL,
			mongoURI)

		if err != nil {
			logger.Fatal(err)
		}

		if len(args) >= 1 {
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
