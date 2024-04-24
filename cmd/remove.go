package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a login",
	Run: func(cmd *cobra.Command, args []string) {
		toRemove, err := cmd.Flags().GetString("login")
		if err != nil {
			log.Fatal(err)
		}
		if toRemove == "" {
			toRemove, err = readStringInput("Enter login to remove: ")
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("Removing { ", toRemove, " }")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringP("login", "l", "", "pass in login name to remove")
}
