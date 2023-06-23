package tokensale

import (
	"tokensale/pkg/tokensale"

	"github.com/spf13/cobra"
)

var pushCommand = &cobra.Command{
	Use:   "push",
	Short: "push a new token to the contract",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tokensale.PushToken(args[0])
	},
}

func init() {
	rootCmd.AddCommand(pushCommand)
}
