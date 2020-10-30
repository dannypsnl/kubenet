package main

import (
	"net"
)

type Controller struct {
	cidr   net.IPMask
	ipList []net.IP
}

func NewController(cidr net.IPMask) *Controller {
	return &Controller{
		cidr:   cidr,
		ipList: make([]net.IP, 0),
	}
}
