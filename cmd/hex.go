/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/3vilM33pl3/hexclient/internal/pkg/hexcloud"

	"github.com/spf13/cobra"
)

// hexCmd represents the hex command
var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Manipulate hexagons",
	Long:  `A longer description that spans multiple lines and likely contains examples.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hex called")
	},
}

var hexPlaceCmd = &cobra.Command{
	Use:   "place [x,y,z] [content]",
	Short: "add hexagon with coordinate [x,y,z] and compressed content file [content]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex place command: %v\n", args)
	},
}

var hexGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get hexagon at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex get command: %v\n", args)
	},
}

var hexRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove hexagon at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex remove command: %v\n", args)
	},
}

var hexInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "show information of hexagon at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex info command: %v\n", args)
	},
}

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manipulate hexagons",
	Long:  `A longer description that spans multiple lines and likely contains examples.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("repo called")
	},
}

var repoAddCmd = &cobra.Command{
	Use:   "add [ref]",
	Short: "add hexagon to repository with reference [ref]",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var refList hexcloud.HexRefList

		for _, ref := range args {
			refList.Ref = append(refList.Ref, &hexcloud.HexReference{Ref: ref})
		}

		err = client.RepoAddHexagon(&refList)

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

	},
}

var repoDelCmd = &cobra.Command{
	Use:   "del [ref]",
	Short: "delete hexagon from repository with reference [ref] ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside repo delete command: %v\n", args)
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
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		status, err := client.StatusServer()
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		} else {
			fmt.Printf("Status server: %s\n", status)
		}

	},
}

var hexStatusStorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "status of hexagon storage",
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

		status, err := client.StatusStorage()
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		} else {
			fmt.Printf("Status storage: %s\n", status)
		}
	},
}

var hexStatusClientCmd = &cobra.Command{
	Use:   "client",
	Short: "status of hexagon network clients",
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

		status, err := client.StatusClients()
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		} else {
			fmt.Printf("Status clients: %s\n", status)
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

func init() {

	rootCmd.AddCommand(hexCmd)
	hexCmd.AddCommand(hexPlaceCmd)
	hexCmd.AddCommand(hexGetCmd)
	hexCmd.AddCommand(hexRemoveCmd)
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
