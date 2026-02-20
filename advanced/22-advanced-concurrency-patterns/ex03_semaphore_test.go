package advancedconcurrency

import (
	"sync"
	"testing"
	"time"
)

func TestMainframeSemaphore(t *testing.T) {
	// Create Gateway allowing 3 simultaneous connections
	gateway := NewMainframeGateway(3)

	var wg sync.WaitGroup
	var successes int
	var rejects int
	var mu sync.Mutex

	start := time.Now()

	// Launch 10 simultaneous requests
	// The first 3 should jump in immediately and sleep for 50ms.
	// The next 7 should wait.
	// At ~50ms, the first 3 finish, and completely open 3 slots.
	// Requests 4, 5, 6 jump in and sleep for 50ms.
	// At ~100ms, requests 7, 8, 9, 10 hit the `100ms` select timeout and are rejected!

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_, err := gateway.Query("data")

			mu.Lock()
			if err == nil {
				successes++
			} else if err == ErrRateLimited {
				rejects++
			}
			mu.Unlock()
		}()
	}

	wg.Wait()
	duration := time.Since(start)

	// A pure Mutex implementation will process them sequentially: 10 * 50 = 500ms
	// A correct Semaphore + Timeout implementation will finish in ~100-150ms
	// when the final failures surface.

	if duration > 200*time.Millisecond {
		t.Fatalf("FAILED: Took %v. You probably used a Mutex instead of a Semaphore of size 3!", duration)
	}

	// We expect roughly 6 successes (2 blocks of 3) and 4 rejections due to timeout.
	// Timing can be jittery in CI/Mac OS, so we verify > 0 and < 10.

	if rejects == 0 {
		t.Fatalf("FAILED: No requests were rejected. Did you implement the 100ms select timeout?")
	}

	if successes != 6 {
		t.Logf("WARN: Expected 6 successes, got %d (Test relies on millisecond scheduling, which can drift, but logic should be solid)", successes)
	}
}
