package advancedconcurrency

import (
	"errors"
	"sync"
	"time"
)

// Context: Semaphore Rate Limiting
// You have an API endpoint that queries an old, fragile Mainframe.
// The Mainframe crashes if it receives more than 3 concurrent connections,
// regardless of how many users are hitting your Go API.
//
// Why this matters: You could use a worker pool, but dynamic web handlers
// are usually just goroutines spun up by `net/http`. You can't put them in a pool.
// Instead, you use a Semaphore (a bounded token bucket) to pause incoming
// requests until a slot opens up.
//
// Requirements:
// 1. Refactor `MainframeGateway` to use a buffered channel of size 3 as a Semaphore.
// 2. In `Query`, the request MUST block by attempting to send a "token" (an empty struct)
//    to the channel BEFORE doing the actual work.
// 3. To prevent hanging a user forever if the Mainframe locks up, wrap the token
//    acquisition in a `select` with a `time.After(100 * time.Millisecond)` timeout.
//    If the timeout triggers, return `ErrRateLimited`.
// 4. If the token is acquired, do the work, and then REMOVE the token from the
//    channel (using `defer <-semaphore`) before returning the result.

var ErrRateLimited = errors.New("rate limited: no available connections")

type MainframeGateway struct {
	// TODO: Define a buffered channel of empty structs to act as a semaphore.
	mu sync.Mutex // Currently just using a basic mutex natively, which limits to 1!
}

func NewMainframeGateway(maxConns int) *MainframeGateway {
	// TODO: Initialize your semaphore channel
	return &MainframeGateway{}
}

func (g *MainframeGateway) Query(data string) (string, error) {
	// BUG: This uses a Mutex, restricting concurrency to exactly 1.
	// It's safe, but unnecessarily slow. It acts like a queue of infinite length.
	// TODO: Replace with the Semaphore token acquisition via `select`
	// for bounds + timeout protection!

	g.mu.Lock()
	defer g.mu.Unlock()

	// Simulate work
	time.Sleep(50 * time.Millisecond)

	return "ACK: " + data, nil
}
