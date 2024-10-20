package cmd

import (
	"fmt"
	"log"
	"os"
	"passfish/internal/database"
	"passfish/internal/stringutils"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a login",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("🔎 No search terms provided.")
			os.Exit(0)
		}

		db, err := database.New(cfg.DbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// if more than 1 arguments then join them
		// and search for the joined string
		var search string
		if len(args) > 1 {
			search = strings.Join(args, " ")
		} else {
			search = args[0]
		}

		titles, err := db.GetTitles()
		if err != nil {
			log.Fatal(err)
		}

		start := time.Now()
		top := stringutils.FindTopMatches(titles, search, 5)
		fmt.Printf("🔎 Top Matches (%s)\n", time.Since(start))
		for i, w := range top {
			fmt.Printf("%d | %s\n", i+1, w)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
