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
Coming soon

## How can I contribute to the development?
Patches, bug reports or feature suggestions can be sent by email to
[~lauralani/ipam@lists.sr.ht](mailto:~lauralani/ipam@lists.sr.ht). If you're
not familiar with sending patches over email see
[here](https://git-send-email.io/).

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
Copyright 2022 Laura Kalb <dev (at) lauka.net>

This program is Free Software: You can use, study share and improve it at your
will. Specifically you can redistribute and/or modify it under the terms of the
[GNU General Public License](https://www.gnu.org/licenses/gpl.html) as
published by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
