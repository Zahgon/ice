// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type candidateBase struct {
	id            string
	networkType   NetworkType
	candidateType CandidateType

	component      uint16
	address        string
	port           int
	relatedAddress *CandidateRelatedAddress
	tcpType        TCPType

	resolvedAddr net.Addr

	lastSent     atomic.Int64
	lastReceived atomic.Int64
	conn         net.PacketConn

	currAgent *Agent
	closeCh   chan struct{}
	closedCh  chan struct{}

	foundationOverride string
	priorityOverride   uint32

	relayLocalPreference uint16

	remoteCandidateCaches sync.Map // map[AddrPort]Candidate
	isLocationTracked     bool
	extensions            []CandidateExtension
}

// Save a time reference to calculate monotonic time for candidate last sent/received.
// nolint: gochecknoglobals
var timeRef = time.Now()

// getMonoNanos returns the monotonic nanoseconds of a time t since timeRef.
func getMonoNanos(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

// getMonoTime returns a time.Time based on monotonic nanos since timeRef.
func getMonoTime(nanos int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Done implements context.Context.
func (c *candidateBase) Done() <-chan struct{} {
	_ = "STUB: not implemented"

	// Err implements context.Context.
	return nil
}

func (c *candidateBase) Err() error { _ = "STUB: not implemented"; return nil }

// Deadline implements context.Context.
func (c *candidateBase) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// Value implements context.Context.
func (c *candidateBase) Value(any) any {
	_ = "STUB: not implemented"

	// setWriteDeadline is used by upper layers to push write deadlines down to the
	// underlying packet connection.
	return *new(any)
}

func (c *candidateBase) setWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// ID returns Candidate ID.
func (c *candidateBase) ID() string { _ = "STUB: not implemented"; return "" }

func (c *candidateBase) Foundation() string { _ = "STUB: not implemented"; return "" }

// Address returns Candidate Address.
func (c *candidateBase) Address() string {
	_ = "STUB: not implemented"

	// Port returns Candidate Port.
	return ""
}

func (c *candidateBase) Port() int {
	_ = "STUB: not implemented"

	// Type returns candidate type.
	return 0
}

func (c *candidateBase) Type() CandidateType {
	_ = "STUB: not implemented"
	return *

	// NetworkType returns candidate NetworkType.
	new(CandidateType)
}

func (c *candidateBase) NetworkType() NetworkType {
	_ = "STUB: not implemented"
	return *

	// Component returns candidate component.
	new(NetworkType)
}

func (c *candidateBase) Component() uint16 { _ = "STUB: not implemented"; return 0 }

func (c *candidateBase) SetComponent(component uint16) { _ = "STUB: not implemented"; return }

// LocalPreference returns the local preference for this candidate.
func (c *candidateBase) LocalPreference() uint16 {
	_ = "STUB: not implemented" //nolint:cyclop
	return 0
}

// RFC 6544, section 4.2
//
// In Section 4.1.2.1 of [RFC5245], a recommended formula for UDP ICE
// candidate prioritization is defined.  For TCP candidates, the same
// formula and candidate type preferences SHOULD be used, and the
// RECOMMENDED type preferences for the new candidate types defined in
// this document (see Section 5) are 105 for NAT-assisted candidates and
// 75 for UDP-tunneled candidates.
//
// (...)
//
// With TCP candidates, the local preference part of the recommended
// priority formula is updated to also include the directionality
// (active, passive, or simultaneous-open) of the TCP connection.  The
// RECOMMENDED local preference is then defined as:
//
//     local preference = (2^13) * direction-pref + other-pref
//
// The direction-pref MUST be between 0 and 7 (both inclusive), with 7
// being the most preferred.  The other-pref MUST be between 0 and 8191
// (both inclusive), with 8191 being the most preferred.  It is
// RECOMMENDED that the host, UDP-tunneled, and relayed TCP candidates
// have the direction-pref assigned as follows: 6 for active, 4 for
// passive, and 2 for S-O.  For the NAT-assisted and server reflexive
// candidates, the RECOMMENDED values are: 6 for S-O, 4 for active, and
// 2 for passive.
//
// (...)
//
// If any two candidates have the same type-preference and direction-
// pref, they MUST have a unique other-pref.  With this specification,
// this usually only happens with multi-homed hosts, in which case
// other-pref is the preference for the particular IP address from which
// the candidate was obtained.  When there is only a single IP address,
// this value SHOULD be set to the maximum allowed value (8191).

// RelatedAddress returns *CandidateRelatedAddress.
func (c *candidateBase) RelatedAddress() *CandidateRelatedAddress {
	_ = "STUB: not implemented"
	return nil
}

func (c *candidateBase) TCPType() TCPType {
	_ = "STUB: not implemented"

	// start runs the candidate using the provided connection.
	return *new(TCPType)
}

func (c *candidateBase) start(a *Agent, conn net.PacketConn, initializedCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

var bufferPool = sync.Pool{ // nolint:gochecknoglobals
	New: func() any {
		buf := make([]byte, receiveMTU)

		return &buf
	},
}

func (c *candidateBase) recvLoop(initializedCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *candidateBase) validateSTUNTrafficCache(addr net.Addr) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *candidateBase) addRemoteCandidateCache(candidate Candidate, srcAddr net.Addr) {
	_ = "STUB: not implemented"
	return
}

func (c *candidateBase) replaceRemoteCandidateCacheValues(oldRemote, newRemote Candidate) {
	_ = "STUB: not implemented"
	return
}

func (c *candidateBase) handleInboundPacket(buf []byte, srcAddr net.Addr) {
	_ = "STUB: not implemented"
	return
}

// Explicitly copy raw buffer so Message can own the memory.

// nolint: contextcheck

//nolint:contextcheck

// Note: This will return packetio.ErrFull if the buffer ever manages to fill up.

// Add received application bytes to the currently selected candidate pair.

// close stops the recvLoop.
func (c *candidateBase) close() error {
	_ = "STUB: not implemented"
	// If conn has never been started will be nil
	return nil
}

// Assert that conn has not already been closed

// Unblock recvLoop

// Close the conn

// Wait until the recvLoop is closed

func (c *candidateBase) writeTo(raw []byte, dst Candidate) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If the connection is closed, we should return the error

// TypePreference returns the type preference for this candidate.
func (c *candidateBase) TypePreference() uint16 { _ = "STUB: not implemented"; return 0 }

// Priority computes the priority for this ICE Candidate
// See: https://www.rfc-editor.org/rfc/rfc8445#section-5.1.2.1
func (c *candidateBase) Priority() uint32 { _ = "STUB: not implemented"; return 0 }

// The local preference MUST be an integer from 0 (lowest preference) to
// 65535 (highest preference) inclusive.  When there is only a single IP
// address, this value SHOULD be set to 65535.  If there are multiple
// candidates for a particular component for a particular data stream
// that have the same type, the local preference MUST be unique for each
// one.

// transportAddressEqual checks if the transport address (IP, Port, NetworkType, TCPType) is equal to another
// candidate.
func (c *candidateBase) transportAddressEqual(other Candidate) bool {
	_ = "STUB: not implemented"
	return false
}

// Equal is used to compare two candidateBases.
func (c *candidateBase) Equal(other Candidate) bool { _ = "STUB: not implemented"; return false }

// DeepEqual is same as Equal but also compares the extensions.
func (c *candidateBase) DeepEqual(other Candidate) bool { _ = "STUB: not implemented"; return false }

// String makes the candidateBase printable.
func (c *candidateBase) String() string { _ = "STUB: not implemented"; return "" }

// LastReceived returns a time.Time indicating the last time
// this candidate was received.
func (c *candidateBase) LastReceived() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *candidateBase) setLastReceived(t time.Time) { _ = "STUB: not implemented"; return }

// LastSent returns a time.Time indicating the last time
// this candidate was sent.
func (c *candidateBase) LastSent() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *candidateBase) setLastSent(t time.Time) { _ = "STUB: not implemented"; return }

