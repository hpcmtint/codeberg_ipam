/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var ipshowCmd = &cobra.Command{
	Use:     "show",
	Short:   "Show IP and associated name",
	Long:    `Show IP and associated name`,
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ip show called")
	},
}

func init() {
	ipCmd.AddCommand(ipshowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
