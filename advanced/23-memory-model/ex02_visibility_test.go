package memorymodel

import (
	"testing"
	"time"
)

func TestOrchestratorVisibility(t *testing.T) {
	orch := NewOrchestrator()

	// Start the background reader goroutine
	orch.StartHotLoop()

	// Spin up 10 writer goroutines slamming the variable
	for i := 0; i < 10; i++ {
		go func(state bool) {
			orch.Pause(state)
		}(i%2 == 0) // alternate true/false
	}

	time.Sleep(50 * time.Millisecond)

	// Verify we can read it safely without crashing
	_ = orch.IsPaused()

	// If the user runs `go test -race`, and they didn't fix `adminPaused`,
	// the race detector will scream loudly and fail the test.
	// We mandate `-race` in our project READMEs for exactly this reason!
}
