package observability

import (
	"context"
	"io"
	"runtime/trace"
	"sync"
	"time"
)

// Context: Execution Tracing
// You are trying to figure out why your 8-core CPU is only running at 15%
// utilization when processing batches of 10,000 requests.
//
// Why this matters: The CPU profile doesn't show any functions burning CPU—it
// just looks "idle". An execution trace (`runtime/trace`) will show you EXACTLY
// why goroutines are sleeping (e.g., waiting on a Mutex, waiting on a Channel,
// or blocking on Syscalls).
//
// Requirements:
// 1. Run this via the test `go test -run TestGenerateTrace`. It will generate
//    a file called `trace.out`.
// 2. Open it: `go tool trace trace.out`.
// 3. Notice the massive "waiting" blocks caused by the global Mutex below.
// 4. Refactor `BatchProcessor` to NOT use a global Mutex around the mock "work".
//    Let the 8 goroutines run truly concurrently!

type BatchProcessor struct {
	mu sync.Mutex
}

func (b *BatchProcessor) Process(ctx context.Context, input []int) {
	// We want to record this exact function span in the trace viewer!
	defer trace.StartRegion(ctx, "BatchProcessRegion").End()

	var wg sync.WaitGroup
	for _, val := range input {
		wg.Add(1)

		go func(v int) {
			defer wg.Done()

			// BUG: This locks out all other goroutines. They will all pile up
			// in the "Waiting" state in the trace viewer!
			// TODO: Remove this lock to allow concurrency!
			b.mu.Lock()
			defer b.mu.Unlock()

			// Simulate doing CPU work
			time.Sleep(1 * time.Millisecond)
		}(val)
	}
	wg.Wait()
}

// Ignore for exercise conceptually, used for test.
func CaptureTrace(w io.Writer, work func()) error {
	if err := trace.Start(w); err != nil {
		return err
	}
	defer trace.Stop()
	work()
	return nil
}
