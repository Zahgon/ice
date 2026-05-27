// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"github.com/pion/stun/v3"
)

// Default STUN Nomination attribute type for ICE renomination.
// Following the specification draft-thatcher-ice-renomination-01.
const (
	// DefaultNominationAttribute represents the default STUN Nomination attribute.
	// This is a custom attribute for ICE renomination support.
	// This value can be overridden via AgentConfig.NominationAttribute.
	DefaultNominationAttribute stun.AttrType = 0xC001 // matching libwebrtc.
)

// NominationAttribute represents a STUN Nomination attribute.
type NominationAttribute struct {
	Value uint32
}

// GetFrom decodes a Nomination attribute from a STUN message.
func (a *NominationAttribute) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFromWithType decodes a Nomination attribute from a STUN message using a specific attribute type.
func (a *NominationAttribute) GetFromWithType(m *stun.Message, attrType stun.AttrType) error {
	_ = "STUB: not implemented"
	return nil
}

// Extract 24-bit value from the last 3 bytes

// AddTo adds a Nomination attribute to a STUN message.
func (a NominationAttribute) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// AddToWithType adds a Nomination attribute to a STUN message using a specific attribute type.
func (a NominationAttribute) AddToWithType(m *stun.Message, attrType stun.AttrType) error {
	_ = "STUB: not implemented"
	// Store as 4 bytes with first byte as 0
	return nil
}

//nolint:gosec
//nolint:gosec
//nolint:gosec

// String returns string representation of the nomination attribute.
func (a NominationAttribute) String() string { _ = "STUB: not implemented"; return "" }

// Nomination creates a new STUN nomination attribute.
func Nomination(value uint32) NominationAttribute {
	_ = "STUB: not implemented"
	return *new(NominationAttribute)
}

// NominationSetter is a STUN setter for nomination attribute with configurable type.
type NominationSetter struct {
	Value    uint32
	AttrType stun.AttrType
}

// AddTo adds a Nomination attribute to a STUN message using the configured attribute type.
func (n NominationSetter) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }
