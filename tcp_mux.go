// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"errors"
	"io"
	"net"
	"sync"
	"time"

	"github.com/pion/logging"
)

// ErrGetTransportAddress can't convert net.Addr to underlying type (UDPAddr or TCPAddr).
var ErrGetTransportAddress = errors.New("failed to get local transport address")

// TCPMux is allows grouping multiple TCP net.Conns and using them like UDP
// net.PacketConns. The main implementation of this is TCPMuxDefault, and this
// interface exists to allow mocking in tests.
type TCPMux interface {
	io.Closer
	GetConnByUfrag(ufrag string, isIPv6 bool, local net.IP) (net.PacketConn, error)
	RemoveConnByUfrag(ufrag string)
}

type ipAddr string

// TCPMuxDefault muxes TCP net.Conns into net.PacketConns and groups them by
// Ufrag. It is a default implementation of TCPMux interface.
type TCPMuxDefault struct {
	params *TCPMuxParams
	closed bool

	// connsIPv4 and connsIPv6 are maps of all tcpPacketConns indexed by ufrag and local address
	connsIPv4, connsIPv6 map[string]map[ipAddr]*tcpPacketConn

	mu sync.Mutex
	wg sync.WaitGroup
}

// TCPMuxParams are parameters for TCPMux.
type TCPMuxParams struct {
	Listener       net.Listener
	Logger         logging.LeveledLogger
	ReadBufferSize int

	// Maximum buffer size for write op. 0 means no write buffer, the write op will block until the whole packet is written
	// if the write buffer is full, the subsequent write packet will be dropped until it has enough space.
	// a default 4MB is recommended.
	WriteBufferSize int

	// A new established connection will be removed if the first STUN binding request is not received within this timeout,
	// avoiding the client with bad network or attacker to create a lot of empty connections.
	// Default 30s timeout will be used if not set.
	FirstStunBindTimeout time.Duration

	// TCPMux will create connection from STUN binding request with an unknown username, if
	// the connection is not used in the timeout, it will be removed to avoid resource leak / attack.
	// Default 30s timeout will be used if not set.
	AliveDurationForConnFromStun time.Duration
}

// NewTCPMuxDefault creates a new instance of TCPMuxDefault.
func NewTCPMuxDefault(params TCPMuxParams) *TCPMuxDefault { _ = "STUB: not implemented"; return nil }

func (m *TCPMuxDefault) start() { _ = "STUB: not implemented"; return }

// LocalAddr returns the listening address of this TCPMuxDefault.
func (m *TCPMuxDefault) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// GetConnByUfrag retrieves an existing or creates a new net.PacketConn.
func (m *TCPMuxDefault) GetConnByUfrag(ufrag string, isIPv6 bool, local net.IP) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), nil
}

func (m *TCPMuxDefault) createConn(ufrag string, isIPv6 bool, local net.IP, fromStun bool) (*tcpPacketConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: this is missing zone for IPv6

// Note: this is missing zone for IPv6

func (m *TCPMuxDefault) closeAndLogError(closer io.Closer) { _ = "STUB: not implemented"; return }

func (m *TCPMuxDefault) handleConn(conn net.Conn) {
	_ = "STUB: not implemented" //nolint:cyclop
	return
}

// Explicitly copy raw buffer so Message can own the memory.

// Not a STUN

// Close closes the listener and waits for all goroutines to exit.
func (m *TCPMuxDefault) Close() error { _ = "STUB: not implemented"; return nil }

// RemoveConnByUfrag closes and removes a net.PacketConn by Ufrag.
func (m *TCPMuxDefault) RemoveConnByUfrag(ufrag string) { _ = "STUB: not implemented"; return }

// Keep lock section small to avoid deadlock with conn lock

// Close the connections outside the critical section to avoid
// deadlocking TCP mux if (*tcpPacketConn).Close() blocks.

func (m *TCPMuxDefault) removeConnByUfragAndLocalHost(ufrag string, localIPAddr ipAddr) {
	_ = "STUB: not implemented"
	return
}

// Keep lock section small to avoid deadlock with conn lock

// Close the connections outside the critical section to avoid
// deadlocking TCP mux if (*tcpPacketConn).Close() blocks.

func (m *TCPMuxDefault) getConn(ufrag string, isIPv6 bool, local net.IP) (val *tcpPacketConn, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Note: this is missing zone for IPv6

const streamingPacketHeaderLen = 2

// readStreamingPacket reads 1 packet from stream
// read packet  bytes https://tools.ietf.org/html/rfc4571#section-2
// 2-byte length header prepends each packet:
//
//	 0                   1                   2                   3
//	 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//	-----------------------------------------------------------------
//	|             LENGTH            |  RTP or RTCP packet ...       |
//	-----------------------------------------------------------------
func readStreamingPacket(conn net.Conn, buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func writeStreamingPacket(conn net.Conn, buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:gosec // G115
