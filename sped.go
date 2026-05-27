// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"github.com/pion/stun/v3"
)

// DtlsInStunAttribute is a STUN attribute for carrying DTLS embedded in STUN.
type DtlsInStunAttribute []byte

// AddTo adds DTLS-in-STUN attribute to message.
func (d DtlsInStunAttribute) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes DTLS-in-STUN attribute from message.
func (d *DtlsInStunAttribute) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// DtlsInStunAckAttribute is a STUN attribute for acknowledging the receipt
// of DTLS packets (embedded in STUN or without embedding).
type DtlsInStunAckAttribute []uint32

// ACKs are 32-bit values, and the attribute can carry up to four of them.
const (
	ackSizeValues = 4
	ackSizeBytes  = ackSizeValues * 4
)

// AddTo adds DTLS-in-STUN-ACK attribute to message.
func (a DtlsInStunAckAttribute) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes DTLS-in-STUN-ACK attribute from message.
func (a *DtlsInStunAckAttribute) GetFrom(m *stun.Message) error {
	_ = "STUB: not implemented"
	return nil
}
