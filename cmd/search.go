package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a login",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("🔎 No search terms provided.")
			os.Exit(0)
		}
		for _, arg := range args {
			fmt.Println("Searching for: ", arg)
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
