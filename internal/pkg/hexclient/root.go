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

	mapCmd.PersistentFlags().Int64P("radius", "r", 0, "radius of hexagon circle")

	rootCmd.AddCommand(hexCmd)
	hexCmd.AddCommand(hexInfoCmd)

	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(repoAddCmd)
	repoAddCmd.AddCommand(repoAddFileCmd)
	repoCmd.AddCommand(repoDelCmd)
	repoCmd.AddCommand(repoGetCmd)
	repoGetCmd.AddCommand(repoGetAllCmd)

	rootCmd.AddCommand(hexStatusCmd)
	hexStatusCmd.AddCommand(hexStatusServerCmd)
	hexStatusCmd.AddCommand(hexStatusStorageCmd)
	hexStatusCmd.AddCommand(hexStatusClientCmd)

	rootCmd.PersistentFlags().BoolP("secure", "s", true, "secure connection")
	rootCmd.PersistentFlags().StringP("addr", "a", "localhost:8080", "server address")

}
