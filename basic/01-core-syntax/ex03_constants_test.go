package core

import (
	"testing"
)

func TestComputeScale(t *testing.T) {
	// This test asserts that the student successfully implemented (Base * Multiplier) / Divisor
	// without intermediate overflow.

	result := ComputeScale()
	var expected int64 = 5000000000000000000

	if result != expected {
		t.Fatalf("ComputeScale() = %v, want %v", result, expected)
	}
}
