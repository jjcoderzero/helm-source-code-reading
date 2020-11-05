package main

import (
	"io"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/cmd/helm/require"
	"helm.sh/helm/v3/pkg/action"
)

const chartExportDesc = `
Export a chart stored in local registry cache.

This will create a new directory with the name of
the chart, in a format that developers can modify
and check into source control if desired.
`

func newChartExportCmd(cfg *action.Configuration, out io.Writer) *cobra.Command {
	client := action.NewChartExport(cfg)

	cmd := &cobra.Command{
		Use:    "export [ref]",
		Short:  "export a chart to directory",
		Long:   chartExportDesc,
		Args:   require.MinimumNArgs(1),
		Hidden: !FeatureGateOCI.IsEnabled(),
		RunE: func(cmd *cobra.Command, args []string) error {
			ref := args[0]
			return client.Run(out, ref)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&client.Destination, "destination", "d", ".", "location to write the chart.")

	return cmd
}
