package cmd

import (
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "basesvc",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Infof("root command")
	},
}

func Run() {
	rootCommand.AddCommand(restCommand)
	rootCommand.AddCommand(migratePqCmd)
	rootCommand.AddCommand(migrateMongoCmd)
	rootCommand.AddCommand(consumerCommand)
	rootCommand.AddCommand(mokaboxCommand)

	if err := rootCommand.Execute(); err != nil {
		logger.Panicf("%v", err)
	}
}
