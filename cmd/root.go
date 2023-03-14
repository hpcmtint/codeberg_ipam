/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ipam",
	Short: "A cli based ipam",
	Long: `A cli based ipam. 
You can manage subnets, single ip addresses within those, and the corresponding A records.
PowerDNS and IPV6-Support will follow`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ipam.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("debug", "d", false, "Enable debug mode. (may print sensitive Information, so please watch out!)")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
}

func initConfig() {
	// Find home directory.
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ipamdir string = homedir + "/.ipam/"
	// Search config in home directory with name ".cobra" (without extension).
	viper.AddConfigPath(ipamdir)
	viper.SetConfigName("ipam")
	viper.SetConfigType("yaml")

	viper.SetDefault("DataPath", ipamdir+"data/")
	viper.SetDefault("EnablePowerDNS", false)
	viper.SetDefault("PowerDNSEndpoint", "")
	viper.SetDefault("PowerDNSApiKey", "")

	if err := viper.ReadInConfig(); err != nil {
		_, patherr := os.Stat(ipamdir)
		if patherr != nil {
			mkerr := os.MkdirAll(ipamdir, 0755)
			if mkerr != nil {
				println("[ERROR] Can't create ipam config directory!", mkerr)
			}
		}

		// I have no idea what's happening here...
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			writeerr := viper.SafeWriteConfig()
			if writeerr != nil {
				println("[ERROR] Can't write config file!", writeerr)
			}
		} else {
			println("[ERROR] Can't read config file!", err)
		}
	}
}
