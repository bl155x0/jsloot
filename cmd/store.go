/*
Copyright © 2024 bl155x0
*/
package cmd

import (
	"os"

	"github.com/bl155x0/jsloot/loot"
	"github.com/spf13/cobra"
)

// storeCmd reads a JSFile as JSON from stdin and stores its content to disk
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Stores a JS file read as JSON from stdin",
	Run: func(cmd *cobra.Command, args []string) {
		handleVerbose(cmd)

		targetDirectory, err := cmd.Flags().GetString("directory")
		cobra.CheckErr(err)
		if targetDirectory == "" {
			targetDirectory, err = os.Getwd()
			cobra.CheckErr(err)
		}

		//Parse the JSFile from JSON
		jsFile, err := loot.ParseJSFile(os.Stdin)
		cobra.CheckErr(err)

		//Store the Content
		fileName, err := loot.StoreJSFile(jsFile, targetDirectory)
		cobra.CheckErr(err)
		cmd.Println(fileName)
	},
}

func init() {
	storeCmd.Flags().StringP("directory", "d", "", "The output directory")
}
