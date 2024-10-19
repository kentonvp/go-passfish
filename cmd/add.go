package cmd

import (
	"fmt"
  "os"
	"log"
	"passfish/internal/clipboard"
	"passfish/internal/config"
	"passfish/internal/database"
	"passfish/internal/models"
	"passfish/internal/passwords"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a login",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.New(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		db, err := database.New(cfg.DbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

    // Lookup passphrase first
    passphrase, found := os.LookupEnv("PASSFISH_PASSPHRASE")
    if !found {
      passphrase, err = readStringInput("Enter the passphrase: ")
      if err != nil {
        log.Fatal(err)
      }
    }

    // TODO: verify_passphrase

		db.CreateCredentialsTable()

		login, err := cmd.Flags().GetString("login")
		if err != nil {
			log.Fatal(err)
		}
		for login == "" {
			login, err = readStringInput("Enter Login: ")
			if login == "" {
				log.Print("❌ Login cannot be empty.")
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
		for username == "" {
			username, err = readStringInput(fmt.Sprintf("Enter a username for %s: ", login))
			if username == "" {
				log.Print("❌ Username cannot be empty.")
			}
			if err != nil {
				log.Fatal(err)
			}
		}

		genPassword, _ := cmd.Flags().GetBool("create")
		passwordLength, _ := cmd.Flags().GetInt("password-length")
		password := passwords.New(passwordLength)
		if !genPassword {
			password, err = readPasswordInput("Enter Password: ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println()
			if password == "" {
				log.Print("Entered empty password, using a generated password instead.")
				password = passwords.New(passwordLength)
			} else {
				confirmation, err := readPasswordInput("Confirm Password: ")
				fmt.Println()
				if err != nil {
					log.Fatal(err)
				}

				if password != confirmation {
					log.Fatal("❌ Passwords do not match!!!")
				}
			}
		}

		if err := clipboard.Copy(password); err != nil {
			log.Println("❌ Error copying password to clipboard.")
		} else {
			fmt.Println("Password is copied to 📋...")
		}
		creds := models.BaseCredentials{
			Title:    login,
			Username: username,
			Password: passwords.Encrypt(password, passphrase),
		}

		if err := db.InsertCredentials(creds); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("login", "l", "", "pass in new login name")
	addCmd.Flags().StringP("username", "u", "", "pass in username")
	addCmd.Flags().BoolP("create", "c", false, "create a secure password automatically")
	addCmd.Flags().IntP("password-length", "N", 16, "length of the password to generate")
}
