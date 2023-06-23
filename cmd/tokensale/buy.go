package tokensale

import (
	"tokensale/pkg/tokensale"

	"github.com/spf13/cobra"
)

var buyCommand = &cobra.Command{
	Use:   "buy",
	Short: "pay for a token if avaiable",
	Run: func(cmd *cobra.Command, args []string) {
		tokensale.BuyToken()
	},
}

func init() {
	rootCmd.AddCommand(buyCommand)
}
