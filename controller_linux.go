package main

import (
	"log"
	"net"

	"github.com/milosgajdos/tenus"
)

func (c *Controller) SetupEnv() {
	br, err := tenus.NewBridgeWithName("kube-bridge")
	handleErr(err)
	err = br.SetLinkIp(c.thisNet.IP, c.thisNet)
	handleErr(err)
	err = br.SetLinkUp()
	handleErr(err)
	c.bridge = br
}

func (c *Controller) NewContainer(name string) {
	veth, err := tenus.NewVethPairWithOptions(name, tenus.VethOptions{PeerName: "kni0"})
	handleErr(err)
	err = veth.SetLinkIp(c.NewUniqueIP(), c.thisNet)
	handleErr(err)

	// same name would reference to the same NIC
	containerNic, err := net.InterfaceByName(name)
	handleErr(err)
	err = c.bridge.AddSlaveIfc(containerNic)
	handleErr(err)

	err = veth.SetLinkUp()
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
