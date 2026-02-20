package observability

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

// Context: CPU Contention and Memory Leaks
// You are profiling a background worker that hashes massive strings.
// When you run the loader (`go run ./cmd/loader/main.go`), cpu usage spikes to
// 100%, and RAM usage climbs indefinitely.
//
// Why this matters: You must learn to read `pprof` flame graphs. If you profile
// this code, you will discover that 95% of the CPU goes to `fmt.Sprintf` and
// `strings.Split`, and the RAM leak is caused by an unbounded cache.
//
// Requirements:
// 1. Run the loader (`go run cmd/loader/main.go`).
// 2. Open a new terminal and generate a CPU profile:
//    `go tool pprof -http=:8080 "http://localhost:6060/debug/pprof/profile?seconds=5"`
// 3. Open a new terminal and generate a Heap profile:
//    `go tool pprof -http=:8081 "http://localhost:6060/debug/pprof/heap"`
// 4. Refactor `ProcessData` below based on what you find:
//    - Eliminate `fmt.Sprintf` for simple string concatenation.
//    - Eliminate `strings.Split` if you only need the first part (use `strings.Index`).
//    - Fix the memory leak (unbounded `resultsCache`).

var resultsCache = make(map[string]string)

func ProcessData(id string, payload string) string {
	// BUG: Unbounded cache causes a massive memory leak!
	// TODO: For this exercise, just disable the cache entirely (delete the map writes)
	// or implement a quick random eviction once it hits 1000 items.

	// BUG: `fmt.Sprintf` is incredibly slow and allocation-heavy for simple concatenation.
	// TODO: Just use `id + ":" + payload`
	combined := fmt.Sprintf("%s:%s", id, payload)

	// BUG: `strings.Split` allocates a brand new slice of strings, destroying cache locality.
	// TODO: Replace with `idx := strings.Index(combined, "|")` and `combined[:idx]`
	parts := strings.Split(combined, "|")
	baseStr := parts[0]

	// Simulate heavy but necessary work
	hash := sha256.Sum256([]byte(baseStr))
	out := fmt.Sprintf("%x", hash)

	// Leak!
	resultsCache[id] = out

	// To prevent the test/loader from finishing in 1 millisecond, we slow it down artificially.
	time.Sleep(1 * time.Millisecond)

	return out
}
