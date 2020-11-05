package main

import (
	"io"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/pkg/action"
)

const chartHelp = `
This command consists of multiple subcommands to work with the chart cache.

The subcommands can be used to push, pull, tag, list, or remove Helm charts.
`

func newChartCmd(cfg *action.Configuration, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "chart",
		Short:             "push, pull, tag, or remove Helm charts",
		Long:              chartHelp,
		Hidden:            !FeatureGateOCI.IsEnabled(),
		PersistentPreRunE: checkOCIFeatureGate(),
	}
	cmd.AddCommand(
		newChartListCmd(cfg, out),
		newChartExportCmd(cfg, out),
		newChartPullCmd(cfg, out),
		newChartPushCmd(cfg, out),
		newChartRemoveCmd(cfg, out),
		newChartSaveCmd(cfg, out),
	)
	return cmd
}
