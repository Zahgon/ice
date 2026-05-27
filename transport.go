// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ice

import (
	"context"
	"net"
	"sync/atomic"
	"time"
)

// AwaitConnect waits until a pair is selected.
func (a *Agent) AwaitConnect(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// StartDial sets the agent up for connecting to the remote agent, acting as the
// controlling agent and returns immediately.
func (a *Agent) StartDial(remoteUfrag, remotePwd string) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Dial blocks until at least one ice candidate pair has successfully connected.
func (a *Agent) Dial(ctx context.Context, remoteUfrag, remotePwd string) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:contextcheck

// StartAccept sets the agent up for connecting to the remote agent, acting as the
// controlled agent and returns immediately.
func (a *Agent) StartAccept(remoteUfrag, remotePwd string) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Accept blocks until at least one ice candidate pair has successfully connected.
func (a *Agent) Accept(ctx context.Context, remoteUfrag, remotePwd string) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:contextcheck

// Conn represents the ICE connection.
// At the moment the lifetime of the Conn is equal to the Agent.
type Conn struct {
	bytesReceived atomic.Uint64
	bytesSent     atomic.Uint64
	agent         *Agent
}

// BytesSent returns the number of bytes sent.
func (c *Conn) BytesSent() uint64 { _ = "STUB: not implemented"; return 0 }

// BytesReceived returns the number of bytes received.
func (c *Conn) BytesReceived() uint64 { _ = "STUB: not implemented"; return 0 }

func (a *Agent) startConnect(isControlling bool, remoteUfrag, remotePwd string) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:contextcheck

// Read implements the Conn Read method.
func (c *Conn) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec // G115

// Write implements the Conn Write method.
func (c *Conn) Write(packet []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Write application data via the selected pair and update stats with actual bytes written.

// GetCandidatePairsInfo returns snapshot information for all candidate pairs.
// Use the returned ID with WriteToPair() to write to a specific pair.
func (c *Conn) GetCandidatePairsInfo() []CandidatePairInfo { _ = "STUB: not implemented"; return nil }

// WriteToPair writes packet to a specific candidate pair identified by its ID.
// Returns ErrCandidatePairNotFound if the pair ID is not found.
// Returns ErrCandidatePairNotSucceeded if the pair is not in Succeeded state.
// This is useful for sending packets over alternate paths
// even if they are not nominated.
func (c *Conn) WriteToPair(pairID uint64, packet []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close implements the Conn Close method. It is used to close
// the connection. Any calls to Read and Write will be unblocked and return an error.
func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

// LocalAddr returns the local address of the current selected pair or nil if there is none.
func (c *Conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// RemoteAddr returns the remote address of the current selected pair or nil if there is none.
func (c *Conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// SetDeadline sets both read and write deadlines on the underlying ICE connection.
func (c *Conn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadDeadline sets the read deadline on the packet buffer used for application data.
func (c *Conn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetWriteDeadline sets the write deadline on the currently selected local candidate connection.
// The deadline applies to the selected candidate pair and will affect all traffic over that pair.
func (c *Conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
