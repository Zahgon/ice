// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package ice ...
//
//nolint:dupl
package ice

// CandidatePeerReflexive ...
type CandidatePeerReflexive struct {
	candidateBase
}

// CandidatePeerReflexiveConfig is the config required to create a new CandidatePeerReflexive.
type CandidatePeerReflexiveConfig struct {
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

// NewCandidatePeerReflexive creates a new peer reflective candidate.
func NewCandidatePeerReflexive(config *CandidatePeerReflexiveConfig) (*CandidatePeerReflexive, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
