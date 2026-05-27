// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"net"
)

// AddressRewriteMode controls whether a rule replaces or appends candidates.
type AddressRewriteMode int

const (
	addressRewriteModeUnspecified AddressRewriteMode = iota
	AddressRewriteReplace
	AddressRewriteAppend
)

// AddressRewriteRule represents a rule for remapping candidate addresses.
type AddressRewriteRule struct {
	// External are the 1:1 external addresses to advertise for this rule.
	// For replace mode, an empty list is treated as "drop the matched local
	// address" (no candidate emitted). For append mode, an empty list is a
	// no-op: the original candidate is kept.
	// Empty External rules are intentional:
	//   - Mode AddressRewriteReplace drops the matched candidate (deny-list style).
	//   - Mode AddressRewriteAppend keeps the original candidate and adds nothing,
	//     which is useful when you combine a catch-all replace with per-interface
	//     allow rules.
	External []string
	// Local optionally pins this rule to a specific local address. When set,
	// external IPs map to that address regardless of IP family. When empty,
	// External acts as a catch-all for the family implied by the local scope
	// (CIDR when set, otherwise the external IP family).
	Local string
	// Iface is the optional interface name to limit the rule to, empty = any.
	Iface string
	// CIDR is the optional CIDR to limit the rule to, empty = any.
	CIDR string
	// AsCandidateType is the candidate type to publish as for this rule. Defaults to host
	// when unspecified. Supported values: host, server reflexive, relay.
	AsCandidateType CandidateType
	// Mode controls whether we replace the original candidate or append extra
	// candidates.
	//
	// If Mode is zero, the default is:
	//   - CandidateTypeHost           -> AddressRewriteReplace
	//   - CandidateTypeServerReflexive, CandidateTypeRelay -> AddressRewriteAppend
	// For replace mode, a match with zero external IPs removes the candidate.
	// For append mode, a match with zero external IPs leaves the original
	// candidate untouched.
	Mode AddressRewriteMode
	// Networks is the optional networks to limit the rule to, nil/empty = all.
	Networks []NetworkType
}

func validateIPString(ipStr string) (net.IP, bool, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), false, nil
}

// ipMapping holds the mapping of local and external IP address
//
//	for a particular IP family.
type ipMapping struct {
	ipSole      []net.IP            // When non-empty, these are the catch-all external IPs for one local IP family
	ipMap       map[string][]net.IP // Local-to-external IP mapping (k: local, v: external IPs)
	valid       bool                // If not set any external IP, valid is false
	catchAllSet bool
}

func newIPMapping() ipMapping { _ = "STUB: not implemented"; return *new(ipMapping) }

func (m *ipMapping) addSoleIP(ip net.IP) { _ = "STUB: not implemented"; return }

func addExternalMappings(
	external []string,
	ruleMapping *addressRewriteRuleMapping,
	hasLocalAddr bool,
	localAddr net.IP,
	localIsIPv4 bool,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func maybeMarkEmptyMapping(
	ruleMapping *addressRewriteRuleMapping,
	added bool,
	hasLocalAddr bool,
	localIsIPv4 bool,
	localAddr net.IP,
) {
	_ = "STUB: not implemented"
	return
}

func (m *ipMapping) addIPMapping(locIP, extIP net.IP) { _ = "STUB: not implemented"; return }

func cloneIPs(src []net.IP) []net.IP { _ = "STUB: not implemented"; return nil }

func (m *ipMapping) findExternalIPs(locIP net.IP) []net.IP { _ = "STUB: not implemented"; return nil }

type addressRewriteRuleMapping struct {
	rule        AddressRewriteRule
	mode        AddressRewriteMode
	ipv4Mapping ipMapping
	ipv6Mapping ipMapping
	cidr        *net.IPNet
	allowIPv4   bool
	allowIPv6   bool
}

func (m *addressRewriteRuleMapping) hasMappings() bool { _ = "STUB: not implemented"; return false }

func (m *addressRewriteRuleMapping) mappingForFamily(isIPv4 bool) *ipMapping {
	_ = "STUB: not implemented"
	return nil
}

func (m *addressRewriteRuleMapping) isFamilyAllowed(isLocalIPv4 bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *addressRewriteRuleMapping) addImplicitMapping(
	extIP net.IP,
	isLocalIPv4 bool,
	hasLocalAddr bool,
	localAddr net.IP,
) {
	_ = "STUB: not implemented"
	return
}

type addressRewriteMapper struct {
	rulesByCandidateType map[CandidateType][]*addressRewriteRuleMapping
}

//nolint:gocognit,gocyclo,cyclop
func newAddressRewriteMapper(rules []AddressRewriteRule) (*addressRewriteMapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nilnil

//nolint:nilnil

func (m *addressRewriteMapper) hasCandidateType(candidateType CandidateType) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *addressRewriteMapper) shouldReplace(candidateType CandidateType) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *addressRewriteMapper) findExternalIPs(
	candidateType CandidateType,
	localIPStr string,
	iface string,
) ([]net.IP, bool, AddressRewriteMode, error) {
	_ = "STUB: not implemented"
	return nil, false, *new(AddressRewriteMode), nil
}

func ruleMappingForLookup(
	rule *addressRewriteRuleMapping,
	locIP net.IP,
	isLocIPv4 bool,
	iface string,
) (*ipMapping, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func catchAllSpecificity(rule *addressRewriteRuleMapping, iface string) int {
	_ = "STUB: not implemented"
	return 0
}

func evaluateRewriteRules(
	rules []*addressRewriteRuleMapping,
	locIP net.IP,
	isLocIPv4 bool,
	iface string,
) (ips []net.IP, matched bool, mode AddressRewriteMode) {
	_ = "STUB: not implemented"
	return nil, false, *new(AddressRewriteMode)
}
