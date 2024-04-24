/*
Copyright © 2024 Kenton Van Peursem <kentonvp@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "passfish",
	Short:   "CLI Password Manager",
	Version: "0.1.0",
}

var cfgFile string

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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", fmt.Sprintf("%s/.passfish.yaml", dirname), "config file")
}
