/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/

package cmd

import (
	"net/netip"
)

type Subnet struct {
	Subnet    netip.Prefix
	Name      string
	Vlan      string
	Addresses []Address
}

// HasIP checks if a Subnet already contains given netip.Addr.
// Returns true if the IP already is present, false otherwise.
func (s Subnet) HasIP(ip netip.Addr) bool {
	iscontained := false

	for _, element := range s.Addresses {
		if element.IP.Compare(ip) == 0 {
			iscontained = true
		}
	}

	return iscontained
}

func (s Subnet) RemoveIP(ip netip.Addr) bool {
	return true
}

func (s Subnet) GetIP(ip netip.Addr) bool {
	return true
}

type Address struct {
	IP   netip.Addr
	FQDN string
}
