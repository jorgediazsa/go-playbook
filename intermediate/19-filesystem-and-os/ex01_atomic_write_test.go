package filesystemos

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAtomicWriteTempFile(t *testing.T) {
	// 1. Setup a dummy workspace
	targetDir := t.TempDir()
	targetFile := filepath.Join(targetDir, "state.json")

	// 2. Perform the atomic write
	data := []byte(`{"version": 1}`)
	err := WriteConfigAtomically(targetFile, data)

	if err != nil {
		t.Fatalf("FAILED: WriteConfigAtomically returned error: %v", err)
	}

	// 3. Verify the final file exists
	content, err := os.ReadFile(targetFile)
	if err != nil {
		t.Fatalf("FAILED: The target file was not created: %v", err)
	}

	if string(content) != `{"version": 1}` {
		t.Fatalf("FAILED: incorrect content written: %q", string(content))
	}

	// 4. Read the directory to verify no `.tmp` files leaked behind.
	files, _ := os.ReadDir(targetDir)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".tmp") {
			t.Fatalf("FAILED: You left the `.tmp` file dangling: %s! Ensure you are calling os.Rename(), which deletes the tmp file natively if it succeeds.", f.Name())
		}
	}

	// How do we enforce that they didn't just call os.WriteFile?
	// It's very difficult to mock the OS mid-write strictly without intercepting syscalls
	// or injecting an `fs` interface (which os.Rename doesn't support natively).
	// For this exercise, visual inspection is assumed, but we check for lingering files.
}
