package cmd

import (
	"errors"
	"fmt"
	"log"
	"passfish/internal/clipboard"
	"passfish/internal/database"
	"passfish/internal/models"
	"passfish/internal/stringutils"
	"strings"
  "os"

	"github.com/spf13/cobra"
)

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Get a login",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.New(cfg.DbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Request login
		var cred *models.Credentials
		var login string = ""
		for {
			if login == "" {
				login, err = readStringInput("Enter the login to copy: ")
				if err != nil {
					log.Fatal(err)
				}
			}

			if cred, err = db.GetCredentials(login); err != nil {
				log.Println("❌", errors.Join(err, fmt.Errorf("could not find login \"%s\"", login)))
				titles, err := db.GetTitles()
				if err != nil {
					log.Fatal(err)
				}
				matches := stringutils.FindTopMatches(titles, login, 3)
				log.Printf("Did you mean %s, or %s\n",
					strings.Join(matches[0:len(matches)-1], ", "),
					matches[len(matches)-1])

				// Reset to trigger user prompt again
				login = ""
        continue
			}

      passphrase, found := os.LookupEnv("PASSFISH_PASSPHRASE")
      if !found {
        passphrase, err = readPasswordInput("Enter the passphrase: ")
        if err != nil {
          log.Fatal(err)
        }
      }

      if !db.VerifyPassphrase(passphrase) {
        log.Fatal("❌ Incorrect passphrase!")
      }

      clipboard.Copy(cred.DecryptPassword(passphrase))
      fmt.Printf("Username %s\n", cred.Base.Username)
      fmt.Printf("Copied 🔑\n")
      db.MarkAccessed(login)
      break
		}
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
}
