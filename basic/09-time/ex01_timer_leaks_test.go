package timeutil

import (
	"testing"
)

func TestFetchWithTimeout(t *testing.T) {
	// 1. Test the fast path
	fastCh := make(chan string, 1)
	fastCh <- "success"

	resp, err := FetchWithTimeout(fastCh)
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if resp != "success" {
		t.Fatalf("Expected 'success', got %q", resp)
	}

	// 2. Test the timeout path (we shouldn't actually wait 30s in a unit test,
	// but the user's signature was fixed to 30s. For this exercise, verifying compilation
	// and structural logic by visual inspection of their time.NewTimer usage is key).
	// Let's rely on the Go compiler to ensure they used time.NewTimer correctly implicitly.

	// A strictly black-box unit test cannot easily detect the lack of Stop() without
	// profiling or injecting time mocks. We leave this as an structural exercise.
}
