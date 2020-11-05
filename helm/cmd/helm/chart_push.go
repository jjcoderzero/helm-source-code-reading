package main

import (
	"io"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/cmd/helm/require"
	"helm.sh/helm/v3/pkg/action"
)

const chartPushDesc = `
Upload a chart to a remote registry.

Note: the ref must already exist in the local registry cache.

Must first run "helm chart save" or "helm chart pull".
`

func newChartPushCmd(cfg *action.Configuration, out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:    "push [ref]",
		Short:  "push a chart to remote",
		Long:   chartPushDesc,
		Args:   require.MinimumNArgs(1),
		Hidden: !FeatureGateOCI.IsEnabled(),
		RunE: func(cmd *cobra.Command, args []string) error {
			ref := args[0]
			return action.NewChartPush(cfg).Run(out, ref)
		},
	}
}
