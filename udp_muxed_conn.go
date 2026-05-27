// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"net"
	"sync"
	"time"

	"github.com/pion/logging"
)

type udpMuxedConnState int

const (
	udpMuxedConnOpen udpMuxedConnState = iota
	udpMuxedConnWaiting
	udpMuxedConnClosed
)

type udpMuxedConnParams struct {
	Mux       *UDPMuxDefault
	AddrPool  *sync.Pool
	Key       string
	LocalAddr net.Addr
	Logger    logging.LeveledLogger
}

// udpMuxedConn represents a logical packet conn for a single remote as identified by ufrag.
type udpMuxedConn struct {
	params *udpMuxedConnParams
	// Remote addresses that we have sent to on this conn
	addresses []ipPort

	// FIFO queue holding incoming packets
	bufHead, bufTail *bufferHolder
	notify           chan struct{}
	closedChan       chan struct{}
	state            udpMuxedConnState
	mu               sync.Mutex
}

func newUDPMuxedConn(params *udpMuxedConnParams) *udpMuxedConn {
	_ = "STUB: not implemented"
	return nil
}

func (c *udpMuxedConn) ReadFrom(b []byte) (n int, rAddr net.Addr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func (c *udpMuxedConn) WriteTo(buf []byte, rAddr net.Addr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Each time we write to a new address, we'll register it with the mux

func (c *udpMuxedConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *udpMuxedConn) SetDeadline(time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *udpMuxedConn) SetReadDeadline(time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *udpMuxedConn) SetWriteDeadline(time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *udpMuxedConn) CloseChannel() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *udpMuxedConn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *udpMuxedConn) isClosed() bool { _ = "STUB: not implemented"; return false }

func (c *udpMuxedConn) getAddresses() []ipPort { _ = "STUB: not implemented"; return nil }

func (c *udpMuxedConn) addAddress(addr ipPort) { _ = "STUB: not implemented"; return }

// Map it on mux

func (c *udpMuxedConn) removeAddress(addr ipPort) { _ = "STUB: not implemented"; return }

func (c *udpMuxedConn) containsAddress(addr ipPort) bool { _ = "STUB: not implemented"; return false }

func (c *udpMuxedConn) writePacket(data []byte, addr *net.UDPAddr) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert
