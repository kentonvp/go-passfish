package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a login",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Remove a login...")
		toRemove, err := cmd.Flags().GetString("login")
		if err != nil {
			log.Fatal(err)
		}
		if toRemove == "" {
			fmt.Print("Enter Login to remove: ")
			reader := bufio.NewReader(os.Stdin)
			toRemove, err = reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			toRemove = toRemove[:len(toRemove)-1]
		}
		fmt.Println("Removing { ", toRemove, " }")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringP("login", "l", "", "pass in login name to remove")
}
