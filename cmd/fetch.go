/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/bl155x0/jsloot/loot"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Downloads the JavaScript files specified in the given file",
	Run: func(cmd *cobra.Command, args []string) {
		handleVerbose(cmd)
		fileName, err := cmd.Flags().GetString("file")
		cobra.CheckErr(err)
		err = loot.FetchFromFile(fileName, true)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringP("file", "f", "", "The lootbox file to store spotted URLs")
	fetchCmd.MarkFlagRequired("file")
}
