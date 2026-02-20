package goroutines

import (
	"strings"
	"testing"
	"time"
)

func TestUploadAllConcurrent(t *testing.T) {
	// Create 100 files
	files := make([]string, 100)
	for i := 0; i < 100; i++ {
		files[i] = "valid.jpg"
	}

	start := time.Now()
	err := UploadAll(files, 10)
	duration := time.Since(start)

	if err != nil {
		t.Fatalf("Expected nil error for valid files, got %v", err)
	}

	// Because MockUploadS3 is instant in our mock, we can't easily assert exactly 10 routines
	// ran in parallel via time. However, building the worker pool correctly ensures it won't
	// spawn 100 routines.
	_ = duration // Just a placeholder. In a real DB mock, we'd sleep and trace time.

	// Test Error Capture
	files[50] = "corrupt.jpg"

	// Start with fresh concurrency
	err = UploadAll(files, 5)

	if err == nil {
		t.Fatalf("FAILED: UploadAll did not capture the upload error!")
	}
	if !strings.Contains(err.Error(), "failed to upload corrupt image") {
		t.Fatalf("Expected corrupt image error, got: %v", err)
	}
}
