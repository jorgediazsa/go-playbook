package ctxexercises

import (
	"context"
	"testing"
	"time"
)

func TestFetchWithRetryCancellation(t *testing.T) {
	// We give the entire fetch operation a maximum budget of 1500ms.
	// Since each retry sleeps for 1000ms, it should fail during the SECOND sleep.

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := FetchWithRetry(ctx)
	duration := time.Since(start)

	// If they used `time.Sleep(1 * time.Second)`, the loop will run for a full
	// 5 seconds before returning "max retries exceeded" because time.Sleep blocks
	// the thread completely, ignoring the context!

	// If they used `select { case <-time.After(): case <-ctx.Done(): ... }`,
	// it will instantly abort at ~1500ms.

	if err != context.DeadlineExceeded {
		t.Fatalf("FAILED: Expected context.DeadlineExceeded, got %v", err)
	}

	if duration > 2000*time.Millisecond {
		t.Fatalf("FAILED: Retry loop ignored the context cancellation during sleep! It took %v", duration)
	}
}
