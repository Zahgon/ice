// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"net/netip"
)

const (
	udp  = "udp"
	tcp  = "tcp"
	udp4 = "udp4"
	udp6 = "udp6"
	tcp4 = "tcp4"
	tcp6 = "tcp6"
)

func supportedNetworkTypes() []NetworkType { _ = "STUB: not implemented"; return nil }

// NetworkType represents the type of network.
type NetworkType int

const (
	// NetworkTypeUDP4 indicates UDP over IPv4.
	NetworkTypeUDP4 NetworkType = iota + 1

	// NetworkTypeUDP6 indicates UDP over IPv6.
	NetworkTypeUDP6

	// NetworkTypeTCP4 indicates TCP over IPv4.
	NetworkTypeTCP4

	// NetworkTypeTCP6 indicates TCP over IPv6.
	NetworkTypeTCP6
)

func (t NetworkType) String() string { _ = "STUB: not implemented"; return "" }

// IsUDP returns true when network is UDP4 or UDP6.
func (t NetworkType) IsUDP() bool { _ = "STUB: not implemented"; return false }

// IsTCP returns true when network is TCP4 or TCP6.
func (t NetworkType) IsTCP() bool { _ = "STUB: not implemented"; return false }

// NetworkShort returns the short network description.
func (t NetworkType) NetworkShort() string { _ = "STUB: not implemented"; return "" }

// IsReliable returns true if the network is reliable.
func (t NetworkType) IsReliable() bool { _ = "STUB: not implemented"; return false }

// IsIPv4 returns whether the network type is IPv4 or not.
func (t NetworkType) IsIPv4() bool { _ = "STUB: not implemented"; return false }

// IsIPv6 returns whether the network type is IPv6 or not.
func (t NetworkType) IsIPv6() bool { _ = "STUB: not implemented"; return false }

// determineNetworkType determines the type of network based on
// the short network string and an IP address.
func determineNetworkType(network string, ip netip.Addr) (NetworkType, error) {
	_ = "STUB: not implemented"
	// we'd rather have an IPv4-mapped IPv6 become IPv4 so that it is usable.
	return *new(NetworkType), nil
}
