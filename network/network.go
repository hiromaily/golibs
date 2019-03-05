package network

import (
	"net"
)

func Lookup(ip string) (net.IP, error) {
	ips, err := net.LookupIP(ip)
	if err != nil {
		return nil, err
	}
	return ips[0], nil
}
