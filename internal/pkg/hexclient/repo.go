package hexclient

import (
	"fmt"
	"github.com/3vilm33pl3/hexcli/internal/pkg/hexcli"
	"github.com/spf13/cobra"
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
