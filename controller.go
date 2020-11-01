// +build linux

package kubenet

import (
	"log"
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
