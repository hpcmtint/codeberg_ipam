/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ipeditCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit an IP address",
	Long:    `Edit an IP address`,
	Aliases: []string{"e"},
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not implemented yet; please delete and readd")
	},
}

func init() {
	ipCmd.AddCommand(ipeditCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
