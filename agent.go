// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package ice implements the Interactive Connectivity Establishment (ICE)
// protocol defined in rfc5245.
package ice

import (
	"context"
	"net"
	"net/netip"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pion/ice/v4/internal/taskloop"
	"github.com/pion/logging"
	"github.com/pion/mdns/v2"
	"github.com/pion/stun/v3"
	"github.com/pion/transport/v4"
	"github.com/pion/transport/v4/packetio"
	"github.com/pion/turn/v5"
	"golang.org/x/net/proxy"
)

type bindingRequest struct {
	timestamp       time.Time
	transactionID   [stun.TransactionIDSize]byte
	destination     net.Addr
	isUseCandidate  bool
	nominationValue *uint32 // Tracks nomination value for renomination requests
}

// Agent represents the ICE agent.
type Agent struct {
	loop *taskloop.Loop

	// constructed is set to true after the agent is fully initialized.
	// Options can check this flag to reject updates that are only valid during construction.
	constructed bool

	onConnectionStateChangeHdlr       atomic.Value // func(ConnectionState)
	onSelectedCandidatePairChangeHdlr atomic.Value // func(Candidate, Candidate)
	onCandidateHdlr                   atomic.Value // func(Candidate)

	onConnected     chan struct{}
	onConnectedOnce sync.Once

	// Force candidate to be contacted immediately (instead of waiting for task ticker)
	forceCandidateContact chan bool

	tieBreaker uint64
	lite       bool

	connectionState ConnectionState
	gatheringState  GatheringState

	mDNSMode MulticastDNSMode
	mDNSName string
	mDNSConn *mdns.Conn

	muHaveStarted sync.Mutex
	startedCh     <-chan struct{}
	startedFn     func()
	isControlling atomic.Bool

	maxBindingRequests uint16

	hostAcceptanceMinWait  time.Duration
	srflxAcceptanceMinWait time.Duration
	prflxAcceptanceMinWait time.Duration
	relayAcceptanceMinWait time.Duration
	stunGatherTimeout      time.Duration

	tcpPriorityOffset uint16
	disableActiveTCP  bool

	portMin uint16
	portMax uint16

	candidateTypes []CandidateType

	// How long connectivity checks can fail before the ICE Agent
	// goes to disconnected
	disconnectedTimeout time.Duration

	// How long connectivity checks can fail before the ICE Agent
	// goes to failed
	failedTimeout time.Duration

	// How often should we send keepalive packets?
	// 0 means never
	keepaliveInterval time.Duration

	// How often should we run our internal taskLoop to check for state changes when connecting
	checkInterval time.Duration

	localUfrag      string
	localPwd        string
	localCandidates map[NetworkType][]Candidate

	remoteUfrag      string
	remotePwd        string
	remoteCandidates map[NetworkType][]Candidate

	checklist  []*CandidatePair
	nextPairID uint64
	pairsByID  map[uint64]*CandidatePair

	selectorLock sync.RWMutex
	selector     pairCandidateSelector

	selectedPair atomic.Value // *CandidatePair

	urls                   []*stun.URI
	networkTypes           []NetworkType
	turnTransportProtocols []NetworkType
	addressRewriteRules    []AddressRewriteRule

	buf *packetio.Buffer

	// LRU of outbound Binding request Transaction IDs
	pendingBindingRequests []bindingRequest

	// Address rewrite (1:1) IP mapping
	addressRewriteMapper *addressRewriteMapper

	// Callback that allows user to implement custom behavior
	// for STUN Binding Requests
	userBindingRequestHandler func(m *stun.Message, local, remote Candidate, pair *CandidatePair) bool

	gatherCandidateCancel func()
	gatherCandidateDone   chan struct{}

	connectionStateNotifier       *handlerNotifier
	candidateNotifier             *handlerNotifier
	selectedCandidatePairNotifier *handlerNotifier

	loggerFactory logging.LoggerFactory
	log           logging.LeveledLogger

	net         transport.Net
	tcpMux      TCPMux
	udpMux      UDPMux
	udpMuxSrflx UniversalUDPMux

	interfaceFilter func(string) (keep bool)
	ipFilter        func(net.IP) (keep bool)
	remoteIPFilter  func(net.IP) (keep bool)
	includeLoopback bool

	insecureSkipVerify bool

	proxyDialer proxy.Dialer

	enableUseCandidateCheckPriority bool

	// Renomination support
	enableRenomination       bool
	nominationValueGenerator func() uint32
	nominationAttribute      stun.AttrType

	// Continual gathering support
	continualGatheringPolicy ContinualGatheringPolicy
	networkMonitorInterval   time.Duration
	lastKnownInterfaces      map[string]netip.Addr // map[iface+ip] for deduplication

	// Automatic renomination
	automaticRenomination bool
	renominationInterval  time.Duration
	lastRenominationTime  time.Time

	turnClientFactory func(*turn.ClientConfig) (turnClient, error)
}

