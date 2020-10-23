package cmd

import (
	"github.com/iwanjunaid/basesvc/shared/logger"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of basesvc",
	Long:  "Print the version number of basesvc",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Infof ("v0.1.0")
	},
}
