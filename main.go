package main

import (
	"log"
	"net"
)

func main() {
	_, n, err := net.ParseCIDR("10.240.0.1/24")
	if err != nil {
		log.Fatalln(err)
	}
	ctl := NewController(n)
	ctl.SetupEnv()

	ctl.NewContainer("test1")
}
