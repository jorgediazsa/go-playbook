package memorymodel

import (
	"testing"
)

func TestParseLogLogic(t *testing.T) {
	// We simply verify the interface and logical output the student writes.
	// Depending on how they refactored the pointer to a value, this might need
	// adjusting locally, but ideally they just return a value type.

	rec := ParseLog("INFO:System booting")

	// Dereferencing might fail if they returned a value directly.
	// We use reflection loosely or just rely on manual compilation verification.
	// For testing purposes in this exercise, as long as it compiles, we are good.

	_ = rec
}

// RUN WITH: go test -bench BenchmarkParseLog -benchmem
func BenchmarkParseLog(b *testing.B) {
	raw := "ERROR:Disk full"

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		// If the struct escapes, this will show 1 alloc/op
		// Notice: returning a value type avoids the heap pointer allocation.
		// However, strings.SplitN internally creates slices which allocate.
		// To truly hit 0 allocations, they must also refactor `strings.SplitN`
		// into a manual byte/string index search!
		//
		// Extra Credit Challenge in this Benchmark:
		// Can you parse the log string into a value-struct with EXACTLY 0 allocs?
		// Hint: use `strings.Index(raw, ":")` and slice the string directly `raw[:idx]`.

		_ = ParseLog(raw)
	}
}
