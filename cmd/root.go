/*
Copyright © 2024 Kenton Van Peursem <kentonvp@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
  "passfish/internal/config"

	"github.com/spf13/cobra"
)

const AppName = "passfish"
const Version = "0.1.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     AppName,
	Short:   "CLI Password Manager",
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		if cfg, err := cmd.Flags().GetString("database"); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Using database file: ", cfg)
		}
	},
}

var cfg config.Config = config.Config{}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Default config file is $HOME/.passfish.yaml
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dirname = path.Join(dirname, ".config", AppName)
	err = os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

  var cfgFile string
	rootCmd.PersistentFlags().StringVar(
    &cfgFile,                                  // pointer to variable
    "config",                                 // flag name
    fmt.Sprintf("%s/passfish.yaml", dirname), // default value
    "config file",                            // help message
  )

  tmp, err := config.New(cfgFile)
  if err != nil {
    log.Fatal(err)
  }

  // Set to global config
  cfg.DbPath = tmp.DbPath
}
