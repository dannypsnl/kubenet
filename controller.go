package main

import (
	"encoding/binary"
	"net"
)

type Controller struct {
	thisNet  *net.IPNet
	ipStatus map[uint32]bool
	curIP    uint32
	maxIP    uint32
}

func NewController(cidr *net.IPNet) *Controller {
	_, bits := cidr.Mask.Size()
	return &Controller{
		thisNet:  cidr,
		ipStatus: make(map[uint32]bool),
		curIP:    1,
		maxIP:    2 << (bits - 1),
	}
}

func (c *Controller) NewUniqueIP() net.IP {
	if c.curIP == c.maxIP {
		c.curIP = 1
	}
	if !c.ipStatus[c.curIP] {
		c.ipStatus[c.curIP] = true
		return ipOr(c.thisNet.IP, intToIP(c.curIP))
	} else {
		c.curIP++
		return c.NewUniqueIP()
	}
}

func ipOr(a, b net.IP) net.IP {
	return net.IPv4(a.To4()[0]|b.To4()[0], a.To4()[1]|b.To4()[1], a.To4()[2]|b.To4()[2], a.To4()[3]|b.To4()[3])
}
func intToIP(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
