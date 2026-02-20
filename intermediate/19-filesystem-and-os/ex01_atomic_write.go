package filesystemos

import (
	"os"
)

// Context: Atomic File Writes & Crash Safety
// You are building an embedded IoT camera that constantly updates its `state.json`
// config file on a slow SD card. Power outages are common.
//
// Why this matters: If a power outage happens EXACTLY while `os.WriteFile`
// is heavily writing to the disk, the file will be brutally truncated and corrupted.
// Standard writes are NOT crash-safe.
//
// Requirements:
// 1. Refactor `WriteConfigAtomically` to prevent corruption.
// 2. You MUST write the data to a temporary file first (e.g., `filename + ".tmp"`).
// 3. You MUST then rename that temporary file over the final `filename` using `os.Rename`.
//    (In Unix, `rename` is an atomic system call).

func WriteConfigAtomically(filename string, data []byte) error {
	// BUG: This writes directly to the target file. It is NOT atomic.
	//
	// TODO:
	// 1. Write the payload to `<filename>.tmp` first.
	// 2. If it succeeds, rename `<filename>.tmp` -> `<filename>`.
	// 3. Ensure any error returns immediately, and clean up the tmp file if needed
	//    (though deferred os.Remove is best effort here).

	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
