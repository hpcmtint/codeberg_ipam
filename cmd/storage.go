/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/

package cmd

import (
	"encoding/json"
	"fmt"
	"net/netip"
	"os"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

// FindBestSubnet tries to load the most fitting IP subnet file
// on disk. It takes an IP object and tries to get the best subnet
// (meaning the subnet with the smallest subnet size).
//
// Returns the best subnet as Subnet object and true if a suitable
// subnet was found, otherwise an empty Subnet object and false.
func FindBestSubnet(ip netip.Addr) (Subnet, bool) {
	subnets := ListSubnets()
	var smallestprefix = 0
	bestmatch, _ := netip.ParsePrefix("::/128")
	var isipv4 = ip.Is4()
	var subnet Subnet

	for _, net := range subnets {
		prefix, _ := netip.ParsePrefix(net)
		if prefix.Addr().Is4() == isipv4 {
			if prefix.Contains(ip) {
				if prefix.Bits() > smallestprefix {
					bestmatch = prefix
				}
			}
		}
	}
	if !bestmatch.Addr().IsUnspecified() {
		var geterr error
		subnet, geterr = GetSubnet(bestmatch)
		if geterr != nil {
			fmt.Println("[ERROR]", geterr)
			os.Exit(1)
		}
		return subnet, true
	} else {
		return Subnet{}, false
	}
}

// SubnetExists will return true if the given subnet already exists
// on file, false otherwise.
func SubnetExists(net netip.Prefix) bool {
	subnets := ListSubnets()

	for _, b := range subnets {
		if b == net.String() {
			return true
		}
	}
	return false
}

// ListSubnets returns a list of strings containing the current
// subnets configured.
func ListSubnets() []string {
	subnets := make([]string, 0)
	var datadir = viper.GetString("DataPath")

	subnetfiles, readerr := os.ReadDir(datadir)

	if len(subnetfiles) == 0 {
		return subnets
	}

	if readerr != nil {
		fmt.Println("[ERROR]", readerr)
		os.Exit(1)
	}
	for _, element := range subnetfiles {
		a := strings.Replace(element.Name(), "_", "/", 1)
		a = strings.Replace(a, ".json", "", 1)
		subnets = append(subnets, a)
	}

	return subnets
}

// WriteSubnet takes a given Subnet object and tries to write it to
// file.
//
// Returns nil on success or the error that happened.
func (s Subnet) WriteSubnet() error {
	var datadir = viper.GetString("DataPath")

	_, direrr := os.Stat(datadir)
	if direrr != nil {
		mkerr := os.MkdirAll(datadir, 0755)
		if mkerr != nil {
			println("[ERROR] Can't create ipam config directory!", mkerr)
			os.Exit(1)
		}
	}

	filename := datadir + strings.Replace(s.Subnet.String(), "/", "_", 1) + ".json"

	data, _ := json.Marshal(s)

	file, fileerr := os.Create(filename)
	if fileerr != nil {
		fmt.Println("[ERROR]", fileerr)
		os.Exit(1)
	}
	defer file.Close()

	_, writeerr := file.Write(data)
	if writeerr != nil {
		fmt.Println("[ERROR]", writeerr)
		os.Exit(1)
	}

	return nil
}

// GetSubnet reads the corresponding file for the given
// netip.Prefix net and constructs a Subnet object.
//
// Returns the Subnet object and nil if the file read was
// successful, an empty Subnet object and the error otherwise.
func GetSubnet(net netip.Prefix) (Subnet, error) {
	var datadir = viper.GetString("DataPath")
	filename := datadir + strings.Replace(net.String(), "/", "_", 1) + ".json"
	var subnet = Subnet{}

	content, readerr := os.ReadFile(filename)
	if readerr != nil {
		return Subnet{}, readerr
	}

	marsherr := json.Unmarshal(content, &subnet)
	if marsherr != nil {
		return Subnet{}, marsherr
	}

	return subnet, nil
}

// SortAddresses sorts the given list of IP addresses
// using netip.Addr.Less() and returns the sorted slice.
func SortAddresses(list []Address) []Address {

	if len(list) <= 1 {
		return list
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].IP.Less(list[j].IP)
	})
	return list
}

// DeleteSubnet deletes the subnet file on disk for netip.Prefix
// net.
//
// Returns nil on success, or a *PathError on failure
func DeleteSubnet(net netip.Prefix) error {
	var datadir = viper.GetString("DataPath")
	filename := datadir + strings.Replace(net.String(), "/", "_", 1) + ".json"

	removeerr := os.Remove(filename)
	if removeerr != nil {
		return removeerr
	} else {
		return nil
	}
}
