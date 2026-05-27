// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"sync/atomic"
	"time"

	"github.com/pion/stun/v3"
)

func newCandidatePair(local, remote Candidate, controlling bool) *CandidatePair {
	_ = "STUB: not implemented"
	return nil
}

// CandidatePair is a combination of a local and remote candidate.
type CandidatePair struct {
	id                       uint64
	iceRoleControlling       bool
	Remote                   Candidate
	Local                    Candidate
	priorityOverride         uint64
	hasPriorityOverride      bool
	bindingRequestCount      uint16
	state                    CandidatePairState
	nominated                bool
	nominateOnBindingSuccess bool

	// stats
	currentRoundTripTime int64 // in ns
	totalRoundTripTime   int64 // in ns

	packetsSent          uint32
	packetsReceived      uint32
	bytesSent            uint64
	bytesReceived        uint64
	lastPacketSentAt     atomic.Value // time.Time
	lastPacketReceivedAt atomic.Value // time.Time

	requestsReceived  uint64
	requestsSent      uint64
	responsesReceived uint64
	responsesSent     uint64

	firstRequestSentAt      atomic.Value // time.Time
	lastRequestSentAt       atomic.Value // time.Time
	firstResponseReceivedAt atomic.Value // time.Time
	lastResponseReceivedAt  atomic.Value // time.Time
	firstRequestReceivedAt  atomic.Value // time.Time
	lastRequestReceivedAt   atomic.Value // time.Time
}

func (p *CandidatePair) String() string { _ = "STUB: not implemented"; return "" }

func (p *CandidatePair) equal(other *CandidatePair) bool { _ = "STUB: not implemented"; return false }

func (p *CandidatePair) setPriorityOverride(prio uint64) { _ = "STUB: not implemented"; return }

// RFC 5245 - 5.7.2.  Computing Pair Priority and Ordering Pairs
// Let G be the priority for the candidate provided by the controlling
// agent.  Let D be the priority for the candidate provided by the
// controlled agent.
// pair priority = 2^32*MIN(G,D) + 2*MAX(G,D) + (G>D?1:0).
func (p *CandidatePair) priority() uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:varnamelen // clearer to use g and d here

// Just implement these here rather
// than fooling around with the math package

// 1<<32 overflows uint32; and if both g && d are
// maxUint32, this result would overflow uint64

func (p *CandidatePair) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (a *Agent) sendSTUN(msg *stun.Message, local, remote Candidate) {
	_ = "STUB: not implemented"
	return
}

// UpdateRoundTripTime sets the current round time of this pair and
// accumulates total round trip time and responses received.
func (p *CandidatePair) UpdateRoundTripTime(rtt time.Duration) { _ = "STUB: not implemented"; return }

// CurrentRoundTripTime returns the current round trip time in seconds
// https://www.w3.org/TR/webrtc-stats/#dom-rtcicecandidatepairstats-currentroundtriptime
func (p *CandidatePair) CurrentRoundTripTime() float64 { _ = "STUB: not implemented"; return 0 }

// TotalRoundTripTime returns the current round trip time in seconds
// https://www.w3.org/TR/webrtc-stats/#dom-rtcicecandidatepairstats-totalroundtriptime
func (p *CandidatePair) TotalRoundTripTime() float64 { _ = "STUB: not implemented"; return 0 }

// RequestsReceived returns the total number of connectivity checks received
// https://www.w3.org/TR/webrtc-stats/#dom-rtcicecandidatepairstats-requestsreceived
func (p *CandidatePair) RequestsReceived() uint64 { _ = "STUB: not implemented"; return 0 }

// RequestsSent returns the total number of connectivity checks sent
// https://www.w3.org/TR/webrtc-stats/#dom-rtcicecandidatepairstats-requestssent
func (p *CandidatePair) RequestsSent() uint64 { _ = "STUB: not implemented"; return 0 }

// ResponsesReceived returns the total number of connectivity responses received
// https://www.w3.org/TR/webrtc-stats/#dom-rtcicecandidatepairstats-responsesreceived
func (p *CandidatePair) ResponsesReceived() uint64 { _ = "STUB: not implemented"; return 0 }

// ResponsesSent returns the total number of connectivity responses sent
// https://www.w3.org/TR/webrtc-stats/#dom-rtcicecandidatepairstats-responsessent
func (p *CandidatePair) ResponsesSent() uint64 { _ = "STUB: not implemented"; return 0 }

// PacketsSent returns total application (non-STUN) packets sent on this pair.
func (p *CandidatePair) PacketsSent() uint32 { _ = "STUB: not implemented"; return 0 }

// PacketsReceived returns total application (non-STUN) packets received on this pair.
func (p *CandidatePair) PacketsReceived() uint32 { _ = "STUB: not implemented"; return 0 }

// BytesSent returns total application bytes sent on this pair.
func (p *CandidatePair) BytesSent() uint64 { _ = "STUB: not implemented"; return 0 }

// BytesReceived returns total application bytes received on this pair.
func (p *CandidatePair) BytesReceived() uint64 { _ = "STUB: not implemented"; return 0 }

// LastPacketSentAt returns the timestamp of the last application packet sent.
func (p *CandidatePair) LastPacketSentAt() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// LastPacketReceivedAt returns the timestamp of the last application packet received.
func (p *CandidatePair) LastPacketReceivedAt() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// UpdatePacketSent increments packet/byte counters and updates timestamp for a sent application packet.
func (p *CandidatePair) UpdatePacketSent(n int) { _ = "STUB: not implemented"; return }

// #nosec G115 -- n > 0 validated above

// UpdatePacketReceived increments packet/byte counters and updates timestamp for a received application packet.
func (p *CandidatePair) UpdatePacketReceived(n int) { _ = "STUB: not implemented"; return }

// #nosec G115 -- n > 0 validated above

// FirstRequestSentAt returns the timestamp of the first connectivity check sent.
func (p *CandidatePair) FirstRequestSentAt() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// LastRequestSentAt returns the timestamp of the last connectivity check sent.
func (p *CandidatePair) LastRequestSentAt() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// Deprecated: use FirstResponseReceivedAt
// FirstReponseReceivedAt returns the timestamp of the first connectivity response received.
func (p *CandidatePair) FirstReponseReceivedAt() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// FirstResponseReceivedAt returns the timestamp of the first connectivity response received.
func (p *CandidatePair) FirstResponseReceivedAt() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// LastResponseReceivedAt returns the timestamp of the last connectivity response received.
func (p *CandidatePair) LastResponseReceivedAt() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// FirstRequestReceivedAt returns the timestamp of the first connectivity check received.
func (p *CandidatePair) FirstRequestReceivedAt() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// LastRequestReceivedAt returns the timestamp of the last connectivity check received.
func (p *CandidatePair) LastRequestReceivedAt() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// UpdateRequestSent increments the number of requests sent and updates the timestamp.
func (p *CandidatePair) UpdateRequestSent() { _ = "STUB: not implemented"; return }

// UpdateResponseSent increments the number of responses sent.
func (p *CandidatePair) UpdateResponseSent() { _ = "STUB: not implemented"; return }

// UpdateRequestReceived increments the number of requests received and updates the timestamp.
func (p *CandidatePair) UpdateRequestReceived() { _ = "STUB: not implemented"; return }

// ID returns the unique identifier for this candidate pair.
func (p *CandidatePair) ID() uint64 { _ = "STUB: not implemented"; return 0 }
