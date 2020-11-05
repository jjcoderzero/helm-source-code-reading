package main

import (
	"io"
	"path/filepath"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/cmd/helm/require"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
)

const chartSaveDesc = `
Store a copy of chart in local registry cache.

Note: modifying the chart after this operation will
not change the item as it exists in the cache.
`

func newChartSaveCmd(cfg *action.Configuration, out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:    "save [path] [ref]",
		Short:  "save a chart directory",
		Long:   chartSaveDesc,
		Args:   require.MinimumNArgs(2),
		Hidden: !FeatureGateOCI.IsEnabled(),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			ref := args[1]

			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			ch, err := loader.Load(path)
			if err != nil {
				return err
			}

			return action.NewChartSave(cfg).Run(out, ch, ref)
		},
	}
}
