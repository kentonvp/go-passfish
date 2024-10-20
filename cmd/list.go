package cmd

import (
	"fmt"
	"log"
	"passfish/internal/database"
	"passfish/internal/models"
	"passfish/internal/stringutils"
	"sort"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display added logins",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.New(cfg.DbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		titles, err := db.GetTitles()
		if err != nil {
			log.Fatal(err)
		}

		creds := make([]models.Credentials, 0, len(titles))
		for _, title := range titles {
			cred, err := db.GetCredentials(title)
			if err != nil {
				log.Fatal(err)
			}
			creds = append(creds, *cred)
		}

    // Display logins sorted by most recently accessed.
    fmt.Println("🕰️ Displaying logins sorted by most recently accessed.")

    // Sort creds by their last accessed time.
    sort.Slice(creds, func(i int, j int) bool {
      return creds[i].LastAccessed.Compare(creds[j].LastAccessed) > 0
    })

    // Display the logins.
		for i, cred := range creds {
			fmt.Printf("%3d | %s | %s | %s\n",
				i,
				stringutils.CenterString(cred.Base.Title, 20),
				stringutils.CenterString(cred.Base.Username, 20),
				cred.LastAccessed)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
