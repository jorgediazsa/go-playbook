package filesystemos

import (
	"io/fs"
	"strings"
)

// Context: Abstraction using `io/fs`
// You are building a Log Search Engine. It traverses a directory structure
// to count how many files end in `.log`.
//
// Why this matters: Historically, Go developers used `filepath.Walk` and
// passed hard-coded paths (`/var/log`). This made testing a nightmare because
// CI pipelines needed real files on disk.
// Go 1.16 introduced `io/fs`. Now, your search engine accepts an `fs.FS` interface.
// Production passes `os.DirFS("/var/log")`. Tests pass `testing/fstest.MapFS{}`
// (a purely in-memory virtual filesystem)!
//
// Requirements:
// 1. Refactor `CountLogFiles` to traverse the `fileSystem` boundary using `fs.WalkDir`.
// 2. Count any file whose name ends with `.log`.
// 3. Skip directories entirely (do not count them as log files!).

func CountLogFiles(fileSystem fs.FS) (int, error) {
	count := 0

	// BUG: The current implementation iterates over the root instantly, but does not
	// recurse into subdirectories like /year/month/day/app.log!
	// TODO: Replace this loop with `fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error)`.

	entries, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".log") {
			count++
		}
	}

	return count, nil
}
