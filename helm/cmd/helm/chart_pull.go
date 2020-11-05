package main

import (
	"io"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/cmd/helm/require"
	"helm.sh/helm/v3/pkg/action"
)

const chartPullDesc = `
Download a chart from a remote registry.

This will store the chart in the local registry cache to be used later.
`

func newChartPullCmd(cfg *action.Configuration, out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:    "pull [ref]",
		Short:  "pull a chart from remote",
		Long:   chartPullDesc,
		Args:   require.MinimumNArgs(1),
		Hidden: !FeatureGateOCI.IsEnabled(),
		RunE: func(cmd *cobra.Command, args []string) error {
			ref := args[0]
			return action.NewChartPull(cfg).Run(out, ref)
		},
	}
}
