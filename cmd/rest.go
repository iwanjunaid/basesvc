package cmd

import (
	"os"

	"github.com/iwanjunaid/basesvc/config"
	"github.com/iwanjunaid/basesvc/infrastructure/rest"
	nr "github.com/newrelic/opentelemetry-exporter-go/newrelic"

	"github.com/spf13/cobra"
)

var restCommand = &cobra.Command{
	Use: "api",
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.WithField("component", "api_command").Infof("PreRun done")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("NEW_RELIC_API_KEY") != "" {
			controller, err := nr.InstallNewPipeline("BaseSvc")
			if err != nil {
				panic(err)
			}
			defer controller.Stop()
		}

		rest.NewRest(config.GetInt("rest.port"), logger, db, mdb, kp, telemetry).Serve()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer db.Close()
	},
}
