package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/cmd/helm/require"
)

var envHelp = `
Env prints out all the environment information in use by Helm.
`

func newEnvCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "env",
		Short: "helm client environment information",
		Long:  envHelp,
		Args:  require.MaximumNArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				keys := getSortedEnvVarKeys()
				return keys, cobra.ShellCompDirectiveNoFileComp
			}

			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		Run: func(cmd *cobra.Command, args []string) {
			envVars := settings.EnvVars()

			if len(args) == 0 {
				// Sort the variables by alphabetical order.
				// This allows for a constant output across calls to 'helm env'.
				keys := getSortedEnvVarKeys()

				for _, k := range keys {
					fmt.Fprintf(out, "%s=\"%s\"\n", k, envVars[k])
				}
			} else {
				fmt.Fprintf(out, "%s\n", envVars[args[0]])
			}
		},
	}
	return cmd
}

func getSortedEnvVarKeys() []string {
	envVars := settings.EnvVars()

	var keys []string
	for k := range envVars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}
