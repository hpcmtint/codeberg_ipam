/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/

package cmd

import (
	"errors"
	"net/netip"
	"time"
)

type Subnet struct {
	Subnet    netip.Prefix `json:"subnet"`
	Name      string       `json:"name"`
	Vlan      string       `json:"vlan"`
	ChangedAt time.Time    `json:"changedat,omitempty"`
	ChangedBy string       `json:"changedby,omitempty"`
	Addresses []Address    `json:"addresses"`
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

// RemoveIP removes the Address object for given ip from
// the Address list of the subnet.
//
// Returns the changed Subnet and nil if delete was
// successful, or an empty Subnet and an error if
// ip could not be deleted.
func (s Subnet) RemoveIP(ip netip.Addr) (Subnet, error) {
	var addrlist []Address

	if !s.HasIP(ip) {

		return Subnet{}, errors.New("IP " + ip.String() + " wasn't found in subnet " + s.Subnet.String())
	}

	for _, item := range s.Addresses {
		if item.IP.Compare(ip) != 0 {
			addrlist = append(addrlist, item)
		}
	}
	s.Addresses = addrlist
	return s, nil
}

// GetIP returns the Address object for the subnet with
// netip.Addr ip.
//
// Returns the Address object and true if a corresponding
// object was found, an empty Address and false otherwise.
func (s Subnet) GetIP(ip netip.Addr) (Address, bool) {
	for _, item := range s.Addresses {
		if item.IP.Compare(ip) == 0 {
			return item, true
		}
	}

	return Address{}, false
}

type Address struct {
	IP        netip.Addr `json:"ip"`
	FQDN      string     `json:"fqdn"`
	ChangedAt time.Time  `json:"changedat,omitempty"`
	ChangedBy string     `json:"changedby,omitempty"`
}
