// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

// Role represents ICE agent role, which can be controlling or controlled.
type Role byte

// Possible ICE agent roles.
const (
	Controlling Role = iota
	Controlled
)

// UnmarshalText implements TextUnmarshaler.
func (r *Role) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText implements TextMarshaler.
func (r Role) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (r Role) String() string { _ = "STUB: not implemented"; return "" }
