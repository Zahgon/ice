// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

// CandidateRelatedAddress convey transport addresses related to the
// candidate, useful for diagnostics and other purposes.
type CandidateRelatedAddress struct {
	Address string
	Port    int
}

// String makes CandidateRelatedAddress printable.
func (c *CandidateRelatedAddress) String() string { _ = "STUB: not implemented"; return "" }

// Equal allows comparing two CandidateRelatedAddresses.
// The CandidateRelatedAddress are allowed to be nil.
func (c *CandidateRelatedAddress) Equal(other *CandidateRelatedAddress) bool {
	_ = "STUB: not implemented"
	return false
}
