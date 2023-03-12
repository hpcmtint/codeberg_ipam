/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var subnetlistCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all subnets",
	Long:    `List all subnets`,
	Aliases: []string{"l"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subnet list called")
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
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
