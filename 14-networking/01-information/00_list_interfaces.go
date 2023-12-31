// example taken from https://github.com/vladimirvivien/go-networking/blob/master/interfaces/lsif.go
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		fmt.Printf("%s: <%s> mtu=%d\n", iface.Name, strings.ToUpper(iface.Flags.String()), iface.MTU)
		if len(iface.HardwareAddr.String()) > 0 {
			fmt.Printf("\tether %s\n", iface.HardwareAddr.String())
		}

		if len(addrs) > 0 {
			for _, addr := range addrs {
				fmt.Printf("\t%s\n", addrInfo(addr))
			}
		}
	}
}

func addrInfo(addr net.Addr) string {
	ipAddr, ipNet, err := net.ParseCIDR(addr.String())
	if err != nil {
		return "unknown"
	}
	var scope string
	switch {
	case ipAddr.IsLoopback():
		scope = "loopback"
	case ipAddr.IsGlobalUnicast():
		scope = "global unicast"
	case ipAddr.IsMulticast():
		scope = "global multicast"
	case ipAddr.IsLinkLocalUnicast():
		scope = "link local unicast"
	case ipAddr.IsLinkLocalMulticast():
		scope = "link local multicast"
	case ipAddr.IsInterfaceLocalMulticast():
		scope = "interface multicast"
	case ipAddr.IsUnspecified():
		scope = "unspecified"
	default:
		scope = "unknown"
	}

	return fmt.Sprintf(
		"%s network=%s addr=%s mask=%v scope=%s",
		ipNet.Network(),
		ipNet.IP.String(),
		ipAddr.String(),
		ipAddr.DefaultMask(),
		scope,
	)
}
