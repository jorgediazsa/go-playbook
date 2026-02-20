package advancedconcurrency

import (
	"context"
)

// Context: Pipelining & Backpressure
// You are building an ETL pipeline.
// Stage 1 (Generator): Reads rows from a CSV.
// Stage 2 (Processor): Performs heavy math on each row.
// Stage 3 (Sink): Writes the results to a Database.
//
// Why this matters: If the Database Sink is slow, the Processor's output channel
// fills up, blocking the Processor. This blocks the Generator. This is Backpressure.
// But what happens if the Sink Encounters a fatal DB error halfway through?
// If the Sink just returns early, the Generator and Processor are permanently stuck
// trying to write to unwatched channels (goroutine leaks!).
//
// Requirements:
// 1. Refactor buildPipeline. Have the main function wait for the Sink to finish.
// 2. The pipeline MUST use a cancellation context (`context.WithCancel`).
// 3. If the Sink encounters an error, it MUST cancel the context.
// 4. The Generator and Processor MUST select on `ctx.Done()` alongside sending
//    to their output channels. If `ctx.Done()` triggers, they must exit immediately,
//    safely closing their output channels on the way out.

func buildPipeline(ctx context.Context, nums []int) (int, error) {
	// 1. We create the context needed for cancellation
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	// Channels
	// genCh := make(chan int)
	// procCh := make(chan int)

	// BUG: The current implementation (below) is fully sequential, uses no channels,
	// and has no backpressure or concurrent stages.

	sum := 0
	for _, n := range nums {
		// Stage 1: Generate
		v1 := n

		// Stage 2: Process (multiply by 2)
		v2 := v1 * 2

		// Stage 3: Sink (sum them up)
		sum += v2
	}

	return sum, nil
}
