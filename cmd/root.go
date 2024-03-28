/*
Copyright © 2024 bl155x0
*/
package cmd

import (
	"os"

	"github.com/bl155x0/jsloot/loot"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jsloot",
	Short: "jsloot - file looter",
	Long: `jsloot is a java script looter.

jsloot can do the following : 
 * Collect given JavaScript URLs and store them in a lootbox file.  
 * Process this lootbox file and download all the JavaScript files
 * Beautify the JavaScript files
`,
	Run: func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose")
}

func handleVerbose(cmd *cobra.Command) {
	v, e := cmd.Flags().GetBool("verbose")
	cobra.CheckErr(e)
	loot.SetVerbose(v)
}
