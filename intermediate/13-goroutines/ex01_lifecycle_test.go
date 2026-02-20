package goroutines

import (
	"runtime"
	"testing"
	"time"
)

func TestCacheLifecycle(t *testing.T) {
	initialRoutines := runtime.NumGoroutine()

	cache := NewCache()
	cache.Start()

	// Allow it to spin up
	time.Sleep(20 * time.Millisecond)

	routineCountWhileRunning := runtime.NumGoroutine()
	if routineCountWhileRunning <= initialRoutines {
		t.Fatalf("Cache.Start() did not spawn a background goroutine!")
	}

	// Stop it. If implemented correctly, Stop() will block until evict() finishes.
	cache.Stop()

	// Wait a tiny bit just in case, though Stop() *should* be synchronous.
	time.Sleep(5 * time.Millisecond)

	finalRoutines := runtime.NumGoroutine()
	if finalRoutines > initialRoutines {
		t.Fatalf("LEAK DETECTED: Expected %d goroutines after Stop(), got %d. The background loop did not exit.", initialRoutines, finalRoutines)
	}
}
