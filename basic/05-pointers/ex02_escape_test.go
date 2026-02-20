package pointers

import "testing"

func TestProcessStream(t *testing.T) {
	// If the user fixes compilation, this will pass.
	// The real test here is manual execution of escape analysis to verify the heap saving!

	count := ProcessStream()
	if count != 1000 {
		t.Fatalf("Expected 1000 processed events, got %d", count)
	}

	// Optional manual verification for the student:
	// Ask them to run `go build -gcflags="-m" ./ex02_escape.go`
	// Before their fix, `&e escapes to heap`.
	// After their fix, `CreateEvent` doesn't escape anything.
}
