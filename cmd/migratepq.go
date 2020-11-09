package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/iwanjunaid/basesvc/config"
	"github.com/spf13/cobra"
)

var migratePqCmd = &cobra.Command{
	Use: "migratepq",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			sourceURL  = config.GetString("migration.postgres.dir")
			searchPath = config.GetString("migration.postgres.search_path")
			user       = config.GetString("database.postgres.write.user")
			pass       = config.GetString("database.postgres.write.pass")
			host       = config.GetString("database.postgres.write.host")
			port       = config.GetInt("database.postgres.write.port")
			db         = config.GetString("database.postgres.write.db")
			sslMode    = config.GetString("database.postgres.write.sslmode")
		)
		databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&search_path=%s",
			user, pass, host, port, db, sslMode, searchPath)

		mig, err := migrate.New(
			sourceURL,
			databaseURL)

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
