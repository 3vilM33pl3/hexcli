/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package hexclient

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "nb",
	Short: "Command line interface for game content manipulation",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(mapCmd)
	mapCmd.AddCommand(mapAddCmd)
	mapCmd.AddCommand(mapGetCmd)
	mapCmd.AddCommand(mapRemoveCmd)

	rootCmd.AddCommand(hexCmd)
	hexCmd.AddCommand(hexInfoCmd)

	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(repoAddCmd)
	repoCmd.AddCommand(repoDelCmd)

	rootCmd.AddCommand(hexStatusCmd)
	hexStatusCmd.AddCommand(hexStatusServerCmd)
	hexStatusCmd.AddCommand(hexStatusStorageCmd)
	hexStatusCmd.AddCommand(hexStatusClientCmd)

	rootCmd.PersistentFlags().BoolP("secure", "s", true, "secure connection")
	rootCmd.PersistentFlags().StringP("addr", "a", "localhost:8080", "server address")

}
