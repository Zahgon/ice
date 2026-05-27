// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"net"
	"net/netip"

	"github.com/pion/logging"
	"github.com/pion/transport/v4"
)

type ifaceAddr struct {
	addr  netip.Addr
	iface string
}

// The conditions of invalidation written below are defined in
// https://tools.ietf.org/html/rfc8445#section-5.1.1.1
// It is partial because the link-local check is done later in various gather local
// candidate methods which conditionally accept IPv6 based on usage of mDNS or not.
func isSupportedIPv6Partial(ip net.IP) bool { _ = "STUB: not implemented"; return false }

// Deprecated IPv4-compatible IPv6 addresses [RFC4291] and IPv6 site-
//   local unicast addresses [RFC3879] MUST NOT be included in the
//   address candidates.
// !(IPv4-compatible IPv6)
// !(IPv6 site-local unicast)

func isZeros(ip net.IP) bool { _ = "STUB: not implemented"; return false }

//nolint:gocognit,cyclop
func localInterfaces(
	n transport.Net,
	interfaceFilter func(string) (keep bool),
	ipFilter func(net.IP) (keep bool),
	networkTypes []NetworkType,
	includeLoopback bool,
) ([]*transport.Interface, []ifaceAddr, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Interface down

// Loopback interface

//nolint:cyclop
func listenUDPInPortRange(
	netTransport transport.Net,
	log logging.LeveledLogger,
	portMax, portMin int,
	network string,
	lAddr *net.UDPAddr,
) (transport.UDPConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.UDPConn), nil
}

// Start at 1024 which is non-privileged

//nolint:nilerr
