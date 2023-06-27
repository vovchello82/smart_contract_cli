package tokensale

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "tokensale",
	Aliases: []string{"ts"},
	Short:   "tokensale - cli dapp for communication with the tokensale smart contract",
	Long:    "tokensale - dapp for communication with the tokensale smart contract",
	Run: func(cmd *cobra.Command, arg []string) {
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("ooops ... unexpected error on cmd execution")
	}
}
