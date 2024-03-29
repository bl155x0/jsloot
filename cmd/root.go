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
	Use: "jsloot",
	Long: `
 ▄▄▄██▀▀▀██████  ██▓     ▒█████   ▒█████  ▄▄▄█████▓ 🐿
   ▒██ ▒██    ▒ ▓██▒    ▒██▒  ██▒▒██▒  ██▒▓  ██▒ ▓▒
   ░██ ░ ▓██▄   ▒██░    ▒██░  ██▒▒██░  ██▒▒ ▓██░ ▒░
▓██▄██▓  ▒   ██▒▒██░    ▒██   ██░▒██   ██░░ ▓██▓ ░ 
 ▓███▒ ▒██████▒▒░██████▒░ ████▓▒░░ ████▓▒░  ▒██▒ ░ 
 ▒▓▒▒░ ▒ ▒▓▒ ▒ ░░ ▒░▓  ░░ ▒░▒░▒░ ░ ▒░▒░▒░   ▒ ░░   
 ▒ ░▒░ ░ ░▒  ░ ░░ ░ ▒  ░  ░ ▒ ▒░   ░ ▒ ▒░     ░    
 ░ ░ ░ ░  ░  ░    ░ ░   ░ ░ ░ ▒  ░ ░ ░ ▒    ░      
 ░   ░       ░      ░  ░    ░ ░      ░ ░           

    	       JavaScript URL looter
                     ★ bl155 ★



`,
	Run: func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func Execute() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(getAllCmd)
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
