package garbagecollector

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSessionCacheEvictionOOM(t *testing.T) {
	// 1. Initialize our cache.
	// We mandate it should cap out at roughly 10 items for the test.
	// Since we didn't force the struct footprint in the prompt, we enforce it structurally:
	// A correct implementation will not keep growing indefinitely.

	cache := NewSessionCache() // User must modify constructor to accept a MaxKeys if necessary, or hardcode it.
	// For testing flexibility, we don't strict-type the constructor, we observe behavior.

	// We will insert 10,000 users. Each user is 1KB.
	// That's 10MB of memory.
	// If they bounded the cache to e.g. 100 items (100KB), memory should flatline.

	// We force a manual GC to get a clean baseline.
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	for i := 0; i < 10000; i++ {
		id := fmt.Sprintf("user_%d", i)
		cache.Set(id, User{ID: id})
	}

	// Force a GC. If the cache is unbounded, all 10,000 users are retained.
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// Calculate total memory held onto by the application (HeapAlloc).
	retainedBytes := m2.HeapAlloc - m1.HeapAlloc

	// If they held onto all 10,000 users, it's roughly 10MB.
	// If they bounded it properly (e.g., evicted items so it holds < 1000),
	// the GC successfully cleaned up the rest, and HeapAlloc will be tiny.

	if retainedBytes > 8*1024*1024 { // 8MB threshold
		t.Fatalf("FAILED: Your cache kept ~%d bytes retained in memory. The GC couldn't free them because you inserted all 10,000 users without evicting old ones. Implement bounded eviction!", retainedBytes)
	}

	// Also ensure we can actually retrieve a recent one.
	// In a random evictor or LRU, the _very last_ item inserted is probably still there.
	_, ok := cache.Get("user_9999")
	if !ok {
		t.Logf("WARN: User 9999 was evicted instantly. Your eviction logic might be overly aggressive, but you prevented an OOM!")
	}
}
