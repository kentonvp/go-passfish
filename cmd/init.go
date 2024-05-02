package cmd

import (
	"github.com/spf13/cobra"

	"passfish/internal/config"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the password manager",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if config file exists.
		cfg, err := config.NewConfig(cfgFile)
		if err != nil {
			// If it does, print an error message and exit.
		}
		// If it doesn't, create the config file.
		// Print a success message.

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
