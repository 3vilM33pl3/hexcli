package hexclient

import (
	"fmt"
	"github.com/3vilm33pl3/hexcli/internal/pkg/hexcli"
	"github.com/spf13/cobra"
	"strconv"
)

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "Manipulate map",
	Long:  `Functionality for adding, updating, removing everything map related`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hex called")
	},
}

var mapAddCmd = &cobra.Command{
	Use:   "add [x,y,z] [content ref]",
	Short: "add hexagon with coordinate [x,y,z] and content reference [content ref]",
	Long: "Example: nb hex place --secure=false -- 0 -5 5 N 0000-0000-0000-0000" +
		"The double dash is needed to indicate no more flags are coming and everything is interpreted as an normal argument. " +
		"This is needed so that negative numbers don't get interpreted as flags",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)

		if len(args) < 4 {
			fmt.Printf("Not enough arguments")
			return
		}
		x, y, z, id, err := extractHexCoord(args)
		if err != nil {
			return
		}

		hex := &hexcli.HexLocation{
			X:     x,
			Y:     y,
			Z:     z,
			HexID: id,
		}

		hexLocList := &hexcli.HexLocationList{}
		hexLocList.HexLoc = append(hexLocList.HexLoc, hex)

		err = client.MapAdd(hexLocList)

		if err != nil {
			fmt.Printf("Error placing hexagon on map: %s", err)
			return
		}

	},
}

var mapGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get hexagon at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")
		radius, _ := cmd.Flags().GetInt64("radius")

		client, err := NewClient(serverAddr, secure)

		x, y, z, _, err := extractHexCoord(args)
		if err != nil {
			return
		}

		result, err := client.MapGet(&hexcli.HexLocation{
			X: x,
			Y: y,
			Z: z,
		}, radius)

		if err != nil {
			fmt.Printf("Error retrieving hexagon information on map: %s", err)
			return
		}

		for _, hex := range result.HexLoc {
			fmt.Printf("%d %d %d %s\n", hex.X, hex.Y, hex.Z, hex.HexID)
		}

	},
}

var mapRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove hexagon at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var hexList hexcli.HexLocationList

		x, y, z, _, err := extractHexCoord(args)
		if err != nil {
			return
		}

		hex := &hexcli.HexLocation{
			X:     x,
			Y:     y,
			Z:     z,
			HexID: "",
		}

		hexList.HexLoc = append(hexList.HexLoc, hex)
		err = client.MapRemove(&hexList)

	},
}

func extractHexCoord(args []string) (x int64, y int64, z int64, id string, err error) {

	x, err = strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Printf("x value is not a number %s : %e", args[0], err)
		return x, y, z, "", err
	}

	y, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		fmt.Printf("y value is not a number %s : %e", args[1], err)
		return x, y, z, "", err
	}

	z, err = strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		fmt.Printf("z value is not a number %s : %e", args[2], err)
		return x, y, z, "", err
	}

	return x, y, z, args[3], nil
}
