package main

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPOr(t *testing.T) {
	assert.Equal(t,
		net.IPv4(10, 244, 0, 1),
		ipOr(net.IPv4(10, 244, 0, 0), net.IPv4(0, 0, 0, 1)))
}
