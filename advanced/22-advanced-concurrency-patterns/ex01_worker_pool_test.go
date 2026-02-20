package advancedconcurrency

import (
	"context"
	"errors"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestProcessBatch(t *testing.T) {
	// Test 1: Happy Path Concurrency
	t.Run("Concurrent Execution", func(t *testing.T) {
		jobs := make([]int, 50)
		for i := range jobs {
			jobs[i] = i
		}

		var processed int32

		start := time.Now()
		err := ProcessBatch(context.Background(), jobs, 10, func(ctx context.Context, job int) error {
			time.Sleep(10 * time.Millisecond) // Simulate work
			atomic.AddInt32(&processed, 1)
			return nil
		})
		duration := time.Since(start)

		if err != nil {
			t.Fatalf("Expected nil error, got %v", err)
		}
		if processed != 50 {
			t.Fatalf("Expected 50 jobs processed, got %d", processed)
		}

		// If sequential, 50 * 10ms = 500ms.
		// If concurrent (10 workers), 5 * 10ms = ~50ms
		if duration > 150*time.Millisecond {
			t.Fatalf("FAILED: Pool ran too slowly (%v). It's likely executing sequentially instead of concurrently.", duration)
		}
	})

	// Test 2: Error Propagation & Cancellation
	t.Run("Cancellation on Error", func(t *testing.T) {
		baselineGoroutines := runtime.NumGoroutine()

		jobs := make([]int, 100)
		for i := range jobs {
			jobs[i] = i
		}

		ErrDiskFull := errors.New("disk full")
		var processed int32

		// We use a custom processFunc that fails on job 5.
		// If cancellation works, the pool should abort almost immediately.
		err := ProcessBatch(context.Background(), jobs, 5, func(ctx context.Context, job int) error {
			// Respect cancellation
			if err := ctx.Err(); err != nil {
				return err
			}

			if job == 5 {
				return ErrDiskFull
			}

			time.Sleep(20 * time.Millisecond)
			atomic.AddInt32(&processed, 1)
			return nil
		})

		if !errors.Is(err, ErrDiskFull) {
			t.Fatalf("FAILED: Expected ErrDiskFull to be propagated, got %v", err)
		}

		// Wait briefly for lingering routines to die
		time.Sleep(100 * time.Millisecond)

		if runtime.NumGoroutine() > baselineGoroutines+2 {
			t.Fatalf("FAILED: Goroutine leak detected! active: %d, baseline: %d. Workers did not exit when the error occurred.", runtime.NumGoroutine(), baselineGoroutines)
		}

		// If processed is close to 100, cancellation didn't work.
		if processed > 20 {
			t.Fatalf("FAILED: Processed %d jobs after an error. The context cancellation was ignored by the other workers!", processed)
		}
	})
}
