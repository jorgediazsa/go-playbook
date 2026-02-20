package memorymodel

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Run with: `go test -bench BenchmarkFalseSharing`
func BenchmarkFalseSharing(b *testing.B) {
	b.Run("Naive (Shared Cache Line)", func(b *testing.B) {
		m := &NaiveMetrics{}
		var wg sync.WaitGroup

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Add(2)
			// Both coroutines slam the exact same memory cache line structure simultaneously
			go func() {
				defer wg.Done()
				for j := 0; j < 10_000; j++ {
					atomic.AddInt64(&m.Counter1, 1) // Core 1
				}
			}()

			go func() {
				defer wg.Done()
				for j := 0; j < 10_000; j++ {
					atomic.AddInt64(&m.Counter2, 1) // Core 2
				}
			}()

			wg.Wait()
		}
	})

	b.Run("Padded (Independent Cache Lines)", func(b *testing.B) {
		m := &PaddedMetrics{}
		var wg sync.WaitGroup

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Add(2)
			go func() {
				defer wg.Done()
				for j := 0; j < 10_000; j++ {
					atomic.AddInt64(&m.Counter1, 1)
				}
			}()

			go func() {
				defer wg.Done()
				for j := 0; j < 10_000; j++ {
					atomic.AddInt64(&m.Counter2, 1)
				}
			}()

			wg.Wait()
		}
	})

	// When the student fixes `PaddedMetrics`, the second benchmark will run
	// significantly faster (often 2x-4x faster on modern multi-core CPUs).
}
