// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import "github.com/pion/stun/v3"

// UseCandidateAttr represents USE-CANDIDATE attribute.
type UseCandidateAttr struct{}

// AddTo adds USE-CANDIDATE attribute to message.
func (UseCandidateAttr) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// IsSet returns true if USE-CANDIDATE attribute is set.
func (UseCandidateAttr) IsSet(m *stun.Message) bool { _ = "STUB: not implemented"; return false }

// UseCandidate is shorthand for UseCandidateAttr.
func UseCandidate() UseCandidateAttr { _ = "STUB: not implemented"; return *new(UseCandidateAttr) }
