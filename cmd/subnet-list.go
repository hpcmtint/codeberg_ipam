/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/
package cmd

import (
	"fmt"
	"math"
	"net/netip"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var subnetlistCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all subnets",
	Long:    `List all subnets`,
	Aliases: []string{"l"},
	Args:    cobra.ExactArgs(0),
	Example: "cmdb subnet list",
	Run: func(cmd *cobra.Command, args []string) {
		//verbose, _ := cmd.Flags().GetBool("verbose")
		subnetlist := ListSubnets()
		var subnets []Subnet

		for _, subnet := range subnetlist {
			prefix, _ := netip.ParsePrefix(subnet)
			net, err := GetSubnet(prefix)
			if err != nil {
				fmt.Println("[ERROR]", err)
				os.Exit(1)
			}
			subnets = append(subnets, net)
		}
		sort.Slice(subnets, func(i, j int) bool {
			return subnets[i].Subnet.Addr().Less(subnets[j].Subnet.Addr())
		})

		for _, subnet := range subnets {
			//if verbose {
			var numip, freeip int

			if subnet.Subnet.Addr().Is4() {
				hostbits := float64(32 - subnet.Subnet.Bits())
				numip = int(math.Pow(2, hostbits)) - 2
				freeip = numip - len(subnet.Addresses)

				fmt.Printf("%v:\t%v\t(vl: %v)\tfree: %v\n", subnet.Subnet, subnet.Name, subnet.Vlan, freeip)
			} else {
				fmt.Printf("%v:\t%v\t(vl: %v)\n", subnet.Subnet, subnet.Name, subnet.Vlan)
			}
			//} else {
			//    fmt.Printf("%v:\t%v\t(vl: %v)\n", subnet.Subnet, subnet.Name, subnet.Vlan)
			//}

		}
	},
}

func init() {
	subnetCmd.AddCommand(subnetlistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subnetlistCmd.Flags().BoolP("verbose", "v", false, "Show verbose output like free IPs")
}
