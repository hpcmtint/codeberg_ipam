/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var subnetaddCmd = &cobra.Command{
	Use:   "add [subnet] [vlan]",
	Short: "Add a new subnet",
	Long: `Add a new subnet`,
	Run: func(cmd *cobra.Command, args []string) {
		var subnet string
		var vlanid string

		if len(args) == 0 {
			fmt.Printf("IP Subnet (Format: 192.168.0.0/24): ")
			fmt.Scan(&subnet)

			fmt.Printf("VLAN Tag: ")
			fmt.Scan(&vlanid)
		}

		if len(args) == 1 {
			subnet = args[0]

			fmt.Printf("VLAN Tag: ")
			fmt.Scan(&vlanid)
		}

		if len(args) == 2 {
			subnet = args[0]
			vlanid = args[1]
		}

		fmt.Printf("Adding Subnet %v with VLAN Tag %v.\n", subnet, vlanid)
	},
}

func init() {
	subnetCmd.AddCommand(subnetaddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
