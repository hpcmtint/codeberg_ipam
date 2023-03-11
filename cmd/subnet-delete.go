/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var subnetdeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete subnet",
	Long:  `Delete a subnet from the cmdb.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: Too few arguments!")
			fmt.Print("Usage:\n  cmdb subnet delete [subnet]")
			os.Exit(1)
		}
		if len(args) > 1 {
			fmt.Println("Error: Too many arguments!")
			fmt.Print("Usage:\n  cmdb subnet delete [subnet]")
			os.Exit(1)
		}
		subnet := args[0]

		fmt.Printf("Deleting %v\n", subnet)
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
