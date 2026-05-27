// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"net"
	"time"

	"github.com/pion/logging"
	"github.com/pion/stun/v3"
	"github.com/pion/transport/v4"
	"golang.org/x/net/proxy"
)

// AgentOption represents a function that can be used to configure an Agent.
type AgentOption func(*Agent) error

// NominationValueGenerator is a function that generates nomination values for renomination.
type NominationValueGenerator func() uint32

// DefaultNominationValueGenerator returns a generator that starts at 1 and increments for each call.
// This provides a simple, monotonically increasing sequence suitable for renomination.
func DefaultNominationValueGenerator() NominationValueGenerator {
	_ = "STUB: not implemented"
	return *new(NominationValueGenerator)
}

// WithAddressRewriteRules appends the provided address rewrite (1:1) rules to the agent's
// existing configuration. Each `AddressRewriteRule` can limit the mapping to a specific
// interface (`Iface`), local address (`Local`), CIDR block (`CIDR`), or subset
// of network types (`Networks`), allowing fine-grained control over which local
// addresses are replaced with the supplied external IPs.
// Use `Mode` to control whether a rule replaces the original candidate (default for
// host) or appends additional candidates (default for other types).
//
// Rules are evaluated in the order they are added; for each candidate type +
// local address, explicit `Local` matches win immediately. Otherwise, the most
// specific catch-all is chosen (iface+CIDR > iface-only > CIDR-only > global),
// with declaration order breaking ties at the same specificity. `Iface` (when
// set) must also match. This lets you layer specificity (e.g., iface+CIDR, then
// iface-only, then global) while still keeping rule order meaningful.
// Overlapping rules in the same scope are logged as warnings.
func WithAddressRewriteRules(rules ...AddressRewriteRule) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

func warnOnAddressRewriteConflicts(agent *Agent) { _ = "STUB: not implemented"; return }

func emptyScopeValue(v string) string { _ = "STUB: not implemented"; return "" }

func appendAddressRewriteRules(agent *Agent, rules ...AddressRewriteRule) error {
	_ = "STUB: not implemented"
	return nil
}

func sanitizeAddressRewriteRule(rule AddressRewriteRule) (AddressRewriteRule, error) {
	_ = "STUB: not implemented"
	return *new(AddressRewriteRule), nil
}

func defaultAddressRewriteMode(candidateType CandidateType) AddressRewriteMode {
	_ = "STUB: not implemented"
	return *new(AddressRewriteMode)
}

