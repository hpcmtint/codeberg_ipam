# ipam
A cli based [IP Address Management](https://en.wikipedia.org/wiki/IP_address_management)
built with [Go](https://go.dev/) using [Cobra](https://cobra.dev/) and
[Viper](https://github.com/spf13/viper) with optional [PowerDNS](https://www.powerdns.com/auth.html)
API integration.

## Features
`ipam` is a cli based IPAM. It supports:
- Adding, listing, viewing and deleting subnets
- Adding, viewing and deleting single IP addresses with FQDNs
- automatic addition and deletion of forward and reverse records
- import and export configuration

## But why?
In my career as a Network Engineer I worked at a company that developed their own internal CMDB
that was set up as a cli application. I really liked the ease of use and the way you could
quickly do stuff.

So I sat down and started my first Go project and started developing this ipam.

## How do I use the cmdb?
`ipam` is built as a cli app. Once you installed `ipam` and added it to your `$PATH`, you can
- add prefixes with `ipam subnet add 192.168.0.0/24 fancy-subnet-name`
- add IPs to those prefixes with `ipam ip add 192.168.0.1 fqdn.example.com`
- add the next free IP in a prefix with `ipam ip add 192.168.0.0/24 fqdn2.example.com`
- much more (see `ipam (command) --help` for examples)


If PowerDNS integration is enabled in `$HOME/.ipam/ipam.yml`, forward and reverse DNS records
are automatically managed when you add or delete IP addresses with hostnames.

## How can I contribute to the development?
The main development is happening at https://codeberg.org/lauralani/ipam.

If you find a bug or would like a feature to get added you can 
[create a new issue](https://codeberg.org/lauralani/ipam/issues)
or collaborate by forking the repository and 
[create a new pull request](https://codeberg.org/lauralani/ipam/pulls).

If you have any questions please also feel free to create a new Issue!

## CLI
```
❯ ipam
A cli based ipam.
You can manage subnets and single IP addresses within those subnets.
ipam can also manage the corresponding DNS records in your PowerDNS Instance.

Usage:
ipam [command]

Available Commands:
completion  Generate the autocompletion script for the specified shell
export      Export ipam configuration
help        Help about any command
import      Import ipam configuration
ip          manage IP addresses
subnet      Manage IP subnets

Flags:
-d, --debug     Enable debug mode. (may print sensitive Information, so please watch out!)
-h, --help      help for ipam
-v, --version   version for ipam

Use "ipam [command] --help" for more information about a command.
```

## License
Copyright 2023 Laura Kalb <dev (at) lauka.net>

This program is Free Software: You can use, study share and improve it at your
will. Specifically you can redistribute and/or modify it under the terms of the
[GNU General Public License](https://www.gnu.org/licenses/gpl.html) as
published by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
