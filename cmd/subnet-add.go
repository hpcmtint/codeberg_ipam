/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/

package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var subnetaddCmd = &cobra.Command{
	Use:     "add subnet [vlan]",
	Short:   "Add a new subnet",
	Long:    `Add a new subnet`,
	Args:    cobra.RangeArgs(1, 2),
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		var subnet string
		var vlanid string
		var netname string

		if len(args) == 1 {
			subnet = args[0]

			vlanid = "-"
		}

		if len(args) == 2 {
			subnet = args[0]
			vlanid = args[1]
		}

		// Parse subnet into ParseCIDR to test if it's a valid subnet
		_, ipnet, err := net.ParseCIDR(subnet)

		// Exit if parsed value is no valid IP
		if err != nil {
			fmt.Println("[ERROR]", err)
			os.Exit(1)
		}

		// Exit if parsed value is an IPv6 Address
		// TODO: Implement IPv6 support
		if ipnet.IP.To4() == nil {
			fmt.Printf("[ERROR] IPv6 is not yet supported!\n")
			os.Exit(1)
		}

		// Ask for Subnet Name
		// TODO: Check if net name only contains letters, numbers and hyphens
		fmt.Printf("Subnet name: ")
		fmt.Scan(&netname)

		if vlanid == "-" {
			fmt.Printf("Adding Subnet %v.\n", subnet)
		} else {
			fmt.Printf("Adding Subnet %v with VLAN Tag %v.\n", subnet, vlanid)
		}

		// TODO: Save subnet to file
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
