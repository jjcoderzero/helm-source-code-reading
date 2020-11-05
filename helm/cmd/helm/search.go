package main

import (
	"io"

	"github.com/spf13/cobra"
)

const searchDesc = `
Search provides the ability to search for Helm charts in the various places
they can be stored including the Helm Hub and repositories you have added. Use
search subcommands to search different locations for charts.
`

func newSearchCmd(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:               "search [keyword]",
		Short:             "search for a keyword in charts",
		Long:              searchDesc,
		ValidArgsFunction: noCompletions, // Disable file completion
	}

	cmd.AddCommand(newSearchHubCmd(out))
	cmd.AddCommand(newSearchRepoCmd(out))

	return cmd
}
