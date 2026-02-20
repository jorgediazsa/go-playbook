package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestImportCycleResolved(t *testing.T) {
	// We use `go list` to attempt to parse the packages.
	// If the cycle still exists, `go list` will fail.

	cmd := exec.Command("go", "list", "./...")
	out, err := cmd.CombinedOutput()

	if err != nil {
		if strings.Contains(string(out), "import cycle not allowed") {
			t.Fatalf("FAILED: Import cycle detected!\n%s\nYou must refactor the packages to break the cycle (e.g., extract shared models).", string(out))
		}
		t.Fatalf("go list failed: %v", err)
	}

	// If it succeeds, the cycle is broken!
}
