package errorsadv

import (
	"errors"
	"testing"
)

// A mock struct that implements the behavior pattern.
type NetworkErr struct{ tmp bool }

func (e NetworkErr) Error() string   { return "network error" }
func (e NetworkErr) Temporary() bool { return e.tmp }

type FatalErr struct{}

func (e FatalErr) Error() string { return "fatal error" }

func TestExecuteWithRetryClassification(t *testing.T) {
	// Test 1: Fatal Error (Should exit instantly attempt 1)
	attempts := 0
	fatalOp := func() error {
		attempts++
		return FatalErr{}
	}

	err := ExecuteWithRetry(fatalOp)

	if !errors.Is(err, FatalErr{}) {
		t.Fatalf("FAILED: Expected to return the fatal error untouched, got %v", err)
	}
	if attempts > 1 {
		t.Fatalf("FAILED: It retried a fatal error! Attempts: %d", attempts)
	}

	// Test 2: Temporary Error (Should retry 3 times)
	attempts = 0
	tempOp := func() error {
		attempts++
		return NetworkErr{tmp: true}
	}

	_ = ExecuteWithRetry(tempOp)

	if attempts != 3 {
		t.Fatalf("FAILED: It failed to retry a temporary error! Attempts: %d", attempts)
	}
}
