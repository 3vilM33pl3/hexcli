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

}
