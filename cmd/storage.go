/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/

package cmd

import (
	"bufio"
	"fmt"
	"net/netip"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Subnet struct {
	Subnet    netip.Prefix
	Name      string
	Vlan      string
	Addresses []Address
}

type Address struct {
	IP   string
	FQDN string
}

// SearchBestSubnet tries to load the most fitting IP subnet file
// on disk. It takes an IP object and tries to get the best subnet
// (meaning the subnet with the smallest subnet size).
//
// Returns the best subnet as Subnet object and true if a suitable
// subnet was found, otherwise an empty Subnet object and false.
func SearchBestSubnet(ip netip.Addr) (Subnet, bool) {
	subnets := ListSubnets()
	var smallestprefix int = 0
	bestmatch, _ := netip.ParsePrefix("0.0.0.0/32")
	var isipv4 bool = ip.Is4()
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
	var datadir string = viper.GetString("DataPath")

	subnetfiles, readerr := os.ReadDir(datadir)

	if len(subnetfiles) == 0 {
		return subnets
	}

	if readerr != nil {
		fmt.Println("[ERROR]", readerr)
		os.Exit(1)
	}
	for _, element := range subnetfiles {
		subnets = append(subnets, strings.Replace(element.Name(), "_", "/", 1))
	}

	return subnets
}

// WriteSubnet takes a given Subnet object and tries to write it to
// file.
//
// Returns nil on success or the error that happened.
func WriteSubnet(subnet Subnet) error {
	var datadir string = viper.GetString("DataPath")

	_, direrr := os.Stat(datadir)
	if direrr != nil {
		mkerr := os.MkdirAll(datadir, 0755)
		if mkerr != nil {
			println("[ERROR] Can't create ipam config directory!", mkerr)
			os.Exit(1)
		}
	}

	filename := datadir + strings.Replace(subnet.Subnet.String(), "/", "_", 1)

	file, fileerr := os.Create(filename)
	if fileerr != nil {
		fmt.Println("[ERROR]", fileerr)
		os.Exit(1)
	}
	defer file.Close()

	_, suberr := file.WriteString(subnet.Subnet.String() + "\n")
	if suberr != nil {
		fmt.Println("[ERROR]", suberr)
		os.Exit(1)
	}

	_, nameerr := file.WriteString(subnet.Name + "\n")
	if nameerr != nil {
		fmt.Println("[ERROR]", nameerr)
		os.Exit(1)
	}

	_, vlanerr := file.WriteString(subnet.Vlan + "\n")
	if vlanerr != nil {
		fmt.Println("[ERROR]", vlanerr)
		os.Exit(1)
	}

	if len(subnet.Addresses) != 0 {
		for _, element := range subnet.Addresses {
			_, err := file.WriteString(element.IP + ":" + element.FQDN + "\n")
			if err != nil {
				fmt.Println("[ERROR]", err)
				os.Exit(1)
			}
		}
	}

	return nil
}

// GetSubnet reads the corresponding file for the given
// netip.Prefix net and constructs a Subnet object.
//
// Returns the Subnet object and nil if the file read was
// successful, an empty Subnet object and the error otherwise.
func GetSubnet(net netip.Prefix) (Subnet, error) {
	var datadir string = viper.GetString("DataPath")
	filename := datadir + strings.Replace(net.String(), "/", "_", 1)
	var subnet Subnet = Subnet{}

	// open file
	file, openerr := os.Open(filename)
	if openerr != nil {
		return Subnet{}, openerr
	}
	// remember to close the file at the end of the program
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	var counter int = 0
	for scanner.Scan() {
		switch counter {
		case 0:
			subnet.Subnet, _ = netip.ParsePrefix(scanner.Text())

		case 1:
			subnet.Name = scanner.Text()

		case 2:
			subnet.Vlan = scanner.Text()

		default:
			s := strings.Split(scanner.Text(), ":")
			a := Address{s[0], s[1]}
			subnet.Addresses = append(subnet.Addresses, a)
		}
		counter = counter + 1
	}

	if scanerr := scanner.Err(); scanerr != nil {
		return Subnet{}, openerr
	}

	return subnet, nil
}
