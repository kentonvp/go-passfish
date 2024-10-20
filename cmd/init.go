package cmd

import (
	"fmt"
	"log"
	"passfish/internal/database"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the password manager",
	Run: func(cmd *cobra.Command, args []string) {
    db, err := database.New(cfg.DbPath)
    if err != nil {
      log.Fatal(err)
    }
    defer db.Close()

    // Create tables if they don't exist
    db.CreateTables()

    passphrase := db.GetPassphrase()
    for passphrase == "" {
      resp, err := readPasswordInput("Insert a passphrase, this should be *extra* secret: ")
      if err != nil {
        log.Fatal(err)
      }
      if len(resp) < 8 {
        fmt.Println("Come on... Your passphrase must be at LEAST 8 characters long.")
        continue
      }

      passphrase = resp
      if err = db.SetPassphrase(passphrase); err != nil {
        log.Fatal(err)
      }
    }
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
