package testingadv

// Context: Benchmarks and Allocations
// You are parsing thousands of comma-separated lines.
//
// Why this matters: `strings.Split` allocates a new slice on the heap, and
// allocates a new underlying backing array for that slice, and copies the pointers.
// In a hot loop processing 10 million lines, this creates massive Garbage Collection
// pressure.
//
// Requirements:
// 1. Run the benchmark: `go test -bench BenchmarkCountCommas -benchmem`. Note the allocations.
// 2. Refactor `CountCommas` to do the EXACT same logical work, but WITHOUT using
//    `strings.Split`. (Hint: Iterate over the string, or use `strings.Count`,
//    or `strings.Index`. For maximum speed, just loop `for i := 0; i < len(s); i++`).
// 3. Re-run the benchmark. You should achieve 0 allocations (`0 B/op, 0 allocs/op`).

import "strings"

func CountCommas(s string) int {
	// BUG: Very slow, highly allocative approach to counting commas.
	// TODO: Rewrite this for 0 allocations and massive speed.

	parts := strings.Split(s, ",")
	return len(parts) - 1
}
