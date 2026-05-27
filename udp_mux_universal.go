// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"net"
	"time"

	"github.com/pion/logging"
	"github.com/pion/stun/v3"
	"github.com/pion/transport/v4"
)

// UniversalUDPMux allows multiple connections to go over a single UDP port for
// host, server reflexive and relayed candidates.
// Actual connection muxing is happening in the UDPMux.
type UniversalUDPMux interface {
	UDPMux
	GetXORMappedAddr(stunAddr net.Addr, deadline time.Duration) (*stun.XORMappedAddress, error)
	GetRelayedAddr(turnAddr net.Addr, deadline time.Duration) (*net.Addr, error)
	GetConnForURL(ufrag string, url string, addr net.Addr) (net.PacketConn, error)
}

// UniversalUDPMuxDefault handles STUN and TURN servers packets by wrapping the original UDPConn overriding ReadFrom.
// It the passes packets to the UDPMux that does the actual connection muxing.
type UniversalUDPMuxDefault struct {
	*UDPMuxDefault
	params UniversalUDPMuxParams

	// Since we have a shared socket, for srflx candidates it makes sense
	// to have a shared mapped address across all the agents
	// stun.XORMappedAddress indexed by the STUN server addr
	xorMappedMap map[string]*xorMapped
}

// UniversalUDPMuxParams are parameters for UniversalUDPMux server reflexive.
type UniversalUDPMuxParams struct {
	Logger                logging.LeveledLogger
	UDPConn               net.PacketConn
	XORMappedAddrCacheTTL time.Duration
	Net                   transport.Net
}

// NewUniversalUDPMuxDefault creates an implementation of UniversalUDPMux embedding UDPMux.
func NewUniversalUDPMuxDefault(params UniversalUDPMuxParams) *UniversalUDPMuxDefault {
	_ = "STUB: not implemented"
	return nil
}

// Wrap UDP connection, process server reflexive messages
// before they are passed to the UDPMux connection handler (connWorker)

// Embed UDPMux

// udpConn is a wrapper around UDPMux conn that overrides ReadFrom and handles STUN/TURN packets.
type udpConn struct {
	net.PacketConn
	mux    *UniversalUDPMuxDefault
	logger logging.LeveledLogger
}

// GetRelayedAddr creates relayed connection to the given TURN service and returns the relayed addr.
// Not implemented yet.
func (m *UniversalUDPMuxDefault) GetRelayedAddr(net.Addr, time.Duration) (*net.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// GetConnForURL add uniques to the muxed connection by concatenating ufrag and URL
	// (e.g. STUN URL) to be able to support multiple STUN/TURN servers
	// and return a unique connection per server.
}

func (m *UniversalUDPMuxDefault) GetConnForURL(ufrag string, url string, addr net.Addr) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), nil
}

// ReadFrom is called by UDPMux connWorker and handles packets coming from the STUN server discovering a mapped address.
// It passes processed packets further to the UDPMux (maybe this is not really necessary).
func (c *udpConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

//nolint:nestif

// Message about this err will be logged in the UDPMux

// isXORMappedResponse indicates whether the message is a XORMappedAddress and is coming from the known STUN server.
func (m *UniversalUDPMuxDefault) isXORMappedResponse(msg *stun.Message, stunAddr string) bool {
	_ = "STUB: not implemented"
	return false
}

// Check first if it is a STUN server address,
// because remote peer can also send similar messages but as a BindingSuccess.

// handleXORMappedResponse parses response from the STUN server, extracts XORMappedAddress attribute.
// and set the mapped address for the server.
func (m *UniversalUDPMuxDefault) handleXORMappedResponse(stunAddr *net.UDPAddr, msg *stun.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// GetXORMappedAddr returns *stun.XORMappedAddress if already present for a given STUN server.
// Makes a STUN binding request to discover mapped address otherwise.
// Blocks until the stun.XORMappedAddress has been discovered or deadline.
// Method is safe for concurrent use.
func (m *UniversalUDPMuxDefault) GetXORMappedAddr(
	serverAddr net.Addr,
	deadline time.Duration,
) (*stun.XORMappedAddress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we already have a mapping for this STUN server (address already received)
// and if it is not too old we return it without making a new request to STUN server

// Otherwise, make a STUN request to discover the address
// or wait for already sent request to complete

//nolint:errorlint

// Block until response was handled by the connWorker routine and XORMappedAddress was updated

// When channel closed, addr was obtained

// writeSTUN sends a STUN request via UDP conn.
//
// The returned channel is closed when the STUN response has been received.
// Method is safe for concurrent use.
func (m *UniversalUDPMuxDefault) writeSTUN(serverAddr net.Addr) (chan struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If record present in the map, we already sent a STUN request,
// just wait when waitAddrReceived will be closed

type xorMapped struct {
	addr             *stun.XORMappedAddress
	waitAddrReceived chan struct{}
	expiresAt        time.Time
}

func (a *xorMapped) closeWaiters() { _ = "STUB: not implemented"; return }

// Notify was close, ok, that means we received duplicate response just exit

// Notify tha twe have a new addr

func (a *xorMapped) pending() bool { _ = "STUB: not implemented"; return false }

func (a *xorMapped) expired() bool { _ = "STUB: not implemented"; return false }

func (a *xorMapped) SetAddr(addr *stun.XORMappedAddress) { _ = "STUB: not implemented"; return }
