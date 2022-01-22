/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package hexclient

import (
	"fmt"
	"github.com/3vilM33pl3/hexclient/internal/pkg/hexclient"
	"github.com/3vilM33pl3/hexclient/internal/pkg/hexcloud"
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
		args = append(args, "0")
		hexGetRingCmd.Run(cmd, args)
		fmt.Printf("Inside get info command: %v\n", args)
	},
}

var hexGetRingCmd = &cobra.Command{
	Use:   "ring",
	Short: "get hexagon ring at `coord:` [x,y,z] radius: [r], fill [bool]",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("server")
		secure, _ := cmd.Flags().GetBool("secure")
		radius, _ := cmd.Flags().GetInt64("radius")
		fill, _ := cmd.Flags().GetBool("fill")

		client, err := hexclient.NewClient(serverAddr, secure)
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		hexes, err := client.GetHexagonRing(&hexcloud.Hex{
					X:         0,
					Y:         0,
					Z:         0,
					Type:      "",
					Direction: "",
				},
		radius,fill )

		if err != nil {
			fmt.Printf("Error getting hexagons #{err}")
			return
		}

		for _, hex := range hexes.GetHc() {
			fmt.Printf("[%d, %d, %d] %s %s\n", hex.X, hex.Y, hex.Z, hex.Direction, hex.Type)
		}

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
		serverAddr, _ := cmd.Flags().GetString("server")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := hexclient.NewClient(serverAddr, secure)
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		status, err := client.StatusServer()
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		} else {
			fmt.Printf("StatusServer server: %s\n", status)
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

	rootCmd.AddCommand(hexCmd)

	hexCmd.AddCommand(hexAddCmd)
	hexCmd.AddCommand(hexDelCmd)
	hexCmd.AddCommand(hexUpdateCmd)
	hexCmd.AddCommand(hexInfoCmd)

	hexCmd.AddCommand(hexGetCmd)
	hexGetCmd.AddCommand(hexGetRingCmd)
	hexGetRingCmd.PersistentFlags().Int64P("radius", "r", 0, "radius of ring")
	hexGetRingCmd.PersistentFlags().BoolP("fill", "f", true, "fill ring")

	rootCmd.AddCommand(hexStatusCmd)
	hexStatusCmd.AddCommand(hexStatusServerCmd)
	hexStatusCmd.AddCommand(hexStatusStorageCmd)
	hexStatusCmd.AddCommand(hexStatusClientCmd)

	rootCmd.PersistentFlags().BoolP("secure", "s", true, "secure connection")
	rootCmd.PersistentFlags().StringP("server", "a","localhost:8080", "server address")

}
