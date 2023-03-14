/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/

package cmd

import (
	"fmt"
	"net/netip"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var subnetaddCmd = &cobra.Command{
	Use:     "add subnet subnet-name [vlan]",
	Short:   "Add a new subnet",
	Long:    `Add a new subnet`,
	Args:    cobra.RangeArgs(2, 3),
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		var subnet string
		var vlanid string
		var netname string

		if len(args) == 2 {
			subnet = args[0]
			netname = args[1]
			vlanid = "-"
		}

		if len(args) == 3 {
			subnet = args[0]
			netname = args[1]
			vlanid = args[2]
		}

		// Parse subnet into ParseCIDR to test if it's a valid subnet
		// _, ipnet, err := net.ParseCIDR(subnet)
		ipnet, parseerr := netip.ParsePrefix(subnet)

		// Exit if subnet already exists, no need to add it then
		if SubnetExists(ipnet) {
			fmt.Printf("[ERROR] Subnet already exists: %v\n", subnet)
			os.Exit(1)
		}

		// Exit if parsed value is no valid IP
		if parseerr != nil {
			fmt.Println("[ERROR]", parseerr)
			os.Exit(1)
		}

		// Exit if parsed value is an IPv6 Address
		// TODO: Implement IPv6 support
		if !ipnet.Addr().Is4() {
			fmt.Printf("[ERROR] IPv6 is not yet supported!\n")
			os.Exit(1)
		}

		subnetobject := Subnet{}
		subnetobject.Subnet = ipnet
		subnetobject.Name = netname
		subnetobject.Vlan = vlanid

		writeerr := WriteSubnet(subnetobject)

		if writeerr != nil {
			fmt.Println("[ERROR]", writeerr)
			os.Exit(1)
		}

		fmt.Printf("added subnet:\nnet: %v\nname: %v\nvlan: %v\n", subnet, netname, vlanid)
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
