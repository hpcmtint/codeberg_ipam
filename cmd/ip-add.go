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

var ipaddCmd = &cobra.Command{
	Use:     "add ipaddress [hostname]",
	Short:   "Add new IP address",
	Long:    `Add new IP address`,
	Aliases: []string{"a"},
	Args:    cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var ipaddress, hostname string

		if len(args) == 1 {
			ipaddress = args[0]
			hostname = ""
		} else {
			ipaddress = args[0]
			hostname = args[1]
		}

		ip, parseerr := netip.ParseAddr(ipaddress)

		// Exit if parsed value is no valid IP
		if parseerr != nil {
			fmt.Println("[ERROR]", parseerr)
			os.Exit(1)
		}

		// Exit if parsed value is an IPv6 Address
		// TODO: Implement IPv6 support
		if !ip.Is4() {
			fmt.Printf("[ERROR] IPv6 is not yet supported!\n")
			os.Exit(1)
		}

		subnet, subnetexists := FindBestSubnet(ip)

		if !subnetexists {
			fmt.Printf("[ERROR] Found no suitable subnet for IP %v\n", ipaddress)
			fmt.Printf("[ERROR] Maybe you need to add it first?\n")
			os.Exit(1)
		}

		if subnet.HasIP(ip) {
			fmt.Printf("[ERROR] IP %v already exists in subnet %v\n", ip.String(), subnet.Subnet.String())
			os.Exit(1)
		}

		subnet.Addresses = append(subnet.Addresses, Address{ip, hostname})

		writeerr := WriteSubnet(subnet)
		if writeerr != nil {
			fmt.Println("[ERROR]", writeerr)
			os.Exit(1)
		}

		if hostname == "" {
			fmt.Printf("added ip:\nip: %v\n", ipaddress)
		} else {
			fmt.Printf("added ip:\nip: %v\nhostname: %v\n", ipaddress, hostname)
			dnserr := AddDNSFqdn(hostname, ip)
			if dnserr != nil {
				fmt.Println("[ERROR]", writeerr)
				os.Exit(1)
			}
		}
	},
}

func init() {
	ipCmd.AddCommand(ipaddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
