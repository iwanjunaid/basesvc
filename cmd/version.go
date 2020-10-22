package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of basesvc",
	Long:  "Print the version number of basesvc",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("v0.1.0")
	},
}
