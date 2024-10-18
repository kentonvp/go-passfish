package cmd

import (
	"fmt"
	"log"
	"passfish/internal/config"
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
		cfg, err := config.New(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		db, err := database.New(cfg.DbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Check if the time_sorted flag is set.
		time_sorted, err := cmd.Flags().GetBool("time-sorted")
		if err != nil {
			log.Fatal(err)
		}

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
		if time_sorted {
			// Display logins sorted by most recently accessed.
			fmt.Println("🕰️ Displaying logins sorted by most recently accessed.")

			// Sort creds by their last accessed time.
			sort.Slice(creds, func(i int, j int) bool {
				return creds[i].LastAccessed.Compare(creds[j].LastAccessed) > 0
			})
		} else {
			// Display logins sorted by name.
			fmt.Println("🔤 Displaying logins sorted by name.")

			// Sort creds by their title.
			sort.Slice(creds, func(i int, j int) bool {
				return creds[i].Base.Title < creds[j].Base.Title
			})
		}
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

	listCmd.Flags().BoolP("time-sorted", "t", false, "sort logins by most recently accessed")
}
