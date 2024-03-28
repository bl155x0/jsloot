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
		handleVerbose(cmd)
		lootBoxFile, err := cmd.Flags().GetString("file")
		cobra.CheckErr(err)
		if lootBoxFile == "" {
			fmt.Fprintf(os.Stderr, "Error: loot file not provided\n")
			cmd.Usage()
			os.Exit(-1)
		}
		if lootBoxFile != "" {
			cobra.CheckErr(loot.AddToFile(lootBoxFile, args))
		}
	},
}

func init() {
	addCmd.Flags().StringP("file", "f", "", "The lootbox file to store spotted URLs")
	rootCmd.AddCommand(addCmd)
}
