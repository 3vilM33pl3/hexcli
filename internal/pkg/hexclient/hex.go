/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package hexclient

import (
	"fmt"
	"github.com/3vilm33pl3/hexcli/internal/pkg/hexcli"
	"github.com/spf13/cobra"
	"strconv"
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
	Use:   "place [x,y,z] [direction] [content]",
	Short: "add hexagon with coordinate [x,y,z] and compressed content file [content]",
	Long: "Example: nb hex place --secure=false -- 0 -5 5 N 0000-0000-0000-0000" +
		"The double dash is needed to indicate no more flags are coming and everything is interpreted as an normal argument. " +
		"This is needed so that negative numbers don't get interpreted as flags",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)

		x, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Printf("x value is not a number %s : %e", args[0], err)
			return
		}

		y, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Printf("y value is not a number %s : %e", args[1], err)
			return
		}

		z, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			fmt.Printf("z value is not a number %s : %e", args[2], err)
			return
		}

		var hexDirection hexcli.Direction
		switch args[3] {
		case "N":
			hexDirection = hexcli.Direction_N
		case "NE":
			hexDirection = hexcli.Direction_NE
		case "E":
			hexDirection = hexcli.Direction_E
		case "SE":
			hexDirection = hexcli.Direction_SE
		case "S":
			hexDirection = hexcli.Direction_S
		case "SW":
			hexDirection = hexcli.Direction_SW
		case "W":
			hexDirection = hexcli.Direction_W
		case "NW":
			hexDirection = hexcli.Direction_NW
		default:
			fmt.Printf("direction value is not a valid direction %s ", args[3])
			return
		}

		hex := &hexcli.Hex{
			X:         x,
			Y:         y,
			Z:         z,
			Direction: hexDirection,
			Reference: "",
		}

		err = client.HexagonPlace(hex)

		if err != nil {
			fmt.Printf("Error placing hexagon on map: %s", err)
			return
		}

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
		var refList hexcli.HexRefList

		for _, ref := range args {
			refList.Ref = append(refList.Ref, &hexcli.HexReference{Ref: ref})
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
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var refList hexcli.HexRefList

		for _, ref := range args {
			refList.Ref = append(refList.Ref, &hexcli.HexReference{Ref: ref})
		}

		err = client.RepoDeleteHexagon(&refList)

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}
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