func sanitizeExternalIPs(ips []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type addressRewriteScopeKey struct {
	candidateType CandidateType
	iface         string
	cidr          string
	networksKey   string
	localKey      string
}

type addressRewriteConflict struct {
	scope               addressRewriteScopeKey
	existingExternalIPs []string
	conflictingExternal string
}

func findAddressRewriteRuleConflicts(rules []AddressRewriteRule) []addressRewriteConflict {
	_ = "STUB: not implemented"
	return nil
}

type addressRewriteExternalEntry struct {
	externalIP    string
	localScopeKey string
}

func enumerateAddressRewriteExternalEntries(rule AddressRewriteRule) []addressRewriteExternalEntry {
	_ = "STUB: not implemented"
	return nil
}

func deriveAddressRewriteLocalScopeKey(local string) string { _ = "STUB: not implemented"; return "" }

func deriveAddressRewriteFamilyScopeKey(ipStr string) string { _ = "STUB: not implemented"; return "" }

func mapKeys(m map[string]struct{}) []string { _ = "STUB: not implemented"; return nil }

// WithICELite configures whether the agent operates in lite mode.
// Lite agents do not perform connectivity checks and only provide host candidates.
func WithICELite(lite bool) AgentOption { _ = "STUB: not implemented"; return *new(AgentOption) }

// WithUrls sets the STUN/TURN server URLs used by the agent.
func WithUrls(urls []*stun.URI) AgentOption { _ = "STUB: not implemented"; return *new(AgentOption) }

// WithPortRange sets the UDP port range for host candidates.
func WithPortRange(portMin, portMax uint16) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithDisconnectedTimeout sets the duration before the agent transitions to disconnected state.
// A timeout of 0 disables the transition.
func WithDisconnectedTimeout(timeout time.Duration) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithFailedTimeout sets the duration before the agent transitions to failed state after disconnected.
// A timeout of 0 disables the transition.
func WithFailedTimeout(timeout time.Duration) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithKeepaliveInterval sets how often ICE keepalive packets are sent.
// An interval of 0 disables keepalives.
func WithKeepaliveInterval(interval time.Duration) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithHostAcceptanceMinWait sets the minimum wait before selecting host candidates.
func WithHostAcceptanceMinWait(wait time.Duration) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithSrflxAcceptanceMinWait sets the minimum wait before selecting srflx candidates.
func WithSrflxAcceptanceMinWait(wait time.Duration) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithPrflxAcceptanceMinWait sets the minimum wait before selecting prflx candidates.
func WithPrflxAcceptanceMinWait(wait time.Duration) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithRelayAcceptanceMinWait sets the minimum wait before selecting relay candidates.
func WithRelayAcceptanceMinWait(wait time.Duration) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithSTUNGatherTimeout sets the STUN gather timeout.
func WithSTUNGatherTimeout(timeout time.Duration) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithIPFilter sets a filter for IP addresses used during candidate gathering.
func WithIPFilter(filter func(net.IP) bool) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithRemoteIPFilter sets a filter for remote candidate IP addresses.
// Candidates for which this function returns false are ignored.
func WithRemoteIPFilter(filter func(net.IP) bool) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithNet sets the underlying network implementation for the agent.
func WithNet(net transport.Net) AgentOption { _ = "STUB: not implemented"; return *new(AgentOption) }

// WithMulticastDNSMode configures mDNS behavior for the agent.
func WithMulticastDNSMode(mode MulticastDNSMode) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithMulticastDNSHostName sets the mDNS host name used by the agent.
func WithMulticastDNSHostName(hostName string) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithLocalCredentials sets the local ICE username fragment and password used during Restart.
// If empty strings are provided, the agent will generate values during Restart.
func WithLocalCredentials(ufrag, pwd string) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

//nolint:varnamelen

// WithTCPMux sets the TCP mux for ICE TCP multiplexing.
func WithTCPMux(tcpMux TCPMux) AgentOption { _ = "STUB: not implemented"; return *new(AgentOption) }

// WithUDPMux sets the UDP mux used for multiplexing host candidates.
func WithUDPMux(udpMux UDPMux) AgentOption { _ = "STUB: not implemented"; return *new(AgentOption) }

// WithUDPMuxSrflx sets the UDP mux for server reflexive candidates.
func WithUDPMuxSrflx(udpMuxSrflx UniversalUDPMux) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithProxyDialer sets the proxy dialer used for TURN over TCP/TLS/DTLS connections.
func WithProxyDialer(dialer proxy.Dialer) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithMaxBindingRequests sets the maximum number of binding requests before considering a pair failed.
func WithMaxBindingRequests(limit uint16) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithCheckInterval sets how often the agent runs connectivity checks while connecting.
func WithCheckInterval(interval time.Duration) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithRenomination enables ICE renomination as described in draft-thatcher-ice-renomination-01.
// When enabled, the controlling agent can renominate candidate pairs multiple times
// and the controlled agent follows "last nomination wins" rule.
//
// The generator parameter specifies how nomination values are generated.
// Use DefaultNominationValueGenerator() for a simple incrementing counter,
// or provide a custom generator for more complex scenarios.
//
// Example:
//
//	agent, err := NewAgentWithOptions(config, WithRenomination(DefaultNominationValueGenerator()))
func WithRenomination(generator NominationValueGenerator) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithNominationAttribute sets the STUN attribute type to use for ICE renomination.
// The default value is 0xC001. This can be configured until the attribute is officially
// assigned by IANA for draft-thatcher-ice-renomination.
//
// This option returns an error if the provided attribute type is invalid.
// Currently, validation ensures the attribute is not 0x0000 (reserved).
// Additional validation may be added in the future.
func WithNominationAttribute(attrType uint16) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// Basic validation: ensure it's not the reserved 0x0000

// WithIncludeLoopback includes loopback addresses in the candidate list.
// By default, loopback addresses are excluded.
//
// Example:
//
//	agent, err := NewAgentWithOptions(WithIncludeLoopback())
func WithIncludeLoopback() AgentOption { _ = "STUB: not implemented"; return *new(AgentOption) }

// WithTCPPriorityOffset sets a number which is subtracted from the default (UDP) candidate type preference
// for host, srflx and prfx candidate types. It helps to configure relative preference of UDP candidates
// against TCP ones. Relay candidates for TCP and UDP are always 0 and not affected by this setting.
// When not set, defaultTCPPriorityOffset (27) is used.
//
// Example:
//
//	agent, err := NewAgentWithOptions(WithTCPPriorityOffset(50))
func WithTCPPriorityOffset(offset uint16) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithDisableActiveTCP disables Active TCP candidates.
// When TCP is enabled, Active TCP candidates will be created when a new passive TCP remote candidate is added
// unless this option is used.
//
// Example:
//
//	agent, err := NewAgentWithOptions(WithDisableActiveTCP())
func WithDisableActiveTCP() AgentOption { _ = "STUB: not implemented"; return *new(AgentOption) }

// WithBindingRequestHandler sets a handler to allow applications to perform logic on incoming STUN Binding Requests.
// This was implemented to allow users to:
//   - Log incoming Binding Requests for debugging
//   - Implement draft-thatcher-ice-renomination
//   - Implement custom CandidatePair switching logic
//
// Example:
//
//	handler := func(m *stun.Message, local, remote Candidate, pair *CandidatePair) bool {
//		log.Printf("Binding request from %s to %s", remote.Address(), local.Address())
//		return true // Accept the request
//	}
//	agent, err := NewAgentWithOptions(WithBindingRequestHandler(handler))
func WithBindingRequestHandler(
	handler func(m *stun.Message, local, remote Candidate, pair *CandidatePair) bool,
) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithEnableUseCandidateCheckPriority enables checking for equal or higher priority when
// switching selected candidate pair if the peer requests USE-CANDIDATE and agent is a lite agent.
// This is disabled by default, i.e. when peer requests USE-CANDIDATE, the selected pair will be
// switched to that irrespective of relative priority between current selected pair
// and priority of the pair being switched to.
//
// Example:
//
//	agent, err := NewAgentWithOptions(WithEnableUseCandidateCheckPriority())
func WithEnableUseCandidateCheckPriority() AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithContinualGatheringPolicy sets the continual gathering policy for the agent.
// When set to GatherContinually, the agent will continuously monitor network interfaces
// and gather new candidates as they become available.
// When set to GatherOnce (default), gathering completes after the initial phase.
//
// Example:
//
//	agent, err := NewAgentWithOptions(WithContinualGatheringPolicy(GatherContinually))
func WithContinualGatheringPolicy(policy ContinualGatheringPolicy) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithNetworkMonitorInterval sets the interval at which the agent checks for network interface changes
// when using GatherContinually policy. This option only has effect when used with
// WithContinualGatheringPolicy(GatherContinually).
// Default is 2 seconds if not specified.
//
// Example:
//
//	agent, err := NewAgentWithOptions(
//		WithContinualGatheringPolicy(GatherContinually),
//		WithNetworkMonitorInterval(5 * time.Second),
//	)
func WithNetworkMonitorInterval(interval time.Duration) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithNetworkTypes sets the enabled candidate network types for candidate gathering.
// This controls the network types exposed in ICE candidates and used for pairing.
// Use WithTURNTransportProtocols to control the local TURN client-to-server transport.
// By default, all network types are enabled.
//
// Example:
//
//	agent, err := NewAgentWithOptions(
//		WithNetworkTypes([]NetworkType{NetworkTypeUDP4, NetworkTypeUDP6}),
//	)
func WithNetworkTypes(networkTypes []NetworkType) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithTURNTransportProtocols restricts protocols used by this agent when
// connecting to TURN servers (TURN client <-> TURN server transport).
//
// This is independent from WithNetworkTypes, which controls ICE candidate
// network types announced to the peer. Supported values are
// NetworkTypeUDP4/UDP6 and NetworkTypeTCP4/TCP6.
func WithTURNTransportProtocols(protocols []NetworkType) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

func sanitizeTransportNetworkTypes(types []NetworkType) ([]NetworkType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithCandidateTypes sets the enabled candidate types for gathering.
// By default, host, server reflexive, and relay candidates are enabled.
//
// Example:
//
//	agent, err := NewAgentWithOptions(
//		WithCandidateTypes([]CandidateType{CandidateTypeHost, CandidateTypeServerReflexive}),
//	)
func WithCandidateTypes(candidateTypes []CandidateType) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithAutomaticRenomination enables automatic renomination of candidate pairs
// when better pairs become available after initial connection establishment.
// This feature requires renomination to be enabled and both agents to support it.
//
// When enabled, the controlling agent will periodically evaluate candidate pairs
// and renominate if a significantly better pair is found (e.g., switching from
// relay to direct connection, or when RTT improves significantly).
//
// The interval parameter specifies the minimum time to wait after connection
// before considering automatic renomination. If set to 0, it defaults to 3 seconds.
//
// Example:
//
//	agent, err := NewAgentWithOptions(
//		WithRenomination(DefaultNominationValueGenerator()),
//		WithAutomaticRenomination(3*time.Second),
//	)
func WithAutomaticRenomination(interval time.Duration) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// Note: renomination must be enabled separately via WithRenomination

// WithInterfaceFilter sets a filter function to whitelist or blacklist network interfaces
// for ICE candidate gathering.
//
// The filter function receives the interface name and should return true to keep the interface,
// or false to exclude it.
//
// Example:
//
//	// Only use interfaces starting with "eth"
//	agent, err := NewAgentWithOptions(
//		WithInterfaceFilter(func(interfaceName string) bool {
//			return len(interfaceName) >= 3 && interfaceName[:3] == "eth"
//		}),
//	)
func WithInterfaceFilter(filter func(string) bool) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}

// WithLoggerFactory sets the logger factory for the agent.
//
// Example:
//
//	import "github.com/pion/logging"
//
//	loggerFactory := logging.NewDefaultLoggerFactory()
//	loggerFactory.DefaultLogLevel = logging.LogLevelDebug
//	agent, err := NewAgentWithOptions(WithLoggerFactory(loggerFactory))
func WithLoggerFactory(loggerFactory logging.LoggerFactory) AgentOption {
	_ = "STUB: not implemented"
	return *new(AgentOption)
}
