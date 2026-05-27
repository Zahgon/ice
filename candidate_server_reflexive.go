// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

// CandidateServerReflexive ...
type CandidateServerReflexive struct {
	candidateBase
}

// CandidateServerReflexiveConfig is the config required to create a new CandidateServerReflexive.
type CandidateServerReflexiveConfig struct {
	CandidateID string
	Network     string
	Address     string
	Port        int
	Component   uint16
	Priority    uint32
	Foundation  string
	RelAddr     string
	RelPort     int
}

// NewCandidateServerReflexive creates a new server reflective candidate.
func NewCandidateServerReflexive(config *CandidateServerReflexiveConfig) (*CandidateServerReflexive, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
