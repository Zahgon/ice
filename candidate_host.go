// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"net/netip"
)

// CandidateHost is a candidate of type host.
type CandidateHost struct {
	candidateBase

	network string
}

// CandidateHostConfig is the config required to create a new CandidateHost.
type CandidateHostConfig struct {
	CandidateID       string
	Network           string
	Address           string
	Port              int
	Component         uint16
	Priority          uint32
	Foundation        string
	TCPType           TCPType
	IsLocationTracked bool
}

// NewCandidateHost creates a new host candidate.
func NewCandidateHost(config *CandidateHostConfig) (*CandidateHost, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Until mDNS candidate is resolved assume it is UDPv4

func (c *CandidateHost) setIPAddr(addr netip.Addr) error { _ = "STUB: not implemented"; return nil }
