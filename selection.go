// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"net"
	"time"

	"github.com/pion/logging"
	"github.com/pion/stun/v3"
)

type pairCandidateSelector interface {
	Start()
	ContactCandidates()
	PingCandidate(local, remote Candidate)
	HandleSuccessResponse(m *stun.Message, local, remote Candidate, remoteAddr net.Addr)
	HandleBindingRequest(m *stun.Message, local, remote Candidate)
}

type controllingSelector struct {
	startTime     time.Time
	agent         *Agent
	nominatedPair *CandidatePair
	log           logging.LeveledLogger
}

func (s *controllingSelector) Start() { _ = "STUB: not implemented"; return }

func (s *controllingSelector) isNominatable(c Candidate) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *controllingSelector) ContactCandidates() { _ = "STUB: not implemented"; return }

// If automatic renomination is enabled, continuously ping all candidate pairs
// to keep them tested with fresh RTT measurements for switching decisions

func (s *controllingSelector) nominatePair(pair *CandidatePair) {
	_ = "STUB: not implemented"
	// The controlling agent MUST include the USE-CANDIDATE attribute in
	// order to nominate a candidate pair (Section 8.1.1).  The controlled
	// agent MUST NOT include the USE-CANDIDATE attribute in a Binding
	// request.
	return
}

func (s *controllingSelector) HandleBindingRequest(message *stun.Message, local, remote Candidate) {
	_ = "STUB: not implemented" //nolint:cyclop
	return
}

func (s *controllingSelector) HandleSuccessResponse(m *stun.Message, local, remote Candidate, remoteAddr net.Addr) {
	_ = "STUB: not implemented"
	return
}

// Assert that NAT is not symmetric
// https://tools.ietf.org/html/rfc8445#section-7.2.5.2.1

// This shouldn't happen

// Handle nomination/renomination

// If this is a renomination request (has nomination value), always update the selected pair
// If it's a standard nomination (no value), only set if no pair is selected yet

func (s *controllingSelector) PingCandidate(local, remote Candidate) {
	_ = "STUB: not implemented"
	return
}

// checkForAutomaticRenomination evaluates if automatic renomination should occur.
// This is called periodically when the agent is in connected state and automatic
// renomination is enabled.
func (s *controllingSelector) checkForAutomaticRenomination() { _ = "STUB: not implemented"; return }

// Update last renomination time to prevent rapid renominations

type controlledSelector struct {
	agent          *Agent
	log            logging.LeveledLogger
	lastNomination *uint32 // For renomination: tracks highest nomination value seen
}

func (s *controlledSelector) Start() { _ = "STUB: not implemented"; return }

// shouldAcceptNomination checks if a nomination should be accepted based on renomination rules.
func (s *controlledSelector) shouldAcceptNomination(nominationValue *uint32) bool {
	_ = "STUB: not implemented"
	// If no nomination value, accept normally (standard ICE nomination)
	return false
}

// If nomination value is present, controlling side is using renomination
// Apply "last nomination wins" rule

// shouldSwitchSelectedPair determines if we should switch to a new nominated pair.
// Returns true if the switch should occur, false otherwise.
func (s *controlledSelector) shouldSwitchSelectedPair(pair, selectedPair *CandidatePair, nominationValue *uint32) bool {
	_ = "STUB: not implemented"
	return false
}

// No current selection, accept the nomination

// Same pair, no change needed

// Renomination is in use (nomination value present)
// Accept the switch based on nomination value alone, not priority
// The shouldAcceptNomination check already validated this is a valid renomination

// Standard ICE nomination without renomination - apply priority rules
// Only switch if we don't check priority, OR new pair has strictly higher priority

func (s *controlledSelector) ContactCandidates() { _ = "STUB: not implemented"; return }

func (s *controlledSelector) PingCandidate(local, remote Candidate) {
	_ = "STUB: not implemented"
	return
}

func (s *controlledSelector) HandleSuccessResponse(m *stun.Message, local, remote Candidate, remoteAddr net.Addr) {
	_ = "STUB: not implemented"
	//nolint:godox
	// TODO according to the standard we should specifically answer a failed nomination:
	// https://tools.ietf.org/html/rfc8445#section-7.3.1.5
	// If the controlled agent does not accept the request from the
	// controlling agent, the controlled agent MUST reject the nomination
	// request with an appropriate error code response (e.g., 400)
	// [RFC5389].
	return
}

// Assert that NAT is not symmetric
// https://tools.ietf.org/html/rfc8445#section-7.2.5.2.1

// This shouldn't happen

func (s *controlledSelector) HandleBindingRequest(message *stun.Message, local, remote Candidate) {
	_ = "STUB: not implemented" //nolint:cyclop
	return
}

//nolint:nestif
// https://tools.ietf.org/html/rfc8445#section-7.3.1.5

// Check for renomination attribute

// Check if we should accept this nomination based on renomination rules

// If the state of this pair is Succeeded, it means that the check
// previously sent by this pair produced a successful response and
// generated a valid pair (Section 7.2.5.3.2).  The agent sets the
// nominated flag value of the valid pair to true.

// If the received Binding request triggered a new check to be
// enqueued in the triggered-check queue (Section 7.3.1.4), once the
// check is sent and if it generates a successful response, and
// generates a valid pair, the agent sets the nominated flag of the
// pair to true.  If the request fails (Section 7.2.5.2), the agent
// MUST remove the candidate pair from the valid list, set the
// candidate pair state to Failed, and set the checklist state to
// Failed.

// Only send a triggered check during ICE checking phase (RFC 8445 §7.3.1.4).
// Once the pair is established (succeeded + selected), sending a triggered check
// on every inbound request creates a ping-pong busy loop: the remote side responds
// and sends its own request, which triggers another check here, repeating at 1/RTT.
// After connection, consent freshness is maintained by checkKeepalive() on a timer.

type liteSelector struct {
	pairCandidateSelector
}

// A lite selector should not contact candidates.
func (s *liteSelector) ContactCandidates() { _ = "STUB: not implemented"; return }

//nolint:godox
// https://github.com/pion/ice/issues/96
// TODO: implement lite controlling agent. For now falling back to full agent.
// This only happens if both peers are lite. See RFC 8445 S6.1.1 and S6.2
