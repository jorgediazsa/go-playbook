package ctxexercises

// Context: Context-Aware Retry Loops
// You are building an SDK that fetches user data from an unreliable 3rd-party API.
// Because it fails often, you want to retry up to 5 times, with a 1-second delay
// between retries.
//
// Why this matters: If the parent HTTP request was canceled by the user, or
// a global timeout was reached, your retry loop should NOT blindly sleep
// for another 1 second and try again. It should abort immediately!
//
// Requirements:
// 1. Refactor `FetchWithRetry` to respect `ctx.Done()`.
// 2. Instead of `time.Sleep(1 * time.Second)`, use a `select` statement that
//    waits for *either* the sleep timer to finish *or* the context to cancel.
// 3. If the context cancels during the "sleep", return `ctx.Err()` immediately.

import (
	"context"
	"errors"
	"time"
)

var ErrServerDown = errors.New("500 internal server error")

func UnreliableAPI() (string, error) {
	return "", ErrServerDown
}

func FetchWithRetry(ctx context.Context) (string, error) {
	// BUG: This loop sleeps and retries even if the context was already canceled!
	// TODO: Replace time.Sleep with a select on time.After and ctx.Done().

	for i := 0; i < 5; i++ {
		// Check context before even making the call
		if err := ctx.Err(); err != nil {
			return "", err
		}

		res, err := UnreliableAPI()
		if err == nil {
			return res, nil
		}

		// Wait 1 second before retrying
		time.Sleep(1 * time.Second)
	}

	return "", errors.New("max retries exceeded")
}
