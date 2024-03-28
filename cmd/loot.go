/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/bl155x0/jsloot/loot"
	"github.com/spf13/cobra"
)

// lootCmd represents the fetch command
var lootCmd = &cobra.Command{
	Use:     "loot",
	Aliases: []string{"fetch"},
	Short:   "Downloads the JavaScript files specified in the given file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Fprintf(os.Stderr, "Error: unknown arguments given\n")
			cmd.Usage()
			os.Exit(-1)
		}
		handleVerbose(cmd)
		fileName, err := cmd.Flags().GetString("file")
		cobra.CheckErr(err)
		err = loot.FetchFromFile(fileName, true)
		cobra.CheckErr(err)
	},
}

func init() {
	lootCmd.Flags().StringP("file", "f", loot.DefaultLootBoxFile, "The lootbox file to store spotted URLs")
}
