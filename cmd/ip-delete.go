/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var ipdeleteCmd = &cobra.Command{
	Use:     "delete ipaddress",
	Short:   "Delete an IP address",
	Long:    `Delete an IP address`,
	Aliases: []string{"d"},
	Args:    cobra.ExactArgs(1),
	Example: "ipam ip delete 192.168.0.1",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	ipCmd.AddCommand(ipdeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
