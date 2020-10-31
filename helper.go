package main

import (
	"encoding/binary"
	"net"
)

func ipOr(a, b net.IP) net.IP {
	return net.IPv4(a.To4()[0]|b.To4()[0], a.To4()[1]|b.To4()[1], a.To4()[2]|b.To4()[2], a.To4()[3]|b.To4()[3])
}
func intToIP(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
