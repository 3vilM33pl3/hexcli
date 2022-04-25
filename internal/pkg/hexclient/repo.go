package hexclient

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/3vilm33pl3/hexcli/internal/pkg/hexcli"
	"github.com/spf13/cobra"
	"log"
	"os"
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

		err = client.RepoAddHexagonInfo(&infoList)
		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

	},
}

var repoAddFileCmd = &cobra.Command{
	Use:   "file [filename]",
	Short: "add hexagon from file [filename] to repository",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var infoList hexcli.HexInfoList

		f, err := os.Open(args[0])
		if err != nil {
			fmt.Printf("Error opening file %s", args[0])
			return
		}
		rc := bufio.NewReader(f)
		csvLines, err := csv.NewReader(rc).ReadAll()
		if err != nil {
			log.Printf("Error reading hexdata file: %v", err)
			return
		}

		for i, line := range csvLines {
			exits, err := strconv.Atoi(line[1])
			if err != nil {
				fmt.Printf("Skipping line %d: %s\n", i, err)
			}
			hexInfo := &hexcli.HexInfo{
				ID:    line[0],
				Exits: uint32(exits),
			}
			infoList.HexInfo = append(infoList.HexInfo, hexInfo)
		}

		err = client.RepoAddHexagonInfo(&infoList)
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

		err = client.RepoDeleteHexagonInfo(&infoList)

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}
	},
}

var repoGetCmd = &cobra.Command{
	Use:   "get [ref]",
	Short: "get hexagon info from repository with reference [ref] ",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)
		var infoList hexcli.HexIDList

		for _, id := range args {
			infoList.HexID = append(infoList.HexID, id)
		}

		result, err := client.RepoGetHexagonInfo(&infoList)

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		for _, hex := range result.HexInfo {
			fmt.Printf("%s %d \n", hex.ID, hex.Exits)
		}
	},
}

var repoGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "get all hexagon info from repository",
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr, _ := cmd.Flags().GetString("addr")
		secure, _ := cmd.Flags().GetBool("secure")

		client, err := NewClient(serverAddr, secure)

		result, err := client.RepoGetAllHexagonInfo()

		if err != nil {
			fmt.Printf("Error connecting %s", err)
			return
		}

		for _, hex := range result.HexInfo {
			fmt.Printf("%s %d \n", hex.ID, hex.Exits)
		}
	},
}
