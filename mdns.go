// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"net"

	"github.com/pion/logging"
	"github.com/pion/mdns/v2"
	"github.com/pion/transport/v4"
)

// MulticastDNSMode represents the different Multicast modes ICE can run in.
type MulticastDNSMode byte

// MulticastDNSMode enum.
const (
	// MulticastDNSModeDisabled means remote mDNS candidates will be discarded, and local host candidates will use IPs.
	MulticastDNSModeDisabled MulticastDNSMode = iota + 1

	// MulticastDNSModeQueryOnly means remote mDNS candidates will be accepted, and local host candidates will use IPs.
	MulticastDNSModeQueryOnly

	// MulticastDNSModeQueryAndGather means remote mDNS candidates will be accepted,
	// and local host candidates will use mDNS.
	MulticastDNSModeQueryAndGather
)

func generateMulticastDNSName() (string, error) {
	_ = "STUB: not implemented"
	// https://tools.ietf.org/id/draft-ietf-rtcweb-mdns-ice-candidates-02.html#gathering
	// The unique name MUST consist of a version 4 UUID as defined in [RFC4122], followed by “.local”.
	return "", nil
}

//nolint:cyclop
func createMulticastDNS(
	netTransport transport.Net,
	networkTypes []NetworkType,
	interfaces []*transport.Interface,
	includeLoopback bool,
	localAddress net.IP,
	mDNSMode MulticastDNSMode,
	mDNSName string,
	log logging.LeveledLogger,
	loggerFactory logging.LoggerFactory,
) (*mdns.Conn, MulticastDNSMode, error) {
	_ = "STUB: not implemented"
	return nil, *new(MulticastDNSMode), nil
}

// If ICE fails to start MulticastDNS server just warn the user and continue

// If ICE fails to start MulticastDNS server just warn the user and continue

//nolint:nilerr
