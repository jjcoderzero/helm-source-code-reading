package main

import (
	"io"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/cmd/helm/require"
	"helm.sh/helm/v3/pkg/action"
)

var getHelp = `
This command consists of multiple subcommands which can be used to
get extended information about the release, including:

- The values used to generate the release
- The generated manifest file
- The notes provided by the chart of the release
- The hooks associated with the release
`

func newGetCmd(cfg *action.Configuration, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get",
		Short:             "download extended information of a named release",
		Long:              getHelp,
		Args:              require.NoArgs,
		ValidArgsFunction: noCompletions, // Disable file completion
	}

	cmd.AddCommand(newGetAllCmd(cfg, out))
	cmd.AddCommand(newGetValuesCmd(cfg, out))
	cmd.AddCommand(newGetManifestCmd(cfg, out))
	cmd.AddCommand(newGetHooksCmd(cfg, out))
	cmd.AddCommand(newGetNotesCmd(cfg, out))

	return cmd
}
