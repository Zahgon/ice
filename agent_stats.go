// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

// GetCandidatePairsStats returns a list of candidate pair stats.
func (a *Agent) GetCandidatePairsStats() []CandidatePairStats {
	_ = "STUB: not implemented"
	return nil
}

// AvailableOutgoingBitrate float64
// AvailableIncomingBitrate float64
// CircuitBreakerTriggerCount uint32

// RetransmissionsReceived uint64
// RetransmissionsSent uint64
// ConsentRequestsSent uint64
// ConsentExpiredTimestamp time.Time

// GetSelectedCandidatePairStats returns a candidate pair stats for selected candidate pair.
// Returns false if there is no selected pair.
func (a *Agent) GetSelectedCandidatePairStats() (CandidatePairStats, bool) {
	_ = "STUB: not implemented"
	return *new(CandidatePairStats), false
}

// FirstRequestTimestamp time.Time
// LastRequestTimestamp time.Time
// LastResponseTimestamp time.Time

// AvailableOutgoingBitrate float64
// AvailableIncomingBitrate float64
// CircuitBreakerTriggerCount uint32
// RequestsReceived uint64
// RequestsSent uint64

// ResponsesSent uint64
// RetransmissionsReceived uint64
// RetransmissionsSent uint64
// ConsentRequestsSent uint64
// ConsentExpiredTimestamp time.Time

// GetLocalCandidatesStats returns a list of local candidates stats.
func (a *Agent) GetLocalCandidatesStats() []CandidateStats { _ = "STUB: not implemented"; return nil }

// GetRemoteCandidatesStats returns a list of remote candidates stats.
func (a *Agent) GetRemoteCandidatesStats() []CandidateStats { _ = "STUB: not implemented"; return nil }

// getCandidatesStats returns a list of candidates stats.
func (a *Agent) getCandidatesStats(isLocal bool) []CandidateStats {
	_ = "STUB: not implemented"
	return nil
}

// URL string
