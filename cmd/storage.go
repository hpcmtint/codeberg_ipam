/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/

package cmd

import (
	"net"
)

type Subnet struct {
	Subnet    net.IPNet
	Name      string
	Vlan      string
	Addresses []Address
}

type Address struct {
	IP   string
	FQDN string
}

// GetSubnet tries to load the most fitting IP subnet file on disk.
// It takes an IP object and tries to get the best subnet (meaning
// the subnet with the smallest subnet size).
//
// Returns the best subnet as Subnet object and true if a suitable
// subnet was found, otherwise an empty Subnet object and false.
func GetSubnet(ip net.IP) (Subnet, bool) {

	return Subnet{}, false
}

// WriteSubnet takes a given Subnet object and tries to write it to
// file.
//
// Returns nil on success or the error that happened.
func WriteSubnet(subnet Subnet) error {
	//if subnet.Subnet

	return nil
}
