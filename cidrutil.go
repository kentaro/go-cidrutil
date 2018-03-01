package cidrutil

import (
	"net"
)

type Parser struct {
}

type Network struct {
	CIDR string
	Ones int
	Bits int
	Network string
	Hosts []string
	Broadcast string
}

// NewParser returns a new `*Parser`.
func NewParser() (p *Parser) {
	p = &Parser{}
	return
}

// Parse parses a CIDR string and returns a `*Network`.
func (p *Parser) Parse(cidr string) (n *Network, err error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	n = &Network{CIDR: cidr}
	n.Ones, n.Bits = ipnet.Mask.Size()

	if n.Ones == 32 {
		n.Hosts = []string{ip.String()}
	} else {
		ips := make([]string, 0)
		for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inclement(ip) {
			ips = append(ips, ip.String())
		}

		if n.Ones == 31 {
			n.Hosts = ips
		} else {
			n.Network = ips[0]
			n.Hosts = ips[1:len(ips)-1]
			n.Broadcast = ips[len(ips)-1]
		}
	}

	return
}

// Borrowed from http://play.golang.org/p/m8TNTtygK0
func inclement(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
