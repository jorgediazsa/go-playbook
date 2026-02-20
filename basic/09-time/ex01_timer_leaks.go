package timeutil

import (
	"errors"
	"time"
)

// Context: Timer Leaks in select blocks
// You are building an HTTP client wrapper that fetches data from a third-party API.
// To protect your system, you enforce a strict 30-second timeout on the request.
//
// Why this matters: The third-party API usually responds in 50ms.
// If you use `time.After(30 * time.Second)` in a high-throughput endpoint,
// you leak a timer struct for 30 seconds on EVERY SINGLE REQUEST. At 1,000 Req/Sec,
// you are holding 30,000 dead timers in memory permanently.
//
// Requirements:
// 1. Refactor `FetchWithTimeout` to avoid leaking the 30-second timer
//    when the fast path `respCh` returns quickly.
// 2. You must use `time.NewTimer` and explicitly `Stop()` it.

var ErrTimeout = errors.New("request timed out")

func FetchWithTimeout(respCh <-chan string) (string, error) {
	// BUG: Large timer leak in the fast-path!
	// TODO: Replace time.After with a properly managed time.NewTimer.
	// Ensure the timer is stopped when the function exits so the GC can clean it up.

	timeout := 30 * time.Second

	select {
	case resp := <-respCh:
		// The 30s timer created by time.After still sits in memory!
		return resp, nil
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}
