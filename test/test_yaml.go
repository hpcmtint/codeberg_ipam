package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Address struct {
	Fqdn string `yaml:"fqdn"`
}

type SubnetFile struct {
	Name      string    `yaml:"name"`
	Vlan      string    `yaml:"vlan"`
	Addresses []Address `yaml:"addresses"`
}

func main() {
	// Read the file
	data, err := ioutil.ReadFile("subnet.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a struct to hold the YAML data
	var asd SubnetFile

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &asd)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data
	fmt.Println(asd)
}
