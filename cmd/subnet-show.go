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

// showCmd represents the show command
var subnetshowCmd = &cobra.Command{
	Use:   "show subnet",
	Short: "Displays a subnet.",
	Long: `Displays a subnets details like name and vlan tag, 
aswell as a list of containing IP addresses`,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"s"},
	Example: "ipam subnet show 192.168.0.0/24\nipam subnet show 2001:db8::/64",
	Run: func(cmd *cobra.Command, args []string) {
		net, parseerr := netip.ParsePrefix(args[0])

		if parseerr != nil {
			fmt.Println("[ERROR]", parseerr)
			os.Exit(1)
		}

		if !SubnetExists(net) {
			fmt.Printf("[ERROR] no subnet found for prefix: %v\n", args[0])
			os.Exit(1)
		}

		subnet, subneterr := GetSubnet(net)
		if subneterr != nil {
			fmt.Println("[ERROR]", subneterr)
			os.Exit(1)
		}
		fmt.Printf("\nName:      %v\n", subnet.Name)
		fmt.Printf("Vlan:      %v\n", subnet.Vlan)
		fmt.Printf("Prefix:    %v\n\n", subnet.Subnet)

		fmt.Printf("%v:\n", subnet.Subnet)
		for _, element := range subnet.Addresses {
			if element.FQDN == "" {
				fmt.Printf("\t%v\n", element.IP.String())
			} else {
				fmt.Printf("\t%v:  %v\n", element.IP.String(), element.FQDN)
			}
		}
	},
}

func init() {
	subnetCmd.AddCommand(subnetshowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
