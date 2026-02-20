package tags

import (
	"os/exec"
	"strings"
	"testing"
)

// This test verifies the build tags by spawning sub-processes to build the package
// with different flags.

func TestBuildTags(t *testing.T) {
	// 1. Build WITHOUT tags (should use safe)
	cmd := exec.Command("go", "test", "-v", "-run", "TestInternalHash")
	out, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(out), "Hash redeclared") {
			t.Fatalf("FAILED: Both files compiled at the same time! You must make them mutually exclusive using build tags. Error: %s", string(out))
		}
		t.Fatalf("Failed to build without tags: %v\nOutput: %s", err, string(out))
	}

	// 2. Build WITH tags (should use fast)
	cmd = exec.Command("go", "test", "-tags", "fast", "-v", "-run", "TestInternalHashFast")
	out, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to build with `fast` tag: %v\nOutput: %s", err, string(out))
	}
}

// These internal tests are called by the subprocesses above.

func TestInternalHash(t *testing.T) {
	// Without the fast tag, hash("A") should be 65 (safe)
	result := Hash("A")
	if result == 9999 {
		t.Fatal("The `fast` algorithm ran even though no tags were provided!")
	}
	if result != 65 {
		t.Fatalf("Expected safe hash 65, got %d", result)
	}
}

func TestInternalHashFast(t *testing.T) {
	// With the fast tag, ANY string should hash to 9999
	result := Hash("A")
	if result != 9999 {
		t.Fatalf("Expected fast hash 9999 when built with -tags fast, but got %d", result)
	}
}
