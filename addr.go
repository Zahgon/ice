// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"net"
	"net/netip"
)

func addrWithOptionalZone(addr netip.Addr, zone string) netip.Addr {
	_ = "STUB: not implemented"
	return *new(netip.Addr)
}

// parseAddrFromIface should only be used when it's known the address belongs to that interface.
// e.g. it's LocalAddress on a listener.
func parseAddrFromIface(in net.Addr, ifcName string) (netip.Addr, int, NetworkType, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), 0, *new(NetworkType), nil
}

// net.IPNet does not have a Zone but we provide it from the interface

func parseAddr(in net.Addr) (netip.Addr, int, NetworkType, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), 0, *new(NetworkType), nil
}

type addrParseError struct {
	addr net.Addr
}

func (e addrParseError) Error() string { _ = "STUB: not implemented"; return "" }

type ipConvertError struct {
	ip []byte
}

func (e ipConvertError) Error() string { _ = "STUB: not implemented"; return "" }

func ipAddrToNetIP(ip []byte, zone string) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// we'd rather have an IPv4-mapped IPv6 become IPv4 so that it is usable.

func createAddr(network NetworkType, ip netip.Addr, port int) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func addrEqual(a, b net.Addr) bool { _ = "STUB: not implemented"; return false }

// AddrPort is  an IP and a port number.
type AddrPort [18]byte

func toAddrPort(addr net.Addr) AddrPort { _ = "STUB: not implemented"; return *new(AddrPort) }

//nolint:gosec // G115  false positive
//nolint:gosec // G115  false positive

//nolint:gosec // G115 false positive
//nolint:gosec // G115 false positive
