package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display added logins",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Displaying all logins...")

		// Check if the time_sorted flag is set.
		time_sorted, err := cmd.Flags().GetBool("time-sorted")
		if err != nil {
			log.Fatal(err)
		}
		if time_sorted {
			// Display logins sorted by most recently accessed.
			fmt.Println("Displaying logins sorted by most recently accessed.")
		} else {
			// Display logins sorted by name.
			fmt.Println("Displaying logins sorted by name.")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("time-sorted", "t", false, "sort logins by most recently accessed")
}
