package memorymodel

import (
	"sync"
	"sync/atomic"
)

// Context: False Sharing (CPU Cache Line Contention)
// You have a high-performance metrics aggregator counting API hits.
// You have 8 worker goroutines, but you give them separate independent
// counters so they don't block each other on a Mutex.
//
// Why this matters: The CPU cache pushes data around in 64-byte chunks (Cache Lines).
// Since `Counter1` and `Counter2` are consecutive `int64`s (8 bytes each),
// they fit into the SAME 64-byte CPU cache line.
// Even though Worker 1 only writes to `Counter1` and Worker 2 only writes to `Counter2`,
// the CPU hardware falsely registers a memory collision and constantly invalidates
// and bounces the exact same cache line back and forth between core L1 caches.
// This plummets multithreaded performance to worse than single-threaded locking!
//
// Requirements:
// 1. Run `go test -bench BenchmarkFalseSharing` to see the baseline speed.
// 2. Add struct padding to `PaddedMetrics`.
//    (Hint: Add `_ [56]byte` between the counters so they exceed the 64-byte cache line).
// 3. Rerun the benchmark to witness the massive throughput explosion when cores
//    no longer falsely share cache lines.

type NaiveMetrics struct {
	Counter1 int64
	Counter2 int64
}

// TODO: Fix false sharing by ensuring these counters are physically separated
// in RAM by at least 64 bytes (the size of a standard CPU cache line).
type PaddedMetrics struct {
	// BUG: These share the same cache line.
	// Insert padding here: `_ [56]byte` (64 bytes - 8 bytes for int64)

	Counter1 int64
	Counter2 int64
}

// These functions are provided to test the benchmarks.
func UpdateNaive(m *NaiveMetrics, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10_000_000; i++ {
		atomic.AddInt64(&m.Counter1, 1)
	}
}

func UpdatePadded(m *PaddedMetrics, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10_000_000; i++ {
		atomic.AddInt64(&m.Counter1, 1)
	}
}
