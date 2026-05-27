// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"io"
	"net"
	"net/netip"
	"sync"

	"github.com/pion/logging"
	"github.com/pion/transport/v4"
)

// UDPMux allows multiple connections to go over a single UDP port.
type UDPMux interface {
	io.Closer
	GetConn(ufrag string, addr net.Addr) (net.PacketConn, error)
	RemoveConnByUfrag(ufrag string)
	GetListenAddresses() []net.Addr
}

// UDPMuxDefault is an implementation of the interface.
type UDPMuxDefault struct {
	params UDPMuxParams

	closedChan chan struct{}
	closeOnce  sync.Once

	// connsIPv4 and connsIPv6 are maps of all udpMuxedConn indexed by ufrag|network|candidateType
	connsIPv4, connsIPv6 map[string]*udpMuxedConn

	addressMapMu sync.RWMutex
	addressMap   map[ipPort]*udpMuxedConn

	// Buffer pool to recycle buffers for net.UDPAddr encodes/decodes
	pool *sync.Pool

	mu sync.Mutex

	// For UDP connection listen at unspecified address
	localAddrsForUnspecified []net.Addr
}

// UDPMuxParams are parameters for UDPMux.
type UDPMuxParams struct {
	Logger        logging.LeveledLogger
	UDPConn       net.PacketConn
	UDPConnString string

	// Required for gathering local addresses
	// in case a un UDPConn is passed which does not
	// bind to a specific local address.
	Net transport.Net
}

// NewUDPMuxDefault creates an implementation of UDPMux.
func NewUDPMuxDefault(params UDPMuxParams) *UDPMuxDefault {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

//nolint:nestif

// For unspecified addresses, the correct behavior is to return errListenUnspecified, but
// it will break the applications that are already using unspecified UDP connection
// with UDPMuxDefault, so print a warn log and create a local address list for mux.

// Big enough buffer to fit both packet and address

// LocalAddr returns the listening address of this UDPMuxDefault.
func (m *UDPMuxDefault) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// GetListenAddresses returns the list of addresses that this mux is listening on.
func (m *UDPMuxDefault) GetListenAddresses() []net.Addr { _ = "STUB: not implemented"; return nil }

// GetConn returns a PacketConn given the connection's ufrag and network address.
// creates the connection if an existing one can't be found.
func (m *UDPMuxDefault) GetConn(ufrag string, addr net.Addr) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	// don't check addr for mux using unspecified address
	return *new(net.PacketConn), nil
}

// RemoveConnByUfrag stops and removes the muxed packet connection.
func (m *UDPMuxDefault) RemoveConnByUfrag(ufrag string) { _ = "STUB: not implemented"; return }

// Keep lock section small to avoid deadlock with conn lock.

// No need to lock if no connection was found.

// IsClosed returns true if the mux had been closed.
func (m *UDPMuxDefault) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Close the mux, no further connections could be created.
func (m *UDPMuxDefault) Close() error { _ = "STUB: not implemented"; return nil }

func (m *UDPMuxDefault) writeTo(buf []byte, rAddr net.Addr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *UDPMuxDefault) registerConnForAddress(conn *udpMuxedConn, addr ipPort) {
	_ = "STUB: not implemented"
	return
}

func (m *UDPMuxDefault) createMuxedConn(key string) *udpMuxedConn {
	_ = "STUB: not implemented"
	return nil
}

func (m *UDPMuxDefault) connWorker() {
	_ = "STUB: not implemented" //nolint:cyclop
	return
}

//nolint:gosec

// If we have already seen this address dispatch to the appropriate destination

// If we haven't seen this address before but is a STUN packet lookup by ufrag

func (m *UDPMuxDefault) getConn(ufrag string, isIPv6 bool) (val *udpMuxedConn, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type bufferHolder struct {
	next *bufferHolder
	buf  []byte
	addr *net.UDPAddr
}

func newBufferHolder(size int) *bufferHolder { _ = "STUB: not implemented"; return nil }

func (b *bufferHolder) reset() { _ = "STUB: not implemented"; return }

type ipPort struct {
	addr netip.Addr
	port uint16
}

// newIPPort create a custom type of address based on netip.Addr and
// port. The underlying ip address passed is converted to IPv6 format
// to simplify ip address handling.
func newIPPort(ip net.IP, zone string, port uint16) (ipPort, error) {
	_ = "STUB: not implemented"
	return *new(ipPort), nil
}
