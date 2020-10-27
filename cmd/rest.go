package cmd

import (
	"fmt"

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
		var host = config.GetString("REST_HOST")
		var port = config.GetInt("REST_PORT")
		var address = fmt.Sprintf("%s:%d", host, port)

		defer logger.WithField("component", "api_command").Println("Run done")
		rest.NewRest(address, db).Serve()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer db.Close()
	},
}
