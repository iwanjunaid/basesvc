package cmd

import (
	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/consumer"
	"github.com/spf13/cobra"
)

var consumerCommand = &cobra.Command{
	Use: "consumer",
	PreRun: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("component", "consumer_command").Println("PreRun done")
	},
	Run: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("component", "consumer_command").Println("Run done")
		consumer.NewConsumer(kc, db).Listen(config.GetString("kafka.topic"))
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer db.Close()
	},
}
