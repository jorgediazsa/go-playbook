package infrastructure

import (
	"os"
	"strings"
	"testing"
)

// To make this test pass, the student must define the variables,
// but the test also checks if we dynamically injected them via the CLI!
// We can't easily force the `go test` command to include ldflags dynamically
// from inside the code, but we can check if they removed the hardcoded values.

func TestGetBuildInfo(t *testing.T) {
	commit, bTime := GetBuildInfo()

	if commit == "hardcoded" || bTime == "hardcoded" {
		t.Fatalf("FAILED: You are returning hardcoded values. You must return package-level variables that can be overridden by ldflags!")
	}

	// If they just defined `var GitCommit = "unknown"`, it will pass this generic check.
	// But let's check if the environment actually injected something when we test it manually.

	// We read the actual command line that started this test purely for educational logging
	cmd := strings.Join(os.Args, " ")
	if !strings.Contains(cmd, "ldflags") {
		t.Logf("WARN: You didn't run this test with ldflags! Try running:\n go test -v -ldflags=\"-X 'go-playbook/advanced/29-go-infrastructure.GitCommit=abcdef' -X 'go-playbook/advanced/29-go-infrastructure.BuildTime=yesterday'\" ./...")
	} else {
		t.Logf("SUCCESS! Injected Commit: %s, BuildTime: %s", commit, bTime)
	}
}
