// +build linux

package kubenet

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController_NewUniqueIP(t *testing.T) {
	ctl := NewController("10.244.0.0/16")
	for i := 2; i < 258; i++ {
		assert.Equal(t, ipOr(net.ParseIP("10.244.0.0"), intToIP(uint32(i))), ctl.NewUniqueIP())
	}
}
