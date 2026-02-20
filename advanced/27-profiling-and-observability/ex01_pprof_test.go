package observability

import (
	"testing"
)

func TestProcessData(t *testing.T) {
	// Our test strictly verifies correctness.
	// The student must use external `pprof` tools to verify performance!

	out := ProcessData("user123", "very_large_payload|ignore_this_part")

	// sha256 of "user123:very_large_payload"
	expected := "8e5d2b7c6c449c25f4aabcf80b1dfc0b4c73abdf2ab96c7324aa8a9fac6f8fb9"

	if out != expected {
		t.Fatalf("FAILED: ProcessData changed output. Expected %s, got %s", expected, out)
	}
}
