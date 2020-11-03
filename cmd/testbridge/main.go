package main

import (
	"net"

	"github.com/docker/libcontainer/netlink"
	"github.com/milosgajdos/tenus"
)

func main() {
	ip, n, err := net.ParseCIDR("10.240.0.1/24")

	bridge, err := tenus.NewBridgeWithName("kube-bridge")
	handleErr(err)
	err = bridge.SetLinkIp(ip, n)
	handleErr(err)
	err = bridge.SetLinkUp()
	handleErr(err)

	name := "testveth"
	veth, err := tenus.NewVethPairWithOptions(name, tenus.VethOptions{PeerName: "kni0"})
	handleErr(err)

	vethip, vethNet, err := net.ParseCIDR("10.240.0.2/24")
	err = veth.SetLinkIp(vethip, vethNet)
	handleErr(err)

	// same name would reference to the same NIC
	containerNic, err := net.InterfaceByName(name)
	handleErr(err)
	err = bridge.AddSlaveIfc(containerNic)
	handleErr(err)

	err = netlink.AddDefaultGw(ip.String(), containerNic.Name)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
