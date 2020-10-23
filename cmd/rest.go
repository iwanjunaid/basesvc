package cmd

import (
	"database/sql"
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
	rest "github.com/iwanjunaid/basesvc/infrastructure/rest"
	"github.com/iwanjunaid/basesvc/shared/logger"
	"github.com/spf13/cobra"
	"log"
)

var (
	db *sql.DB
	cfg *config.Config

	restCommand = &cobra.Command{
		Use: "api",
		PreRun: func(cmd *cobra.Command, args []string) {
			// load configuration
			cfg = config.GetConfig()

			// database configuration
			db = datastore.NewDB(cfg)

			// logging configuration
			logConfig := logger.Configuration{
				EnableConsole:     cfg.Logger.Console.Enable,
				ConsoleJSONFormat: cfg.Logger.Console.JSON,
				ConsoleLevel:      cfg.Logger.Console.Level,
				EnableFile:        cfg.Logger.File.Enable,
				FileJSONFormat:    cfg.Logger.File.JSON,
				FileLevel:         cfg.Logger.File.Level,
				FileLocation:      cfg.Logger.File.Path,
			}

			if err := logger.NewLogger(logConfig, logger.InstanceZapLogger); err != nil {
				log.Fatalf("could not instantiate log %v", err)
			}

			defer logger.WithFields(logger.Fields{"component": "api_command"}).Infof("PreRun done")
		},
		Run: func(cmd *cobra.Command, args []string) {
			defer logger.WithFields(logger.Fields{"component": "api_command"}).Infof("Run done")
			rest.NewRest(cfg.Server.Rest.Port, db).Serve()
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			defer db.Close()
		},
	}
)


