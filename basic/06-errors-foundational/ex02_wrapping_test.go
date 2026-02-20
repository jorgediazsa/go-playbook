package errorsfunc

import (
	"errors"
	"strings"
	"testing"
)

func TestOrchestrateFetchFail(t *testing.T) {
	err := Orchestrate("user123")

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	// 1. Does it have the required context?
	if !strings.Contains(err.Error(), "failed to fetch user:") {
		t.Errorf("Error missing context: %v", err)
	}

	// 2. Does it retain the original root cause?
	if !errors.Is(err, ErrConnectionLost) {
		t.Errorf("Error was not properly wrapped using %%w. errors.Is() failed.")
	}
}
