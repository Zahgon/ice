// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

const (
	// These preference values come from libwebrtc
	//nolint:lll
	// https://source.chromium.org/chromium/chromium/src/+/main:third_party/webrtc/p2p/base/p2p_constants.h;l=126;drc=bf712ec1a13783224debb691ba88ad5c15b93194
	preferenceRelayTLS  = 0
	preferenceRelayTCP  = 1
	preferenceRelayDTLS = 2
	preferenceRelayUDP  = 3
)

// CandidateRelay ...
type CandidateRelay struct {
	candidateBase

	relayProtocol string
	onClose       func() error
}

// CandidateRelayConfig is the config required to create a new CandidateRelay.
type CandidateRelayConfig struct {
	CandidateID   string
	Network       string
	Address       string
	Port          int
	Component     uint16
	Priority      uint32
	Foundation    string
	RelAddr       string
	RelPort       int
	RelayProtocol string
	OnClose       func() error
}

// NewCandidateRelay creates a new relay candidate.
func NewCandidateRelay(config *CandidateRelayConfig) (*CandidateRelay, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RelayProtocol returns the protocol used between the endpoint and the relay server.
func (c *CandidateRelay) RelayProtocol() string { _ = "STUB: not implemented"; return "" }

func (c *CandidateRelay) close() error { _ = "STUB: not implemented"; return nil }

func (c *CandidateRelay) copy() (Candidate, error) {
	_ = "STUB: not implemented"
	return *new(Candidate), nil
}

// relayProtocolPreference returns the preference for the relay protocol.
func relayProtocolPreference(relayProtocol string) uint16 { _ = "STUB: not implemented"; return 0 }
