package defers

import (
	"strings"
	"testing"
)

func TestRunWorkersPanicRecovery(t *testing.T) {
	errCh := make(chan error, 1)

	// Since we are running tests, an unrecovered panic inside a goroutine
	// will crash the test runner entirely.
	// If the user fixes the code, this test will pass smoothly.

	RunWorkers(42, errCh)

	err := <-errCh
	if err == nil {
		t.Fatalf("Expected the panic to be caught and returned as an error, got nil")
	}

	if !strings.Contains(err.Error(), "worker panicked: nil pointer dereference") {
		t.Errorf("Recovered error did not match expected format. Got: %v", err)
	}
}
