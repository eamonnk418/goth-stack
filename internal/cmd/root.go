package cmd

import "github.com/spf13/cobra"

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goth-stack",
		Short: "goth-stack",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}

func Execute() {
	cobra.CheckErr(newRootCmd().Execute())
}
