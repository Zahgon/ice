// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import "sync"

// OnConnectionStateChange sets a handler that is fired when the connection state changes.
func (a *Agent) OnConnectionStateChange(f func(ConnectionState)) error {
	_ = "STUB: not implemented"
	return nil
}

// OnSelectedCandidatePairChange sets a handler that is fired when the final candidate.
// pair is selected.
func (a *Agent) OnSelectedCandidatePairChange(f func(Candidate, Candidate)) error {
	_ = "STUB: not implemented"
	return nil
}

// OnCandidate sets a handler that is fired when new candidates gathered. When
// the gathering process complete the last candidate is nil.
func (a *Agent) OnCandidate(f func(Candidate)) error { _ = "STUB: not implemented"; return nil }

func (a *Agent) onSelectedCandidatePairChange(p *CandidatePair) { _ = "STUB: not implemented"; return }

func (a *Agent) onCandidate(c Candidate) { _ = "STUB: not implemented"; return }

func (a *Agent) onConnectionStateChange(s ConnectionState) { _ = "STUB: not implemented"; return }

type handlerNotifier struct {
	sync.Mutex
	runningConnectionStates bool
	runningCandidates       bool
	runningCandidatePairs   bool
	notifiers               sync.WaitGroup

	connectionStates    []ConnectionState
	connectionStateFunc func(ConnectionState)

	candidates    []Candidate
	candidateFunc func(Candidate)

	selectedCandidatePairs []*CandidatePair
	candidatePairFunc      func(*CandidatePair)

	// State for closing
	done chan struct{}
}

func (h *handlerNotifier) Close(graceful bool) {
	_ = "STUB: not implemented"

	// if we were closed ungracefully before, we now
	// want ot wait.
	return
}

func (h *handlerNotifier) EnqueueConnectionState(state ConnectionState) {
	_ = "STUB: not implemented"
	return
}

func (h *handlerNotifier) EnqueueCandidate(cand Candidate) { _ = "STUB: not implemented"; return }

func (h *handlerNotifier) EnqueueSelectedCandidatePair(pair *CandidatePair) {
	_ = "STUB: not implemented"
	return
}
