package main

import (
	"io"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/cmd/helm/require"
	"helm.sh/helm/v3/pkg/action"
)

const chartRemoveDesc = `
Remove a chart from the local registry cache.

Note: the chart content will still exist in the cache,
but it will no longer appear in "helm chart list".

To remove all unlinked content, please run "helm chart prune". (TODO)
`

func newChartRemoveCmd(cfg *action.Configuration, out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:     "remove [ref]",
		Aliases: []string{"rm"},
		Short:   "remove a chart",
		Long:    chartRemoveDesc,
		Args:    require.MinimumNArgs(1),
		Hidden:  !FeatureGateOCI.IsEnabled(),
		RunE: func(cmd *cobra.Command, args []string) error {
			ref := args[0]
			return action.NewChartRemove(cfg).Run(out, ref)
		},
	}
}
