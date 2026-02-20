package collections

import (
	"strconv"
	"sync"
	"testing"
)

func TestSessionCacheConcurrency(t *testing.T) {
	// If the implementation is wrong, run this test with `go test -race`
	// or it may just naturally panic with "concurrent map read and map write".

	cache := NewSessionCache()
	var wg sync.WaitGroup

	// Start 100 writer goroutines
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			cache.Set("key"+strconv.Itoa(n), "value")
		}(i)
	}

	// Start 100 reader goroutines
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			cache.Get("key" + strconv.Itoa(n))
		}(i)
	}

	wg.Wait()
	// If it didn't panic, they successfully protected the map!
}
