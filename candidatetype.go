// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

// CandidateType represents the type of candidate.
type CandidateType byte

// CandidateType enum.
const (
	CandidateTypeUnspecified CandidateType = iota
	CandidateTypeHost
	CandidateTypeServerReflexive
	CandidateTypePeerReflexive
	CandidateTypeRelay
)

// String makes CandidateType printable.
func (c CandidateType) String() string { _ = "STUB: not implemented"; return "" }

// Preference returns the preference weight of a CandidateType
//
// 4.1.2.2.  Guidelines for Choosing Type and Local Preferences
// The RECOMMENDED values are 126 for host candidates, 100
// for server reflexive candidates, 110 for peer reflexive candidates,
// and 0 for relayed candidates.
func (c CandidateType) Preference() uint16 { _ = "STUB: not implemented"; return 0 }

func containsCandidateType(candidateType CandidateType, candidateTypeList []CandidateType) bool {
	_ = "STUB: not implemented"
	return false
}
