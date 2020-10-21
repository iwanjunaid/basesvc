package cmd

import (
	"github.com/iwanjunaid/basesvc/config"
	rest "github.com/iwanjunaid/basesvc/infrastructure/router"
	"github.com/spf13/cobra"
)

var restCommand = &cobra.Command{
	Use: "api",
	PreRun: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("commponent", "api_command").Println("PreRun done")
	},
	Run: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("component", "api_command").Println("Run done")
		rest.NewRest(config.C.Server.Address, db).Serve()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer db.Close()
	},
}
