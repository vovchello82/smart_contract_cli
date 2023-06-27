package tokensale

import (
	"tokensale/pkg/tokensale"

	"github.com/spf13/cobra"
)

var getCurrentPriceCommand = &cobra.Command{
	Use:   "price",
	Short: "returns current price in wei",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tokensale.GetCurrentPriceInWei()
	},
}

func init() {
	rootCmd.AddCommand(getCurrentPriceCommand)
}
