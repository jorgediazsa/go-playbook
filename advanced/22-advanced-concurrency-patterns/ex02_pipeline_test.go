package advancedconcurrency

import (
	"context"
	"runtime"
	"testing"
	"time"
)

func TestPipelineBackpressure(t *testing.T) {
	// The problem with testing a "build your own pipeline" structurally is that
	// the student has total freedom over the internal stage functions.
	// We verify the outer contract: Concurrency + Leak Prevention.

	t.Run("Happy Path Sum", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		sum, err := buildPipeline(context.Background(), nums)
		if err != nil {
			t.Fatalf("Expected nil error, got %v", err)
		}

		expected := (1 + 2 + 3 + 4 + 5) * 2
		if sum != expected {
			t.Fatalf("Expected sum %d, got %d", expected, sum)
		}
	})

	t.Run("Cancellation cascading on upstream stages", func(t *testing.T) {
		baselineGoroutines := runtime.NumGoroutine()

		nums := make([]int, 1000)

		// If the user's pipeline correctly uses `select { case ch <- val: case <-ctx.Done(): ... }`,
		// when we forcibly cancel the context from the outside, the internal generator and
		// processor should instantly abort and exit, avoiding a leak.

		ctx, cancel := context.WithCancel(context.Background())

		// We launch the pipeline in the background so we can cancel it midway.
		done := make(chan struct{})
		go func() {
			_, _ = buildPipeline(ctx, nums)
			close(done)
		}()

		time.Sleep(10 * time.Millisecond)
		cancel() // FIRE THE CANCELLATION!

		select {
		case <-done:
			// Good, it exited.
		case <-time.After(1 * time.Second):
			t.Fatalf("FAILED: The pipeline did not exit when the context was canceled! Your stages are likely deadlocked trying to send to channels.")
		}

		time.Sleep(10 * time.Millisecond) // buffer for goroutines to clean up

		if runtime.NumGoroutine() > baselineGoroutines+2 {
			t.Fatalf("FAILED: Goroutine leak detected! active: %d, baseline: %d.", runtime.NumGoroutine(), baselineGoroutines)
		}
	})
}
