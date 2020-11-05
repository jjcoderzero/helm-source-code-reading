package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/cmd/helm/require"
	"helm.sh/helm/v3/pkg/action"
)

const verifyDesc = `
Verify that the given chart has a valid provenance file.

Provenance files provide cryptographic verification that a chart has not been
tampered with, and was packaged by a trusted provider.

This command can be used to verify a local chart. Several other commands provide
'--verify' flags that run the same validation. To generate a signed package, use
the 'helm package --sign' command.
`

func newVerifyCmd(out io.Writer) *cobra.Command {
	client := action.NewVerify()

	cmd := &cobra.Command{
		Use:   "verify PATH",
		Short: "verify that a chart at the given path has been signed and is valid",
		Long:  verifyDesc,
		Args:  require.ExactArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				// Allow file completion when completing the argument for the path
				return nil, cobra.ShellCompDirectiveDefault
			}
			// No more completions, so disable file completion
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := client.Run(args[0])
			if err != nil {
				return err
			}

			fmt.Fprint(out, client.Out)

			return nil
		},
	}

	cmd.Flags().StringVar(&client.Keyring, "keyring", defaultKeyring(), "keyring containing public keys")

	return cmd
}
