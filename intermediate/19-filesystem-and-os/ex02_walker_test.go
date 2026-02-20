package filesystemos

import (
	"testing"
	"testing/fstest"
)

func TestCountLogFilesVirtual(t *testing.T) {
	// 1. We define an IN-MEMORY filesystem. No disk I/O required!
	// This makes the test hundreds of times faster and completely deterministic.

	virtualFS := fstest.MapFS{
		"root.log":                    &fstest.MapFile{Data: []byte("log1")},
		"some_image.png":              &fstest.MapFile{Data: []byte("img")},
		"deep/nested/folder/app.log":  &fstest.MapFile{Data: []byte("log2")},
		"deep/nested/folder/db.trace": &fstest.MapFile{Data: []byte("trace")},
		"deep/nested/second.log":      &fstest.MapFile{Data: []byte("log3")},
	}

	count, err := CountLogFiles(virtualFS)
	if err != nil {
		t.Fatalf("FAILED: CountLogFiles returned an error: %v", err)
	}

	// 2. Verify we navigated deep sub-directories correctly.
	if count != 3 {
		t.Fatalf("FAILED: Expected exactly 3 `.log` files, but counted %d. Did you use `fs.WalkDir` to recursively explore the tree?", count)
	}
}
