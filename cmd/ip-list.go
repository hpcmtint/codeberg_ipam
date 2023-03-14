/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var iplistCmd = &cobra.Command{
	Use:     "edit",
	Short:   "List all IP addresses with subnets",
	Long:    `List all IP addresses with subnets`,
	Aliases: []string{"l"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ip list called")
	},
}

func init() {
	ipCmd.AddCommand(iplistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
