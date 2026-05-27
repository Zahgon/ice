// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"context"
	"net"
	"net/netip"
	"sync/atomic"
	"time"

	"github.com/pion/logging"
	"github.com/pion/transport/v4/packetio"
)

type activeTCPConn struct {
	readBuffer, writeBuffer *packetio.Buffer
	localAddr, remoteAddr   atomic.Value
	conn                    atomic.Value // stores net.Conn
	closed                  atomic.Bool
}

func newActiveTCPConn(
	ctx context.Context,
	localAddress string,
	remoteAddress netip.AddrPort,
	log logging.LeveledLogger,
) (a *activeTCPConn) {
	_ = "STUB: not implemented"
	return nil
}

func (a *activeTCPConn) ReadFrom(buff []byte) (n int, srcAddr net.Addr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

// RemoteAddr is assuredly set *after* we can read from the buffer

func (a *activeTCPConn) WriteTo(buff []byte, _ net.Addr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (a *activeTCPConn) Close() error { _ = "STUB: not implemented"; return nil }

func (a *activeTCPConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// RemoteAddr returns the remote address of the connection which is only
// set once a background goroutine has successfully dialed. That means
// this may return ":0" for the address prior to that happening. If this
// becomes an issue, we can introduce a synchronization point between Dial
// and these methods.
func (a *activeTCPConn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (a *activeTCPConn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (a *activeTCPConn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (a *activeTCPConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func getTCPAddrOnInterface(address string) (*net.TCPAddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
