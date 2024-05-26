package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blonet",
	Short: "Blonet is a platform for smart contract development and deployment",
	Long:  `Blonet is a platform for developing, deploying, and interacting with smart contracts on various blockchain networks.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Root command logic
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add sub-commands here
}
