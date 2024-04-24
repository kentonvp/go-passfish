package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a login",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Search for a logins...")
		if len(args) == 0 {
			log.Fatal("No search terms provided.")
		}
		for _, arg := range args {
			fmt.Println("Searching for: ", arg)
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
