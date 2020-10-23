package cmd

import (
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/rest"
	"github.com/spf13/cobra"
)

var restCommand = &cobra.Command{
	Use: "api",
	PreRun: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("component", "api_command").Println("PreRun done")
	},
	Run: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("component", "api_command").Println("Run done")
		rest.NewRest(config.GetString("host.address"), db).Serve()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer db.Close()
	},
}
