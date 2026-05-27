// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build !js

package fakenet

import (
	"net"
	"time"
)

// MockPacketConn for tests.
type MockPacketConn struct{}

func (m *MockPacketConn) ReadFrom([]byte) (n int, addr net.Addr, err error) {
	_ = "STUB: not implemented" //nolint:revive
	return 0, *new(net.Addr), nil
}
func (m *MockPacketConn) WriteTo([]byte, net.Addr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, //nolint:revive
		nil
}
func (m *MockPacketConn) Close() error                     { _ = "STUB: not implemented"; return nil }            //nolint:revive
func (m *MockPacketConn) LocalAddr() net.Addr              { _ = "STUB: not implemented"; return *new(net.Addr) } //nolint:revive
func (m *MockPacketConn) SetDeadline(time.Time) error      { _ = "STUB: not implemented"; return nil }            //nolint:revive
func (m *MockPacketConn) SetReadDeadline(time.Time) error  { _ = "STUB: not implemented"; return nil }            //nolint:revive
func (m *MockPacketConn) SetWriteDeadline(time.Time) error { _ = "STUB: not implemented"; return nil }            //nolint:revive
