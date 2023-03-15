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

// deleteCmd represents the delete command
var subnetdeleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete subnet",
	Long:    `Delete a subnet from the ipam.`,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"d"},
	Example: "ipam subnet delete 192.168.0.0/24",
	Run: func(cmd *cobra.Command, args []string) {
		subnet, parseerr := netip.ParsePrefix(args[0])
		if parseerr != nil {
			fmt.Println("[ERROR]", parseerr)
			os.Exit(1)
		}

		if !SubnetExists(subnet) {
			fmt.Printf("[ERROR] Couldn't find subnet %v\n", subnet.String())
			os.Exit(1)
		}

		var confirmation string
		fmt.Printf("[WARNING] Do you really want to delete subnet %v? [y/N] ", subnet.String())
		fmt.Scan(&confirmation)

		if (confirmation == "y") || (confirmation == "Y") {
			deleteerr := DeleteSubnet(subnet)
			if deleteerr != nil {
				fmt.Println("[ERROR]", deleteerr)
				os.Exit(1)
			} else {
				fmt.Printf("deleted subnet %v\n", subnet.String())
			}
		}
	},
}

func init() {
	subnetCmd.AddCommand(subnetdeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
