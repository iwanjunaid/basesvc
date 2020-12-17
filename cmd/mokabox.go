package cmd

import (
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/mokabox"
	"github.com/spf13/cobra"
)

var mokaboxCommand = &cobra.Command{
	Use: "mokabox",
	PreRun: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("component", "consumer_command").Println("PreRun done")
	},
	Run: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("component", "consumer_command").Println("Run done")
		mokabox.NewMokaBox(kc, db, mdb).Listen(config.GetStringSlice("kafka.topics"))
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer db.Close()
	},
}