func (c *candidateBase) seen(outbound bool) { _ = "STUB: not implemented"; return }

func (c *candidateBase) addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *candidateBase) filterForLocationTracking() bool { _ = "STUB: not implemented"; return false }

func (c *candidateBase) agent() *Agent { _ = "STUB: not implemented"; return nil }

func (c *candidateBase) context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *candidateBase) copy() (Candidate, error) {
	_ = "STUB: not implemented"
	return *new(Candidate), nil
}

func removeZoneIDFromAddress(addr string) string { _ = "STUB: not implemented"; return "" }

// Marshal returns the string representation of the ICECandidate.
func (c *candidateBase) Marshal() string { _ = "STUB: not implemented"; return "" }

// CandidateExtension represents a single candidate extension
// as defined in https://tools.ietf.org/html/rfc5245#section-15.1
// .
type CandidateExtension struct {
	Key   string
	Value string
}

func (c *candidateBase) Extensions() []CandidateExtension { _ = "STUB: not implemented"; return nil }

// We store the TCPType in c.tcpType, but we need to return it as an extension.

// Get returns the value of the given key if it exists.
func (c *candidateBase) GetExtension(key string) (CandidateExtension, bool) {
	_ = "STUB: not implemented"
	return *new(CandidateExtension), false
}

// TCPType was manually set.
//nolint:goconst

