// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package stun contains ICE specific STUN code
package stun

import (
	"errors"
	"net"
	"time"

	"github.com/pion/stun/v3"
)

var (
	errGetXorMappedAddrResponse = errors.New("failed to get XOR-MAPPED-ADDRESS response")
	errMismatchUsername         = errors.New("username mismatch")
)

// GetXORMappedAddr initiates a STUN requests to serverAddr using conn, reads the response and returns
// the XORMappedAddress returned by the STUN server.
func GetXORMappedAddr(conn net.PacketConn, serverAddr net.Addr, timeout time.Duration) (*stun.XORMappedAddress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset timeout after completion
//nolint:errcheck

//nolint:errorlint

// AssertUsername checks that the given STUN message m has a USERNAME attribute with a given value.
func AssertUsername(m *stun.Message, expectedUsername string) error {
	_ = "STUB: not implemented"
	return nil
}
