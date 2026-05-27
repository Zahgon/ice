// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package taskloop implements a task loop to run
// tasks sequentially in a separate Goroutine.
package taskloop

import (
	"context"
	"errors"
	"time"

	atomicx "github.com/pion/ice/v4/internal/atomic"
)

// ErrClosed indicates that the loop has been stopped.
var ErrClosed = errors.New("the agent is closed")

type task struct {
	fn   func(context.Context)
	done chan struct{}
}

// Loop runs submitted task serially in a dedicated Goroutine.
type Loop struct {
	tasks chan task

	// State for closing
	done         chan struct{}
	taskLoopDone chan struct{}
	err          atomicx.Error
}

// New creates and starts a new task loop.
func New(onClose func()) *Loop { _ = "STUB: not implemented"; return nil }

// runLoop handles registered tasks and agent close.
func (l *Loop) runLoop(onClose func()) { _ = "STUB: not implemented"; return }

// Close stops the loop after finishing the execution of the current task.
// Other pending tasks will not be executed.
func (l *Loop) Close() { _ = "STUB: not implemented"; return }

// Run serially executes the submitted callback.
// Blocking tasks must be cancelable by context.
func (l *Loop) Run(ctx context.Context, t func(context.Context)) error {
	_ = "STUB: not implemented"
	return nil
}

// The following methods implement context.Context for TaskLoop

// Done returns a channel that's closed when the task loop has been stopped.
func (l *Loop) Done() <-chan struct{} {
	_ = "STUB: not implemented"

	// Err returns nil if the task loop is still running.
	// Otherwise it return errClosed if the loop has been closed/stopped.
	return nil
}

func (l *Loop) Err() error { _ = "STUB: not implemented"; return nil }

// Deadline returns the no valid time as task loops have no deadline.
func (l *Loop) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// Value is not supported for task loops.
func (l *Loop) Value(any) any { _ = "STUB: not implemented"; return *new(any) }
