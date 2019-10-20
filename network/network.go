package network

import (
	"net"
)

// Lookup to call net.LookupIP
func Lookup(ip string) (net.IP, error) {
	ips, err := net.LookupIP(ip)
	if err != nil {
		return nil, err
	}
	return ips[0], nil
}
