package observability

import (
	"context"
	"os"
	"testing"
	"time"
)

// This test mostly serves as a harness to generate the trace.out file for the student to look at.
func TestGenerateTrace(t *testing.T) {
	f, err := os.Create("trace.out")
	if err != nil {
		t.Fatalf("Could not create trace file: %v", err)
	}
	defer f.Close()

	processor := &BatchProcessor{}

	// Huge input
	input := make([]int, 200)

	start := time.Now()

	err = CaptureTrace(f, func() {
		ctx := context.Background()
		processor.Process(ctx, input)
	})

	if err != nil {
		t.Fatalf("Trace failed: %v", err)
	}

	duration := time.Since(start)

	// With the Mutex: 200 items * 1ms sequentially = > 200ms duration.
	// Without the Mutex: 200 items spread across ~8-10 cores concurrently = ~20-30ms.

	if duration > 100*time.Millisecond {
		t.Fatalf("FAILED: Processed took %v. That's too slow! Run `go tool trace trace.out` and observe the Mutex contention. Then fix it!", duration)
	}

	t.Logf("Success! Processed securely in %v. Trace saved to trace.out", duration)
}
