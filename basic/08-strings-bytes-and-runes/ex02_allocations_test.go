package stringsbytes

import (
	"testing"
)

func TestGenerateCSVRow(t *testing.T) {
	data := []int{10, 20, 30}

	result := GenerateCSVRow(data)
	expected := "10,20,30"

	if result != expected {
		t.Fatalf("Expected %q, got %q", expected, result)
	}
}

// Benchmark shows the massive difference when implemented correctly
func BenchmarkGenerateCSVRow(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Run go test -bench . -benchmem
		// If you use += it will allocate MBs of data.
		// If you use strings.Builder, it will be 1 or 2 allocations.
		_ = GenerateCSVRow(data)
	}
}
