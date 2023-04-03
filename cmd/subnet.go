/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// subnetCmd represents the subnet command
var subnetCmd = &cobra.Command{
	Use:     "subnet",
	Short:   "Manage IP subnets",
	Long:    `Add, delete and show IP subnets`,
	Aliases: []string{"s"},
}

func init() {
	rootCmd.AddCommand(subnetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subnetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subnetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
