// Package command sets the entrypoint commands for running this cli tool
package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "vaultcrawler",
	Short: "vaultcrawler is a cli tool for simplifying interactions with hashicorp vault endpoints",
	Long:  "vaultcrawler is a cli wrapper for the vault cli tool made for interacting with Hashicorp Vault endpoints primarily through the use of the token issued by vault. This token is currently obtained via the GUI, with plans to expand to unified token acquisition",
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("Successful test of vaultcrawler")
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
