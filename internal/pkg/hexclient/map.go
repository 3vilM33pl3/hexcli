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
	Use:   "add [x,y,z] [direction] [content ref]",
	Short: "add hexagon with coordinate [x,y,z] and content reference [content ref]",
	Long: "Example: nb hex place --secure=false -- 0 -5 5 N 0000-0000-0000-0000" +
		"The double dash is needed to indicate no more flags are coming and everything is interpreted as an normal argument. " +
		"This is needed so that negative numbers don't get interpreted as flags",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)

		x, y, z, err := extractHexCoord(args)
		if err != nil {
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

		err = client.MapAdd(hex)

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

		x, y, z, err := extractHexCoord(args)
		if err != nil {
			return
		}

		result, err := client.MapGet(&hexcli.Hex{
			X: x,
			Y: y,
			Z: z,
		}, radius)

		if err != nil {
			fmt.Printf("Error retrieving hexagon information on map: %s", err)
			return
		}

		for _, hex := range result.Hex {
			fmt.Printf("%d %d %d %s %s\n", hex.X, hex.Y, hex.Z, hex.Direction, hex.Reference)
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
		var refList hexcli.HexRefList

		for _, ref := range args {
			refList.Ref = append(refList.Ref, &hexcli.HexReference{Ref: ref})
		}

		var hexList hexcli.HexList

		x, y, z, err := extractHexCoord(args)
		if err != nil {
			return
		}

		hex := &hexcli.Hex{
			X:         x,
			Y:         y,
			Z:         z,
			Direction: hexcli.Direction_N,
			Reference: "",
		}

		hexList.Hex = append(hexList.Hex, hex)
		err = client.MapRemove(&hexList)

	},
}

func extractHexCoord(args []string) (x int64, y int64, z int64, err error) {
	x, err = strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Printf("x value is not a number %s : %e", args[0], err)
		return x, y, z, err
	}

	y, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		fmt.Printf("y value is not a number %s : %e", args[1], err)
		return x, y, z, err
	}

	z, err = strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		fmt.Printf("z value is not a number %s : %e", args[2], err)
		return x, y, z, err
	}
	return x, y, z, nil
}
