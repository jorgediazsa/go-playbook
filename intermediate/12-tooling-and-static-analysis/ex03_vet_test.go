package vet

import (
	"os/exec"
	"strings"
	"testing"
)

func TestGoVetPasses(t *testing.T) {
	// The problem with `go vet` errors is that they still compile and run,
	// often producing strange output at runtime (e.g. %!d(string=INFO) ).

	// We use an exec sub-process to run `go vet`.
	cmd := exec.Command("go", "vet", ".")
	out, err := cmd.CombinedOutput()

	if err != nil {
		// go vet returns a non-zero exit code if it finds issues
		if strings.Contains(string(out), "verb %d") || strings.Contains(string(out), "arg") {
			t.Fatalf("FAILED: `go vet` caught a bug in your code!\nOutput:\n%s\nFix the formatting verbs!", string(out))
		}
		t.Fatalf("go vet failed unexpectedly: %v\n%s", err, string(out))
	}

	// Verify runtime output is also correct
	result := FormatLog("INFO", 12345)
	expected := "Level: INFO | RequestID: 12345"
	if result != expected {
		t.Errorf("FormatLog() = %q, want %q", result, expected)
	}
}
