package hexclient

import (
	"fmt"
	"github.com/3vilm33pl3/hexcli/internal/pkg/hexcli"
	"github.com/spf13/cobra"
	"strconv"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manipulate hexagons",
	Long:  `A longer description that spans multiple lines and likely contains examples.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("repo called")
	},
}

var repoAddCmd = &cobra.Command{
	Use:   "add [ref] [exits]",
	Short: "add hexagon to repository with reference [ref] and [exits]",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var infoList hexcli.HexInfoList

		ref := args[0]
		exits, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error parsing exits %s", err)
		}

		infoList.HexInfo = append(infoList.HexInfo, &hexcli.HexInfo{ID: ref, Exits: uint32(exits)})

		err = client.RepoAddHexagon(&infoList)

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
		var infoList hexcli.HexIDList

		for _, id := range args {
			infoList.HexID = append(infoList.HexID, id)
		}

		err = client.RepoDeleteHexagon(&infoList)

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}
	},
}
