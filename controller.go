// +build linux

package kubenet

import (
	"log"
	"net"
	"os/exec"
	"time"

	"github.com/milosgajdos/tenus"
)

type Controller struct {
	thisIP   net.IP
	thisNet  *net.IPNet
	ipStatus map[uint32]bool
	curIP    uint32
	maxIP    uint32

	bridge tenus.Bridger
}

func NewController(firstNet string) *Controller {
	ip, cidr, err := net.ParseCIDR(firstNet)
	if err != nil {
		log.Fatalln(err)
	}
	_, bits := cidr.Mask.Size()
	return &Controller{
		thisIP:   ip,
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
	err = br.SetLinkIp(c.thisIP, c.thisNet)
	handleErr(err)
	err = br.SetLinkUp()
	handleErr(err)
	c.bridge = br
}

func (c *Controller) NewContainer(name, image string) {
	veth, err := tenus.NewVethPairWithOptions(name, tenus.VethOptions{PeerName: "kni0"})
	handleErr(err)

	// same name would reference to the same NIC
	containerNic, err := net.InterfaceByName(name)
	handleErr(err)
	err = c.bridge.AddSlaveIfc(containerNic)
	handleErr(err)

	err = veth.SetLinkUp()
	handleErr(err)
	err = veth.SetPeerLinkUp()
	handleErr(err)

	err = exec.Command(
		"docker", "run",
		"--cap-add", "NET_ADMIN",
		"-d",
		"--name", name,
		image).Run()
	handleErr(err)

	time.Sleep(5 * time.Second)

	pid, err := tenus.DockerPidByName(name, "/var/run/docker.sock")
	handleErr(err)

	err = veth.SetPeerLinkNsPid(pid)
	handleErr(err)

	// FIXME: although this should add default gateway to peer NIC, but get
	// 		error: Unable to set Default gateway: 10.240.0.1 in pid: 22151 network namespace
	err = veth.SetPeerLinkNetInNs(pid, c.NewUniqueIP(), c.thisNet, &c.thisIP)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
