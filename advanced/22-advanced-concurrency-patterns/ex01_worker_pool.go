package advancedconcurrency

import (
	"context"
)

// Context: Resilient Worker Pool
// You are building a batch processing system that processes thousands of
// uploaded images.
//
// Why this matters: You must never spawn 10,000 goroutines at once. You must
// use a bounded pool. Furthermore, if ONE image processing fails critically
// (e.g. disk full), the ENTIRE batch must abort immediately, and NO goroutines
// should leak.
//
// Requirements:
// 1. Refactor `ProcessBatch` to spawn exactly `concurrency` worker goroutines.
// 2. The workers should pull jobs from the `jobs` channel until it's closed.
// 3. If a worker encounters an error from `processFunc`, it should:
//    a. Cancel the context for all other workers.
//    b. Record the FIRST error encountered so `ProcessBatch` can return it.
// 4. Ensure all workers exit cleanly before `ProcessBatch` returns (Wait!).
// 5. Do not leak goroutines. Do not panic on closed channels.

func ProcessBatch(ctx context.Context, jobs []int, concurrency int, processFunc func(ctx context.Context, job int) error) error {
	// BUG: This current implementation processes sequentially.
	// And if it fails, it just returns, but in a concurrent system,
	// returning early leaves background workers orphaned (leaked)!

	// TODO: Create a derived context with cancellation so workers can abort each other.
	// TODO: Spawn `concurrency` workers.
	// TODO: Feed the `jobs` slice into a channel. Don't forget to close it!
	// TODO: Wait for all workers to finish. Return the first error if any.

	for _, j := range jobs {
		if err := ctx.Err(); err != nil {
			return err
		}

		if err := processFunc(ctx, j); err != nil {
			return err
		}
	}

	return nil
}
