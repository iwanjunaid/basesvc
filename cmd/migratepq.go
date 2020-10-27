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
			sourceURL   = config.GetString("POSTGRES_MIGRATION_DIR")
			user        = config.GetString("POSTGRES_USER")
			pass        = config.GetString("POSTGRES_PASS")
			host        = config.GetString("POSTGRES_HOST")
			port        = config.GetInt("POSTGRES_PORT")
			db          = config.GetString("POSTGRES_DB")
			sslMode     = config.GetString("POSTGRES_SSL_MODE")
			searchPath  = config.GetString("POSTGRES_MIGRATION_SEARCH_PATH")
			databaseURL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&search_path=%s", user, pass,
				host, port, db, sslMode, searchPath)
		)

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
