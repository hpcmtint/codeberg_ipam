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

		ip := net.ParseIP(ipaddress)

		// Exit if parsed value is no valid IP
		if ip == nil {
			fmt.Printf("[ERROR] not a valid IP: %v\n", ipaddress)
			os.Exit(1)
		}

		// Exit if parsed value is an IPv6 Address
		// TODO: Implement IPv6 support
		if ip.To4() == nil {
			fmt.Printf("[ERROR] IPv6 is not yet supported!\n")
			os.Exit(1)
		}

		// TODO: Check if there is already a subnet that can contain this IP, err if not

		if hostname == "" {
			fmt.Printf("Adding IP %v\n", ipaddress)
		} else {
			fmt.Printf("Adding IP %v with hostname %v\n", ipaddress, hostname)
		}

		// TODO: Save to file
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
