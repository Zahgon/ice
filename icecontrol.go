// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"github.com/pion/stun/v3"
)

// tiebreaker is common helper for ICE-{CONTROLLED,CONTROLLING}
// and represents the so-called tiebreaker number.
type tiebreaker uint64

const tiebreakerSize = 8 // 64 bit

// AddToAs adds tiebreaker value to m as t attribute.
func (a tiebreaker) AddToAs(m *stun.Message, t stun.AttrType) error {
	_ = "STUB: not implemented"
	return nil
}

// GetFromAs decodes tiebreaker value in message getting it as for t type.
func (a *tiebreaker) GetFromAs(m *stun.Message, t stun.AttrType) error {
	_ = "STUB: not implemented"
	return nil
}

// AttrControlled represents ICE-CONTROLLED attribute.
type AttrControlled uint64

// AddTo adds ICE-CONTROLLED to message.
func (c AttrControlled) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes ICE-CONTROLLED from message.
func (c *AttrControlled) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// AttrControlling represents ICE-CONTROLLING attribute.
type AttrControlling uint64

// AddTo adds ICE-CONTROLLING to message.
func (c AttrControlling) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes ICE-CONTROLLING from message.
func (c *AttrControlling) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// AttrControl is helper that wraps ICE-{CONTROLLED,CONTROLLING}.
type AttrControl struct {
	Role       Role
	Tiebreaker uint64
}

// AddTo adds ICE-CONTROLLED or ICE-CONTROLLING attribute depending on Role.
func (c AttrControl) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes Role and Tiebreaker value from message.
func (c *AttrControl) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }
