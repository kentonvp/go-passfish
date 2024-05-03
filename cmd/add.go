package cmd

import (
	"fmt"
	"log"
	"passfish/internal/clipboard"
	"passfish/internal/models"
	"passfish/internal/passwords"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a login",
	Run: func(cmd *cobra.Command, args []string) {
		login, err := cmd.Flags().GetString("login")
		if err != nil {
			log.Fatal(err)
		}
		if login == "" {
			login, err = readStringInput("Enter Login: ")
			if login == "" {
				log.Fatal("❌ Login cannot be empty.")
			}
			if err != nil {
				log.Fatal(err)
			}
		}

		// TODO: Check if login already exists. If it does, ask user if they want to overwrite.

		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatal(err)
		}
		if username == "" {
			username, err = readStringInput(fmt.Sprintf("Enter a username for %s: ", login))
			if username == "" {
				log.Fatal("❌ Username cannot be empty.")
			}
			if err != nil {
				log.Fatal(err)
			}
		}

		genPassword, err := cmd.Flags().GetBool("create")
		passwordLength, err := cmd.Flags().GetInt("password-length")
		password := passwords.New(passwordLength)
		if !genPassword {
			password, err = readPasswordInput("Enter Password: ")
			fmt.Println()
			if password == "" {
				log.Fatal("❌ Password cannot be empty.")
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
				log.Fatal("❌ Passwords do not match!!!")
			}
		}

		if err := clipboard.CopyToClipboard(password); err != nil {
			log.Println("❌ Error copying password to clipboard.")
		} else {
			fmt.Println("Password is copied to 📋...")
		}
		creds := models.Credentials{
			Title: login,
			Username: username,
			Password: passwords.Encrypt(password, config.Passphrase())
		}
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("login", "l", "", "pass in new login name")
	addCmd.Flags().StringP("username", "u", "", "pass in username")
	addCmd.Flags().BoolP("create", "c", false, "create a secure password automatically")
	addCmd.Flags().IntP("password-length", "N", 16, "length of the password to generate")
}
