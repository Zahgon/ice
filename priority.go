// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"github.com/pion/stun/v3"
)

// PriorityAttr represents PRIORITY attribute.
type PriorityAttr uint32

const prioritySize = 4 // 32 bit

// AddTo adds PRIORITY attribute to message.
func (p PriorityAttr) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes PRIORITY attribute from message.
func (p *PriorityAttr) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }
