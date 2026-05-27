// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"io"
	"net"
	"sync"
	"time"

	"github.com/pion/logging"
	"github.com/pion/transport/v4/packetio"
)

type bufferedConn struct {
	net.Conn
	buf    *packetio.Buffer
	logger logging.LeveledLogger
	closed int32
}

func newBufferedConn(conn net.Conn, bufSize int, logger logging.LeveledLogger) net.Conn {
	_ = "STUB: not implemented"
	return *new(net.Conn)
}

func (bc *bufferedConn) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (bc *bufferedConn) writeProcess() { _ = "STUB: not implemented"; return }

func (bc *bufferedConn) Close() error { _ = "STUB: not implemented"; return nil }

type tcpPacketConn struct {
	params *tcpPacketParams

	// conns is a map of net.Conns indexed by remote net.Addr.String()
	conns map[string]net.Conn

	recvChan chan streamingPacket

	mu         sync.Mutex
	wg         sync.WaitGroup
	closedChan chan struct{}
	closeOnce  sync.Once
	aliveTimer *time.Timer
}

type streamingPacket struct {
	Data  []byte
	RAddr net.Addr
	Err   error
}

type tcpPacketParams struct {
	ReadBuffer    int
	LocalAddr     net.Addr
	Logger        logging.LeveledLogger
	WriteBuffer   int
	AliveDuration time.Duration
}

func newTCPPacketConn(params tcpPacketParams) *tcpPacketConn { _ = "STUB: not implemented"; return nil }

func (t *tcpPacketConn) ClearAliveTimer() { _ = "STUB: not implemented"; return }

func (t *tcpPacketConn) AddConn(conn net.Conn, firstPacketData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: recvChan can fill up and never drain in edge
// cases while closing a connection, which can cause the
// packetConn to never finish closing. Bail out early
// here to prevent that.

func (t *tcpPacketConn) startReading(conn net.Conn) { _ = "STUB: not implemented"; return }

// Only propagate connection closure errors if no other open connection exists.

func (t *tcpPacketConn) handleRecv(pkt streamingPacket) { _ = "STUB: not implemented"; return }

func (t *tcpPacketConn) isClosed() bool { _ = "STUB: not implemented"; return false }

// WriteTo is for passive and s-o candidates.
func (t *tcpPacketConn) ReadFrom(b []byte) (n int, rAddr net.Addr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

// WriteTo is for active and s-o candidates.
func (t *tcpPacketConn) WriteTo(buf []byte, rAddr net.Addr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (t *tcpPacketConn) closeAndLogError(closer io.Closer) { _ = "STUB: not implemented"; return }

func (t *tcpPacketConn) removeConn(conn net.Conn) bool { _ = "STUB: not implemented"; return false }

// wait for some time to flush pending writes

// read deadline as well just in case

func (t *tcpPacketConn) Close() error { _ = "STUB: not implemented"; return nil }

// wait for some time to flush pending writes

// read deadline as well just in case

func (t *tcpPacketConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (t *tcpPacketConn) SetDeadline(d time.Time) error { _ = "STUB: not implemented"; return nil }

func (t *tcpPacketConn) SetReadDeadline(d time.Time) error { _ = "STUB: not implemented"; return nil }

func (t *tcpPacketConn) SetWriteDeadline(d time.Time) error { _ = "STUB: not implemented"; return nil }

func (t *tcpPacketConn) CloseChannel() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (t *tcpPacketConn) String() string { _ = "STUB: not implemented"; return "" }
