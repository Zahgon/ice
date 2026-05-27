// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"context"
	"io"
	"net"
	"net/netip"
	"sync"

	"github.com/pion/logging"
	"github.com/pion/stun/v3"
	"github.com/pion/turn/v5"
)

type turnClient interface {
	Listen() error
	Allocate() (net.PacketConn, error)
	Close()
}

func defaultTurnClient(cfg *turn.ClientConfig) (turnClient, error) {
	_ = "STUB: not implemented"
	return *new(turnClient), nil
}

func configuredNetworkTypes(networkTypes []NetworkType) []NetworkType {
	_ = "STUB: not implemented"
	return nil
}

func effectiveURLProtoType(url stun.URI) stun.ProtoType {
	_ = "STUB: not implemented"
	return *new(stun.ProtoType)
}

func urlSupportsSrflxGathering(url stun.URI) bool { _ = "STUB: not implemented"; return false }

func relayNetworkTypesForConfiguredCandidates(networkTypes []NetworkType) []NetworkType {
	_ = "STUB: not implemented"
	// Relay allocations currently produce UDP relay endpoints, so relay candidate
	// publication must be gated by configured UDP candidate network types.
	return nil
}

func turnNetworkTypesForURL(url stun.URI, networkTypes []NetworkType) []NetworkType {
	_ = "STUB: not implemented"
	return nil
}

// Close a net.Conn and log if we have a failure.
func closeConnAndLog(c io.Closer, log logging.LeveledLogger, msg string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// GatherCandidates initiates the trickle based gathering process.
func (a *Agent) GatherCandidates() error { _ = "STUB: not implemented"; return nil }

// Cancel previous gathering routine

func (a *Agent) gatherCandidates(ctx context.Context, done chan struct{}) {
	_ = "STUB: not implemented" //nolint:cyclop
	return
}

//nolint:contextcheck

//nolint:contextcheck

// Initialize known interfaces before starting monitoring

func (a *Agent) shouldRewriteCandidateType(candidateType CandidateType) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Agent) shouldRewriteHostCandidates() bool { _ = "STUB: not implemented"; return false }

