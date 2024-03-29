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

// getAllCmd represents the fetch command
var getAllCmd = &cobra.Command{
	Use:     "getall",
	Aliases: []string{"get-all", "all", "loot"},
	Short:   "Downloads all the JavaScript files specified in the given file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Fprintf(os.Stderr, "Error: unknown arguments given\n")
			cmd.Usage()
			os.Exit(-1)
		}
		handleVerbose(cmd)
		//loot from file
		fileName, err := cmd.Flags().GetString("file")
		cobra.CheckErr(err)
		err = loot.FetchFromFile(fileName, true)
		cobra.CheckErr(err)
	},
}

func init() {
	getAllCmd.Flags().StringP("file", "f", loot.DefaultLootBoxFile, "The file containing the URLs to download")
}