func (c *candidateBase) AddExtension(ext CandidateExtension) error {
	_ = "STUB: not implemented"
	return nil
}

// per spec, Extensions aren't explicitly unique, we only set the first one.
// If the exteion is set multiple times.

func (c *candidateBase) RemoveExtension(key string) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// marshalExtensions returns the string representation of the candidate extensions.
func (c *candidateBase) marshalExtensions() string { _ = "STUB: not implemented"; return "" }

// Equal returns true if the candidate extensions are equal.
func (c *candidateBase) extensionsEqual(other []CandidateExtension) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *candidateBase) setExtensions(extensions []CandidateExtension) {
	_ = "STUB: not implemented"
	return
}

// UnmarshalCandidate Parses a candidate from a string
// https://datatracker.ietf.org/doc/html/rfc5245#section-15.1
func UnmarshalCandidate(raw string) (Candidate, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	// Handle candidates with the "candidate:" prefix as defined in RFC 5245 section 15.1.
	return *new(Candidate), nil
}

// foundation ( 1*32ice-char ) But we allow for empty foundation,

//nolint:errorlint // we wrap the error

// Empty foundation, not RFC 8445 compliant but seen in the wild

// component-id ( 1*5DIGIT )

//nolint:errorlint // we wrap the error

// transport ( "UDP" / transport-extension ; from RFC 3261 ) SP

// priority ( 1*10DIGIT ) SP

//nolint:errorlint // we wrap the error

// connection-address SP     ;from RFC 4566

// Remove IPv6 ZoneID: https://github.com/pion/ice/pull/704

// port from RFC 4566

//nolint:errorlint // we wrap the error

// "typ" SP

// SP cand-type ("host" / "srflx" / "prflx" / "relay")

//nolint:errorlint // we wrap the error

// this code is ugly because we can't break backwards compatibility
// with the old way of parsing candidates

//nolint:gosec // G115 no overflow we read 5 digits
//nolint:gosec // G115 no overflow we read 5 digits

//nolint:gosec // G115 no overflow we read 5 digits
//nolint:gosec // G115 no overflow we read 5 digits

//nolint:gosec // G115 no overflow we read 5 digits
//nolint:gosec // G115 no overflow we read 5 digits

//nolint:gosec // G115 no overflow we read 5 digits
//nolint:gosec // G115 no overflow we read 5 digits

// Read an ice-char token from the raw string
// ice-char = ALPHA / DIGIT / "+" / "/"
// stop reading when a space is encountered or the end of the string.
func readCandidateCharToken(raw string, start int, limit int) (string, int, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return "", 0, nil
}

// SP

//nolint: err113 // handled by caller

//nolint: err113 // handled by caller

// Read an ice string token from the raw string until a space is encountered
// Or the end of the string, we imply that ice string are UTF-8 encoded.
func readCandidateStringToken(raw string, start int) (string, int) {
	_ = "STUB: not implemented"
	return "", 0
}

// SP

// Read a digit token from the raw string
// stop reading when a space is encountered or the end of the string.
func readCandidateDigitToken(raw string, start, limit int) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// SP

//nolint: err113 // handled by caller

//nolint: err113 // handled by caller

// Read and validate RFC 4566 port from the raw string.
func readCandidatePort(raw string, start int) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

//nolint: err113 // handled by caller

// Read a byte-string token from the raw string
// As defined in RFC 4566  1*(%x01-09/%x0B-0C/%x0E-FF) ;any byte except NUL, CR, or LF
// we imply that extensions byte-string are UTF-8 encoded.
func readCandidateByteString(raw string, start int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// SP

// 1*(%x01-09/%x0B-0C/%x0E-FF)

//nolint: err113 // handled by caller

// Read and validate raddr and rport from the raw string
// [SP rel-addr] [SP rel-port]
// defined in https://datatracker.ietf.org/doc/html/rfc5245#section-15.1
// .
func tryReadRelativeAddrs(raw string, start int) (raddr string, rport, pos int, err error) {
	_ = "STUB: not implemented"
	return "", 0, 0, nil
}

//nolint:errorlint // we wrap the error

// UnmarshalCandidateExtensions parses the candidate extensions from the raw string.
// *(SP extension-att-name SP extension-att-value)
// Where extension-att-name, and extension-att-value are byte-strings
// as defined in https://tools.ietf.org/html/rfc5245#section-15.1
func unmarshalCandidateExtensions(raw string) (extensions []CandidateExtension, rawTCPTypeRaw string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// SP

//nolint: errorlint // we wrap the error

// while not spec-compliant, we allow for empty values, as seen in the wild

//nolint: errorlint // we are wrapping the error