func (a *Agent) applyHostAddressRewrite(addr netip.Addr, mappedAddrs []netip.Addr, iface string) ([]netip.Addr, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func appendHostMappedAddrs(
	mappedAddrs []netip.Addr,
	mappedIPs []net.IP,
	addr netip.Addr,
	log logging.LeveledLogger,
) []netip.Addr {
	_ = "STUB: not implemented"
	return nil
}

// we'd rather have an IPv4-mapped IPv6 become IPv4 so that it is usable

func (a *Agent) applyHostRewriteForUDPMux(candidateIPs []net.IP, udpAddr *net.UDPAddr) ([]net.IP, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// gatherCandidatesInternal performs the actual candidate gathering for all configured types.
func (a *Agent) gatherCandidatesInternal(ctx context.Context) { _ = "STUB: not implemented"; return }

// Block until all STUN and TURN URLs have been gathered (or timed out)

func (a *Agent) gatherServerReflexiveCandidates(ctx context.Context, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

//nolint:gocognit,gocyclo,cyclop,maintidx
func (a *Agent) gatherCandidatesLocal(ctx context.Context, networkTypes []NetworkType) {
	_ = "STUB: not implemented"
	return
}

// When UDPMux is enabled, skip other UDP candidates

// Here, we are not doing multicast gathering, so we will need to skip this address so
// that we don't accidentally reveal location tracking information. Otherwise, the
// case above hides the IP behind an mDNS address.

// TCPMux maintains a single listener per interface. Avoid duplicating passive TCP candidates
// for additional mapped IPs until connection sharing is supported.

// Only advertise TCP candidates for addresses that the mux listener is actually
// bound to. When the listener is bound to a specific IP, exposing other interface
// addresses would generate unreachable passive candidates and can stall active
// TCP connect attempts.

// Handle ICE TCP passive mode

// Note: this is missing zone for IPv6 by just grabbing the IP slice

// Note: this is missing zone for IPv6 by just grabbing the IP slice

// Extract the port for each PacketConn we got.

// Didn't succeed with any, try the next network.

// Is there a way to verify that the listen address is even
// accessible from the current interface.

// we will still process this candidate so that we start up the right
// listeners.

// shouldFilterLocationTrackedIP returns if this candidate IP should be filtered out from
// any candidate publishing/notification for location tracking reasons.
func shouldFilterLocationTrackedIP(candidateIP netip.Addr) bool {
	_ = "STUB: not implemented"
	// https://tools.ietf.org/html/rfc8445#section-5.1.1.1
	// Similarly, when host candidates corresponding to
	// an IPv6 address generated using a mechanism that prevents location
	// tracking are gathered, then host candidates corresponding to IPv6
	// link-local addresses [RFC4291] MUST NOT be gathered.
	return false
}

// shouldFilterLocationTracked returns if this candidate IP should be filtered out from
// any candidate publishing/notification for location tracking reasons.
func shouldFilterLocationTracked(candidateIP net.IP) bool { _ = "STUB: not implemented"; return false }

func (a *Agent) gatherCandidatesLocalUDPMux(ctx context.Context) error {
	_ = "STUB: not implemented" //nolint:gocognit,cyclop
	return nil
}

// Unlike MultiUDPMux Default, UDPMuxDefault doesn't have
// a separate param to include loopback, so we respect agent config

// Here, we are not doing multicast gathering, so we will need to skip this address so
// that we don't accidentally reveal location tracking information. Otherwise, the
// case above hides the IP behind an mDNS address.

// Detect a duplicate candidate before calling addCandidate().
// otherwise, addCandidate() detects the duplicate candidate
// and close its connection, invalidating all candidates
// that share the same connection.

func (a *Agent) gatherCandidatesSrflxMapped(ctx context.Context, networkTypes []NetworkType) {
	_ = "STUB: not implemented" //nolint:gocognit,cyclop
	return
}

//nolint:gocognit,cyclop
func (a *Agent) gatherCandidatesSrflxUDPMux(ctx context.Context, urls []*stun.URI, networkTypes []NetworkType) {
	_ = "STUB: not implemented"
	return
}

//nolint:cyclop,gocognit
func (a *Agent) gatherCandidatesSrflx(ctx context.Context, urls []*stun.URI, networkTypes []NetworkType) {
	_ = "STUB: not implemented"
	return
}

// If the agent closes midway through the connection
// we end it early to prevent close delay.

//nolint:forcetypeassert

//nolint:maintidx,gocognit,gocyclo,cyclop
func (a *Agent) gatherCandidatesRelay(ctx context.Context, urls []*stun.URI) {
	_ = "STUB: not implemented"
	return
}

// IPv6 TURN support is not finished yet, so skip for now.

// nolint:nestif

//nolint:forcetypeassert
//nolint:forcetypeassert

//nolint:forcetypeassert
//nolint:forcetypeassert

//nolint:forcetypeassert
//nolint:forcetypeassert

//nolint:gosec

//nolint:forcetypeassert
//nolint:forcetypeassert

//nolint:gosec

//nolint:forcetypeassert
//nolint:forcetypeassert

//nolint:forcetypeassert

// Relay allocations currently produce UDP relay endpoints regardless of
// whether the TURN control connection uses UDP/TCP/TLS/DTLS.

type relayEndpoint struct {
	network   string
	address   net.IP
	port      int
	relAddr   string
	relPort   int
	protocol  string
	iface     string
	conn      net.PacketConn
	onClose   func() error
	closeConn func()
}

func (a *Agent) resolveRelayAddresses(ep relayEndpoint) ([]net.IP, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a *Agent) resolveSrflxAddresses(localIP net.IP, iface string) ([]net.IP, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func findIfaceForIP(ifaces []ifaceAddr, ip net.IP) string { _ = "STUB: not implemented"; return "" }

func (a *Agent) createRelayCandidate(ctx context.Context, ep relayEndpoint, ip net.IP, onClose func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Agent) addRelayCandidates(ctx context.Context, ep relayEndpoint) {
	_ = "STUB: not implemented"
	return
}

// startNetworkMonitoring starts a goroutine that periodically checks for network changes
// and re-gathers candidates when changes are detected. This is only used with GatherContinually policy.
func (a *Agent) startNetworkMonitoring(ctx context.Context) { _ = "STUB: not implemented"; return }

// detectNetworkChanges checks if the network interfaces have changed since the last check.
func (a *Agent) detectNetworkChanges() bool {
	_ = "STUB: not implemented"
	// Try to refresh interfaces if using stdnet
	return false
}
