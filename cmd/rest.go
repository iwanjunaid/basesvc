package cmd

import (
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/rest"
	"github.com/iwanjunaid/basesvc/internal/logger"
	"github.com/spf13/cobra"
)

var restCommand = &cobra.Command{
	Use: "api",
	PreRun: func(cmd *cobra.Command, args []string) {
		defer logger.WithFields(logger.Fields{"component": "api_command"}).Infof("PreRun done")
	},
	Run: func(cmd *cobra.Command, args []string) {
		defer logger.WithFields(logger.Fields{"component": "api_command"}).Infof("Run done")
		rest.NewRest(config.GetString("host.address"), db).Serve()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer db.Close()
	},
}
