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
	Use:   "add [URL1..URLn]",
	Short: "Appends URLs to a file",
	Run: func(cmd *cobra.Command, args []string) {
		lootBoxFile, err := cmd.Flags().GetString("file")
		cobra.CheckErr(err)
		if lootBoxFile == "" {
			fmt.Fprintf(os.Stderr, "Error: loot file not provided\n")
			cmd.Usage()
			os.Exit(-1)
		}
		handleVerbose(cmd)
		if lootBoxFile != "" {
			cobra.CheckErr(loot.AddToFile(lootBoxFile, args))
		}
	},
}

func init() {
	addCmd.Flags().StringP("file", "f", "", "The file to append the URLs to")
}
