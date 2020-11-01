// +build linux

package kubenet

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController_NewUniqueIP(t *testing.T) {
	_, n, _ := net.ParseCIDR("10.244.0.0/16")
	ctl := NewController(n)
	for i := 2; i < 258; i++ {
		assert.Equal(t, ipOr(n.IP, intToIP(uint32(i))), ctl.NewUniqueIP())
	}
}
