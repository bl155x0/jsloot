/*
Copyright © 2024 bl155x0
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/bl155x0/jsloot/loot"
	"github.com/spf13/cobra"
)

// addCmd represents the spot command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds URLs to the lootbox",
	Run: func(cmd *cobra.Command, args []string) {
		lootBoxFile, err := cmd.Flags().GetString("file")
		cobra.CheckErr(err)
		lootBoxDir, err := cmd.Flags().GetString("directory")
		cobra.CheckErr(err)
		if lootBoxFile == "" && lootBoxDir == "" {
			fmt.Fprintf(os.Stderr, "Error: lootbox file or path missing\n")
			cmd.Usage()
			os.Exit(-1)
		}
		if lootBoxFile != "" && lootBoxDir != "" {
			fmt.Fprintf(os.Stderr, "Error: lootbox file and path specified while only one is allowed\n")
			cmd.Usage()
			os.Exit(-1)
		}

		if lootBoxFile != "" {
			cobra.CheckErr(loot.AddToFile(lootBoxFile, args))
		} else if lootBoxDir != "" {
			cobra.CheckErr(loot.AddToDir(lootBoxDir, args))
		}
	},
}

func init() {
	addCmd.Flags().StringP("file", "f", "", "The lootbox file to store spotted URLs")
	addCmd.Flags().StringP("directory", "d", "", "The lootbox path to store spotted URLs")
	rootCmd.AddCommand(addCmd)
}
