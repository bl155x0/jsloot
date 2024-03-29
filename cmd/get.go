/*
Copyright © 2024 bl155x0
*/
package cmd

import (
	"os"

	"github.com/bl155x0/jsloot/loot"
	"github.com/spf13/cobra"
)

// getCmd represents the fetch command
var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"fetch"},
	Short:   "Downloads a single JavaScript file",
	Run: func(cmd *cobra.Command, args []string) {
		handleVerbose(cmd)
		url, err := cmd.Flags().GetString("url")
		cobra.CheckErr(err)
		targetDirectory, err := cmd.Flags().GetString("directory")
		cobra.CheckErr(err)
		if targetDirectory == "" {
			var err error
			targetDirectory, err = os.Getwd()
			cobra.CheckErr(err)
		}
		loot.Fetch(url, targetDirectory, true)
	},
}

func init() {
	getCmd.Flags().StringP("url", "u", loot.DefaultLootBoxFile, "The URL to download the JavaScript file from")
	getCmd.Flags().StringP("directory", "d", "", "The output directory")
	getCmd.MarkFlagRequired("url")
}
