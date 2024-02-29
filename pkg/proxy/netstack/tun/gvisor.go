//go:build linux
// +build linux

package tun

import (
	"github.com/toteb/gvisor/pkg/tcpip/link/fdbased"
	"github.com/toteb/gvisor/pkg/tcpip/link/rawfile"
	"github.com/toteb/gvisor/pkg/tcpip/link/tun"
	"github.com/toteb/gvisor/pkg/tcpip/stack"
)

func Open(tunName string) (stack.LinkEndpoint, error) {
	mtu, err := rawfile.GetMTU(tunName)
	if err != nil {
		return nil, err
	}

	fd, err := tun.Open(tunName)
	if err != nil {
		return nil, err
	}

	linkEP, err := fdbased.New(&fdbased.Options{FDs: []int{fd}, MTU: mtu})
	if err != nil {
		return nil, err
	}
	return linkEP, nil
}
