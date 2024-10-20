package cmd

import (
	"errors"
	"fmt"
	"log"
	"passfish/internal/database"
	"passfish/internal/models"
	"passfish/internal/stringutils"
	"strings"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a login",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.New(cfg.DbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		toRemove, err := cmd.Flags().GetString("login")
		if err != nil {
			log.Fatal(err)
		}

		if db.NumberOfCredentials() == 0 {
			log.Println("❌", errors.New("no logins to remove"))
			return
		}

		var cred *models.Credentials
		for {
			if toRemove == "" {
				toRemove, err = readStringInput("Enter the login to remove: ")
				if err != nil {
					log.Fatal(err)
				}
			}
			fmt.Printf("🗑️ Removed %s\n", toRemove)

			if cred, err = db.GetCredentials(toRemove); err != nil {
				log.Println("❌", errors.Join(err, fmt.Errorf("could not find login \"%s\"", toRemove)))
				titles, err := db.GetTitles()
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("Did you mean [%s]\n", strings.Join(stringutils.FindTopMatches(titles, toRemove, 3), ", "))

				// Reset to trigger user input again.
				toRemove = ""
			} else {
				db.DeleteCredentials(cred.Base.Title)
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringP("login", "l", "", "pass in login name to remove")
}
