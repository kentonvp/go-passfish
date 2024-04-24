package cmd

import (
	"fmt"
	"log"
	"passfish/clipboard"
	"passfish/passwords"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a login",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add a login...")

		login, err := cmd.Flags().GetString("login")
		if err != nil {
			log.Fatal(err)
		}
		if login == "" {
			login, err = readStringInput("Enter Login: ")
			if login == "" {
				log.Fatal("Login cannot be empty.")
			}
			if err != nil {
				log.Fatal(err)
			}
		}

		// Check if login already exists. If it does, ask user if they want to overwrite.

		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatal(err)
		}
		if username == "" {
			username, err = readStringInput(fmt.Sprintf("Enter a username for %s: ", login))
			if username == "" {
				log.Fatal("Username cannot be empty.")
			}
			if err != nil {
				log.Fatal(err)
			}
		}

		genPassword, err := cmd.Flags().GetBool("generate-password")
		password := passwords.GeneratePassword()
		if !genPassword {
			password, err = readPasswordInput("Enter Password: ")
			fmt.Println()
			if password == "" {
				log.Fatal("Password cannot be empty.")
			}
			if err != nil {
				log.Fatal(err)
			}
			confirmation, err := readPasswordInput("Confirm Password: ")
			fmt.Println()
			if err != nil {
				log.Fatal(err)
			}

			if password != confirmation {
				log.Fatal("Passwords do not match!!!")
			}
		}

		copiedToClipboard, err := clipboard.CopyToClipboard(password)
		if err != nil {
			log.Fatal(err)
		}

		if copiedToClipboard {
			fmt.Println("Password is copied to clipboard...")
		}
		fmt.Printf("Adding {%s, %s, %s}", login, username, password)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("login", "l", "", "pass in new login name")
	addCmd.Flags().StringP("username", "u", "", "pass in username")
	addCmd.Flags().BoolP("generate-password", "g", false, "generate a password automatically")
}
