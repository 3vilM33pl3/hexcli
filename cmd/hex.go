/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hexCmd represents the hex command
var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Manipulate hexagons",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hex called")
	},
}

var hexConnectCmd = &cobra.Command{
	Use:   "connect [server]",
	Short: "connect to hexcloud server with address [server]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside connect command: %v\n", args)
		var serverAddr string
		if len(args) == 0 {
			fmt.Printf("Using default server localhost:8080\n")
			serverAddr = "localhost:8080"
		} else {
			serverAddr = args[0]
		}

		c, err := NewClient(serverAddr, true)
		if err != nil {
			fmt.Printf("Unable to connect %s", err)
		} else {
			fmt.Printf("Connected %x", c)
		}

	},
}

var hexAddCmd = &cobra.Command{
	Use:   "add [x,y,z] [content]",
	Short: "add hexagon with coordinate [x,y,z] and compressed content file [content]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex add command: %v\n", args)
	},
}

var hexDelCmd = &cobra.Command{
	Use:   "del [coords]",
	Short: "delete hexagon at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex del command: %v\n", args)
	},
}

var hexUpdateCmd = &cobra.Command{
	Use:   "update [x,y,z] [content]",
	Short: "update hexagon at coordinate [x,y,z] with compressed content file [content] ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex update command: %v\n", args)
	},
}

var hexInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "show information of hexagon at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex info command: %v\n", args)
	},
}

var hexGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get hexagon at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex info command: %v\n", args)
	},
}

var hexStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "status of hexagon network",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex status command: %v\n", args)

	},
}

var hexStatusServerCmd = &cobra.Command{
	Use:   "server",
	Short: "status of hexagon server",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr := "localhost:8080"
		if len(args) > 0 {
			serverAddr = args[0]
		}

		secure, _ := cmd.Flags().GetBool("secure")
		client, err := NewClient(serverAddr, secure)

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		status, err := client.Status()
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		} else {
			fmt.Printf("Status server: %s\n", status)
		}

	},
}

var hexPackageCmd = &cobra.Command{
	Use:   "pack [dir]",
	Short: "packagae content in directory [dir]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex info command: %v\n", args)
	},
}

var hexStatusStorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "status of hexagon storage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside status storage command: %v\n", args)
	},
}

var hexStatusClientCmd = &cobra.Command{
	Use:   "client",
	Short: "status of hexagon network clients",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside status client command: %v\n", args)
	},
}

func init() {
	rootCmd.AddCommand(hexConnectCmd)

	rootCmd.AddCommand(hexCmd)
	hexCmd.AddCommand(hexAddCmd)
	hexCmd.AddCommand(hexDelCmd)
	hexCmd.AddCommand(hexUpdateCmd)
	hexCmd.AddCommand(hexInfoCmd)
	hexCmd.AddCommand(hexGetCmd)

	rootCmd.AddCommand(hexStatusCmd)
	hexStatusCmd.AddCommand(hexStatusServerCmd)
	hexStatusCmd.AddCommand(hexStatusStorageCmd)
	hexStatusCmd.AddCommand(hexStatusClientCmd)

	rootCmd.PersistentFlags().BoolP("secure", "s", true, "secure connection")

}
