package main

import (
	"net"

	"github.com/milosgajdos/tenus"
)

type Controller struct {
	thisNet  *net.IPNet
	ipStatus map[uint32]bool
	curIP    uint32
	maxIP    uint32

	bridge tenus.Bridger
}

func NewController(cidr *net.IPNet) *Controller {
	_, bits := cidr.Mask.Size()
	return &Controller{
		thisNet:  cidr,
		ipStatus: make(map[uint32]bool),
		curIP:    2,
		maxIP:    2 << (bits - 1),
	}
}

func (c *Controller) NewUniqueIP() net.IP {
	if c.curIP == c.maxIP {
		c.curIP = 2
	}
	if !c.ipStatus[c.curIP] {
		c.ipStatus[c.curIP] = true
		return ipOr(c.thisNet.IP, intToIP(c.curIP))
	} else {
		c.curIP++
		return c.NewUniqueIP()
	}
}
