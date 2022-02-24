package hexclient

import (
	"fmt"
	"github.com/spf13/cobra"
)

var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Manipulate hexagons",
	Long:  `A longer description that spans multiple lines and likely contains examples.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hex called")
	},
}

var hexInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "show information of hexagon at `coord` [x,y,z]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex info command: %v\n", args)
	},
}

var hexPackageCmd = &cobra.Command{
	Use:   "pack [dir]",
	Short: "packagae content in directory [dir]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside hex info command: %v\n", args)
	},
}
