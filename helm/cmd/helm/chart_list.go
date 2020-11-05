package main

import (
	"io"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/pkg/action"
)

const chartListDesc = `
List all charts in the local registry cache.

Charts are sorted by ref name, alphabetically.
`

func newChartListCmd(cfg *action.Configuration, out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list all saved charts",
		Long:    chartListDesc,
		Hidden:  !FeatureGateOCI.IsEnabled(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return action.NewChartList(cfg).Run(out)
		},
	}
}
