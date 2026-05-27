// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"net"

	"github.com/pion/logging"
	"github.com/pion/transport/v4"
)

// MultiUDPMuxDefault implements both UDPMux and AllConnsGetter,
// allowing users to pass multiple UDPMux instances to the ICE agent
// configuration.
type MultiUDPMuxDefault struct {
	muxes          []UDPMux
	localAddrToMux map[string]UDPMux
}

// NewMultiUDPMuxDefault creates an instance of MultiUDPMuxDefault that
// uses the provided UDPMux instances.
func NewMultiUDPMuxDefault(muxes ...UDPMux) *MultiUDPMuxDefault {
	_ = "STUB: not implemented"
	return nil
}

// GetConn returns a PacketConn given the connection's ufrag and network
// creates the connection if an existing one can't be found.
func (m *MultiUDPMuxDefault) GetConn(ufrag string, addr net.Addr) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), nil
}

// RemoveConnByUfrag stops and removes the muxed packet connection
// from all underlying UDPMux instances.
func (m *MultiUDPMuxDefault) RemoveConnByUfrag(ufrag string) { _ = "STUB: not implemented"; return }

// Close the multi mux, no further connections could be created.
func (m *MultiUDPMuxDefault) Close() error { _ = "STUB: not implemented"; return nil }

// GetListenAddresses returns the list of addresses that this mux is listening on.
func (m *MultiUDPMuxDefault) GetListenAddresses() []net.Addr { _ = "STUB: not implemented"; return nil }

// NewMultiUDPMuxFromPort creates an instance of MultiUDPMuxDefault that
// listen all interfaces on the provided port.
func NewMultiUDPMuxFromPort(port int, opts ...UDPMuxFromPortOption) (*MultiUDPMuxDefault, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil, nil
}

// UDPMuxFromPortOption provide options for NewMultiUDPMuxFromPort.
type UDPMuxFromPortOption interface {
	apply(*multiUDPMuxFromPortParam)
}

type multiUDPMuxFromPortParam struct {
	ifFilter        func(string) (keep bool)
	ipFilter        func(ip net.IP) (keep bool)
	networks        []NetworkType
	readBufferSize  int
	writeBufferSize int
	logger          logging.LeveledLogger
	includeLoopback bool
	net             transport.Net
}

type udpMuxFromPortOption struct {
	f func(*multiUDPMuxFromPortParam)
}

func (o *udpMuxFromPortOption) apply(p *multiUDPMuxFromPortParam) {
	_ = "STUB: not implemented"

	// UDPMuxFromPortWithInterfaceFilter set the filter to filter out interfaces that should not be used.
	return
}

func UDPMuxFromPortWithInterfaceFilter(f func(string) (keep bool)) UDPMuxFromPortOption {
	_ = "STUB: not implemented"
	return *new(UDPMuxFromPortOption)
}

// UDPMuxFromPortWithIPFilter set the filter to filter out IP addresses that should not be used.
func UDPMuxFromPortWithIPFilter(f func(ip net.IP) (keep bool)) UDPMuxFromPortOption {
	_ = "STUB: not implemented"
	return *new(UDPMuxFromPortOption)
}

// UDPMuxFromPortWithNetworks set the networks that should be used. default is both IPv4 and IPv6.
func UDPMuxFromPortWithNetworks(networks ...NetworkType) UDPMuxFromPortOption {
	_ = "STUB: not implemented"
	return *new(UDPMuxFromPortOption)
}

// UDPMuxFromPortWithReadBufferSize set the UDP connection read buffer size.
func UDPMuxFromPortWithReadBufferSize(size int) UDPMuxFromPortOption {
	_ = "STUB: not implemented"
	return *new(UDPMuxFromPortOption)
}

// UDPMuxFromPortWithWriteBufferSize set the UDP connection write buffer size.
func UDPMuxFromPortWithWriteBufferSize(size int) UDPMuxFromPortOption {
	_ = "STUB: not implemented"
	return *new(UDPMuxFromPortOption)
}

// UDPMuxFromPortWithLogger set the logger for the created UDPMux.
func UDPMuxFromPortWithLogger(logger logging.LeveledLogger) UDPMuxFromPortOption {
	_ = "STUB: not implemented"
	return *new(UDPMuxFromPortOption)
}

// UDPMuxFromPortWithLoopback set loopback interface should be included.
func UDPMuxFromPortWithLoopback() UDPMuxFromPortOption {
	_ = "STUB: not implemented"
	return *new(UDPMuxFromPortOption)
}

// UDPMuxFromPortWithNet sets the network transport to use.
func UDPMuxFromPortWithNet(n transport.Net) UDPMuxFromPortOption {
	_ = "STUB: not implemented"
	return *new(UDPMuxFromPortOption)
}
