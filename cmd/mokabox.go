package cmd

import (
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/datastore"
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
		mClient, err := datastore.MongoConnect(config.GetString(CfgMongoURI))
		if err != nil {
			panic(err)
		}
		mokabox.NewMokaBox(mClient).Listen(config.GetStringSlice("kafka.topics"))
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer db.Close()
	},
}