// NewAgent creates a new Agent.
//
// Deprecated: use NewAgentWithOptions instead.
func NewAgent(config *AgentConfig) (*Agent, error) { _ = "STUB: not implemented"; return nil, nil }

// NewAgentWithOptions creates a new Agent with options only.
func NewAgentWithOptions(opts ...AgentOption) (*Agent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newAgentFromConfig(config *AgentConfig, opts ...AgentOption) (*Agent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateLegacyNAT1To1IPs(ips []string) error { _ = "STUB: not implemented"; return nil }

func validateLegacyNAT1To1Entry(mapping string, hasIPv4CatchAll, hasIPv6CatchAll bool) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func legacyNAT1To1Rules(ips []string, candidateType CandidateType) ([]AddressRewriteRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createAgentBase(config *AgentConfig) (*Agent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default to GatherOnce

// Default matching libwebrtc

func applyAddressRewriteMapping(agent *Agent) error { _ = "STUB: not implemented"; return nil }

// for mDNS QueryAndGather we never advertise rewritten host IPs to avoid
// leaking local addresses, this matches the legacy NAT1:1 behavior.

// surface misconfiguration when host candidates are disabled but a host
// rewrite rule was provided.

// surface misconfiguration when srflx candidates are disabled but a srflx
// rewrite rule was provided.

// setupMDNSConfig validates and returns mDNS configuration.
func setupMDNSConfig(config *AgentConfig) (string, MulticastDNSMode, error) {
	_ = "STUB: not implemented"
	return "", *new(MulticastDNSMode), nil
}

// newAgentWithConfig finalizes a pre-configured agent with optional overrides.
//
//nolint:gocognit,cyclop
func newAgentWithConfig(agent *Agent, opts ...AgentOption) (*Agent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Opportunistic mDNS: If we can't open the connection, that's ok: we
// can continue without it.

// Make sure the buffer doesn't grow indefinitely.
// NOTE: We actually won't get anywhere close to this limit.
// SRTP will constantly read from the endpoint and drop packets if it's full.

// Restart is also used to initialize the agent for the first time

func mDNSLocalAddressFromTCPMux(tcpMux TCPMux, networkTypes []NetworkType) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

func allNetworkTypesTCP(networkTypes []NetworkType) bool { _ = "STUB: not implemented"; return false }

func localTCPAddrFromMux(tcpMux TCPMux) (*net.TCPAddr, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func mDNSLocalAddressFromIP(ip net.IP) (net.IP, bool) {
	_ = "STUB: not implemented"
	return *new(net.IP), false
}

// mdns.Config.LocalAddress has no zone support for link-local IPv6.

func (a *Agent) startConnectivityChecks(isControlling bool, remoteUfrag, remotePwd string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:contextcheck

//nolint:contextcheck

func (a *Agent) connectivityChecks() {
	_ = "STUB: not implemented" //nolint:cyclop
	return
}

// The connection is currently failed so don't send any checks
// In the future it may be restarted though

// We have just entered checking for the first time so update our checking timer

// We have been in checking longer then Disconnect+Failed timeout, set the connection to Failed

// While connecting, check candidates more frequently

// Ensure we run our task loop as quickly as the minimum of our various configured timeouts

func (a *Agent) updateConnectionState(newState ConnectionState) { _ = "STUB: not implemented"; return }

// Connection has gone to failed, release all gathered candidates

func (a *Agent) setSelectedPair(pair *CandidatePair) { _ = "STUB: not implemented"; return }

// Signal connected: notify any Connect() calls waiting on onConnected

// Update connection state to Connected and notify state change handlers

// Notify when the selected candidate pair changes

func (a *Agent) pingAllCandidates() { _ = "STUB: not implemented"; return }

// keepAliveCandidatesForRenomination pings all candidate pairs to keep them tested
// and ready for automatic renomination. Unlike pingAllCandidates, this:
// - Pings pairs in succeeded state to keep RTT measurements fresh
// - Ignores maxBindingRequests limit (we want to keep testing alternate paths)
// - Only pings pairs that are not failed.
func (a *Agent) keepAliveCandidatesForRenomination() { _ = "STUB: not implemented"; return }

// Skip failed pairs

// Transition waiting pairs to in-progress

// Continue pinging in-progress and succeeded pairs

// Ping all non-failed pairs (including succeeded ones)
// to keep RTT measurements fresh for renomination decisions

func (a *Agent) getBestAvailableCandidatePair() *CandidatePair {
	_ = "STUB: not implemented"
	return nil
}

func (a *Agent) getBestValidCandidatePair() *CandidatePair { _ = "STUB: not implemented"; return nil }

func (a *Agent) addPair(local, remote Candidate) *CandidatePair {
	_ = "STUB: not implemented"
	return nil
}

func (a *Agent) findPair(local, remote Candidate) *CandidatePair {
	_ = "STUB: not implemented"
	return nil
}

// validateSelectedPair checks if the selected pair is (still) valid
// Note: the caller should hold the agent lock.
func (a *Agent) validateSelectedPair() bool { _ = "STUB: not implemented"; return false }

// Only allow transitions to failed if a.failedTimeout is non-zero

func (a *Agent) connectionStateForDisconnection(
	disconnectedTime time.Duration,
	totalTimeToFailure time.Duration,
) ConnectionState {
	_ = "STUB: not implemented"
	return *new(ConnectionState)
}

// If we never reported disconnected but both thresholds are already exceeded,
// emit disconnected first so callers can observe both transitions.

// checkKeepalive sends STUN Binding Indications to the selected pair
// if no packet has been sent on that pair in the last keepaliveInterval
// Note: the caller should hold the agent lock.
func (a *Agent) checkKeepalive() { _ = "STUB: not implemented"; return }

// We use binding request instead of indication to support refresh consent schemas
// see https://tools.ietf.org/html/rfc7675

// AddRemoteCandidate adds a new remote candidate.
func (a *Agent) AddRemoteCandidate(cand Candidate) error { _ = "STUB: not implemented"; return nil }

// TCP Candidates with TCP type active will probe server passive ones, so
// no need to do anything with them.

// If we have a mDNS Candidate lets fully resolve it before adding it locally

// nolint: contextcheck

func (a *Agent) resolveAndAddMulticastCandidate(cand *CandidateHost) {
	_ = "STUB: not implemented"
	return
}

// nolint: contextcheck

func (a *Agent) mDNSQueryTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (a *Agent) requestConnectivityCheck() { _ = "STUB: not implemented"; return }

func (a *Agent) addRemotePassiveTCPCandidate(remoteCandidate Candidate) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // G115, no overflow, a port

func remoteDialIPForLocalInterface(remoteIP, localIP netip.Addr) netip.Addr {
	_ = "STUB: not implemented"
	return *new(netip.Addr)
}

func copyAtomicValue(dst, src *atomic.Value) { _ = "STUB: not implemented"; return }

type candidateActivitySetter interface {
	setLastReceived(time.Time)
	setLastSent(time.Time)
}

func copyCandidateActivity(dst, src Candidate) { _ = "STUB: not implemented"; return }

func replacePairRemote(pair *CandidatePair, remote Candidate) *CandidatePair {
	_ = "STUB: not implemented"
	return nil
}

func (a *Agent) retargetKnownPairHolders(oldPair, newPair *CandidatePair) {
	_ = "STUB: not implemented"
	return
}

func removeRedundantPrflxFromSet(set []Candidate, cand Candidate) ([]Candidate, []Candidate) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Agent) replaceRemoteInPairs(oldRemote, newRemote Candidate) {
	_ = "STUB: not implemented"
	return
}

func (a *Agent) replaceRemoteInLocalCaches(oldRemote, newRemote Candidate) {
	_ = "STUB: not implemented"
	return
}

// replaceRedundantPeerReflexiveCandidates removes any peer-reflexive candidates
// from the given set that have the same transport address as cand.
// It also updates any candidate pairs and local candidate caches that
// referenced the removed peer-reflexive candidates to reference cand instead.
// It is implemented according to RFC 8838 §11.4.
// It returns the updated set of candidates.
func (a *Agent) replaceRedundantPeerReflexiveCandidates(set []Candidate, cand Candidate) []Candidate {
	_ = "STUB: not implemented"
	return nil
}

// addRemoteCandidate assumes you are holding the lock (must be execute using a.run).
// Returns true when the candidate is accepted (including duplicates).
func (a *Agent) addRemoteCandidate(cand Candidate) bool {
	_ = "STUB: not implemented" //nolint:cyclop
	return false
}

// RFC 8838 §11.4: If a trickled candidate is redundant with an existing
// peer-reflexive candidate (same transport address), prefer the signaled
// candidate and replace the peer-reflexive one.

// Assert that TCP4 or TCP6 is a enabled NetworkType locally

func (a *Agent) shouldAcceptRemoteCandidate(cand Candidate) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Agent) addCandidate(ctx context.Context, cand Candidate, candidateConn net.PacketConn) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Agent) setCandidateExtensions(cand Candidate) { _ = "STUB: not implemented"; return }

// GetRemoteCandidates returns the remote candidates.
func (a *Agent) GetRemoteCandidates() ([]Candidate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLocalCandidates returns the local candidates.
func (a *Agent) GetLocalCandidates() ([]Candidate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGatheringState returns the current gathering state of the Agent.
func (a *Agent) GetGatheringState() (GatheringState, error) {
	_ = "STUB: not implemented"
	return *new(GatheringState), nil
}

// GetLocalUserCredentials returns the local user credentials.
func (a *Agent) GetLocalUserCredentials() (frag string, pwd string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// GetRemoteUserCredentials returns the remote user credentials.
func (a *Agent) GetRemoteUserCredentials() (frag string, pwd string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (a *Agent) removeUfragFromMux() { _ = "STUB: not implemented"; return }

// Close cleans up the Agent.
func (a *Agent) Close() error { _ = "STUB: not implemented"; return nil }

// GracefulClose cleans up the Agent and waits for any goroutines it started
// to complete. This is only safe to call outside of Agent callbacks or if in a callback,
// in its own goroutine.
func (a *Agent) GracefulClose() error { _ = "STUB: not implemented"; return nil }

func (a *Agent) close(graceful bool) error {
	_ = "STUB: not implemented"
	// the loop is safe to wait on no matter what
	return nil
}

// but we are in less control of the notifiers, so we will
// pass through `graceful`.

// Remove all candidates. This closes any listening sockets
// and removes both the local and remote candidate lists.
//
// This is used for restarts, failures and on close.
func (a *Agent) deleteAllCandidates() { _ = "STUB: not implemented"; return }

func (a *Agent) findRemoteCandidate(networkType NetworkType, addr net.Addr) Candidate {
	_ = "STUB: not implemented"
	return *new(Candidate)
}

func (a *Agent) sendBindingRequest(msg *stun.Message, local, remote Candidate) {
	_ = "STUB: not implemented"
	return
}

// Extract nomination value if present

func (a *Agent) sendBindingSuccess(m *stun.Message, local, remote Candidate) {
	_ = "STUB: not implemented"
	return
}

// Removes pending binding requests that are over maxBindingRequestTimeout old
//
// Let HTO be the transaction timeout, which SHOULD be 2*RTT if
// RTT is known or 500 ms otherwise.
// https://tools.ietf.org/html/rfc8445#appendix-B.1
func (a *Agent) invalidatePendingBindingRequests(filterTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// Assert that the passed TransactionID is in our pendingBindingRequests and returns the destination
// If the bindingRequest was valid remove it from our pending cache.
func (a *Agent) handleInboundBindingSuccess(id [stun.TransactionIDSize]byte) (bool, *bindingRequest, time.Duration) {
	_ = "STUB: not implemented"
	return false, nil, *new(time.Duration)
}

func (a *Agent) handleRoleConflict(msg *stun.Message, local, remote Candidate, remoteTieBreaker *AttrControl) {
	_ = "STUB: not implemented"
	return
}

// https://datatracker.ietf.org/doc/html/rfc8445#section-7.3.1.1
//  An agent MUST examine the Binding request for either the ICE-
//  CONTROLLING or ICE-CONTROLLED attribute.  It MUST follow these
// procedures:

// If the agent's tiebreaker value is larger than or equal to the contents of the ICE-CONTROLLING attribute
// If the agent's tiebreaker value is less than the contents of the ICE-CONTROLLED attribute
//  the agent generates a Binding error response

// handleInbound processes STUN traffic from a remote candidate.
func (a *Agent) handleInbound(msg *stun.Message, local Candidate, remote net.Addr) {
	_ = "STUB: not implemented"
	return
}

func canHandleInbound(msg *stun.Message) bool { _ = "STUB: not implemented"; return false }

func (a *Agent) handleInboundResponse(
	remoteCandidate, local Candidate, remote net.Addr, msg *stun.Message,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Agent) handleInboundRequest(
	remoteCandidate, local Candidate, remote net.Addr, msg *stun.Message,
) (remoteCand Candidate, ok bool) {
	_ = "STUB: not implemented"
	return *new(Candidate), false
}

// A peer-reflexive candidate SHOULD take its priority from the PRIORITY
// attribute in the Binding Request that discovered it.

// Support Remotes that don't set a TIE-BREAKER. Not standards compliant, but
// keeping to maintain backwards compat

// validateNonSTUNTraffic processes non STUN traffic from a remote candidate,
// and returns true if it is an actual remote candidate.
func (a *Agent) validateNonSTUNTraffic(local Candidate, remote net.Addr) (Candidate, bool) {
	_ = "STUB: not implemented"
	return *new(Candidate), false
}

// GetSelectedCandidatePair returns the selected pair or nil if there is none.
func (a *Agent) GetSelectedCandidatePair() (*CandidatePair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nilnil

func (a *Agent) getSelectedPair() *CandidatePair { _ = "STUB: not implemented"; return nil }

func (a *Agent) closeMulticastConn() { _ = "STUB: not implemented"; return }

// SetRemoteCredentials sets the credentials of the remote agent.
func (a *Agent) SetRemoteCredentials(remoteUfrag, remotePwd string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateOptions applies the given options to the agent at runtime.
// Only a subset of options can be updated after agent creation:
//   - WithUrls: updates STUN/TURN server URLs (takes effect on next GatherCandidates call)
//
// Returns an error if the agent is closed or if an unsupported option is provided.
func (a *Agent) UpdateOptions(opts ...AgentOption) error { _ = "STUB: not implemented"; return nil }

// Restart restarts the ICE Agent with the provided ufrag/pwd
// If no ufrag/pwd is provided the Agent will generate one itself
//
// If there is a gatherer routine currently running, Restart will
// cancel it.
// After a Restart, the user must then call GatherCandidates explicitly
// to start generating new ones.
func (a *Agent) Restart(ufrag, pwd string) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

// Clear all agent needed to take back to fresh state

// Restart is used by NewAgent. Accept/Connect should be used to move to checking
// for new Agents

func (a *Agent) setGatheringState(newState GatheringState) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Agent) needsToCheckPriorityOnNominated() bool { _ = "STUB: not implemented"; return false }

func (a *Agent) role() Role { _ = "STUB: not implemented"; return *new(Role) }

func (a *Agent) setSelector() { _ = "STUB: not implemented"; return }

func (a *Agent) getSelector() pairCandidateSelector {
	_ = "STUB: not implemented"
	return *new(pairCandidateSelector)
}

// getNominationValue returns a nomination value if generator is available, otherwise 0.
func (a *Agent) getNominationValue() uint32 { _ = "STUB: not implemented"; return 0 }

// RenominateCandidate allows the controlling ICE agent to nominate a new candidate pair.
// This implements the continuous renomination feature from draft-thatcher-ice-renomination-01.
func (a *Agent) RenominateCandidate(local, remote Candidate) error {
	_ = "STUB: not implemented"
	return nil
}

// Find the candidate pair

// Send nomination with custom attribute

// sendNominationRequest sends a nomination request with custom nomination value.
func (a *Agent) sendNominationRequest(pair *CandidatePair, nominationValue uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// Add nomination attribute if renomination is enabled and value > 0

// evaluateCandidatePairQuality calculates a quality score for a candidate pair.
// Higher scores indicate better quality. The score considers:
// - Candidate types (host > srflx > relay)
// - RTT (lower is better)
// - Connection stability.
func (a *Agent) evaluateCandidatePairQuality(pair *CandidatePair) float64 {
	_ = "STUB: not implemented" //nolint:cyclop
	return 0
}

// Type preference scoring (host=100, srflx=50, prflx=30, relay=10)

// Combined type score (average of local and remote)

// RTT scoring (convert to penalty, lower RTT = higher score)
// Use current RTT if available, otherwise assume high latency

// Convert RTT to Duration for cleaner calculation

// Minimum 1ms to avoid log(0)

// Subtract RTT penalty (logarithmic to reduce impact of very high RTTs)

// No RTT data available, apply moderate penalty

// Boost score if pair has been stable (received responses recently)

// Stability bonus

// shouldRenominate determines if automatic renomination should occur.
// It compares the current selected pair with a candidate pair and decides
// if switching would provide significant benefit.
func (a *Agent) shouldRenominate(current, candidate *CandidatePair) bool {
	_ = "STUB: not implemented" //nolint:cyclop
	return false
}

// Type-based switching (always prefer direct over relay)

// RTT-based switching (must improve by at least 10ms)

// Only compare RTT if both values are valid

// Quality score comparison (must improve by at least 15%)

// findBestCandidatePair finds the best available candidate pair based on quality assessment.
func (a *Agent) findBestCandidatePair() *CandidatePair { _ = "STUB: not implemented"; return nil }
