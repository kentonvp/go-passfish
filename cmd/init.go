package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the password manager",
	Run: func(cmd *cobra.Command, args []string) {
		if err == nil {
			fmt.Println("WARNING!!! Configuration file already exists!")
			var resp string = ""
			for resp == "" {
				resp, err = readStringInput("Do you want to overwrite the configuration file? (y/n): ")
				if err != nil {
					log.Fatal(err)
				}
				if resp != "y" && resp != "n" {
					fmt.Printf("Please try again, \"%v\" is not a valid response.\n", resp)
					resp = ""
				}
			}

			if resp == "n" {
				fmt.Println("Exiting...")
				os.Exit(0)
			}
		}

		// Overwrite the configuration file
		if err := config.CreateConfigFile(cfgFile); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
