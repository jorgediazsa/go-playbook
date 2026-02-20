package defers

import "fmt"

// Context: Defer order execution (LIFO) and scoping
// You are managing temporary file systems for a build pipeline.
//
// Why this matters: `defer` evaluates the function variables at the time the
// defer is *declared*, but executes the function at the time the surrounding
// function *returns*. And multiple defers execute Last-In, First-Out.
//
// Requirements:
// 1. `ReleaseResources` mimics acquiring locks, DBs, and files.
// 2. Fix the `defer` statements so that the sequence of cleanup actions
//    recorded in the `cleanupLog` slice matches the expected LIFO order:
//    "File_Closed", "DB_Closed", "Lock_Released".
// 3. Fix the variable scoping bug where the `defer` logs the wrong filename.

func ReleaseResources() []string {
	var cleanupLog []string

	addLog := func(msg string) {
		cleanupLog = append(cleanupLog, msg)
	}

	// 1. Acquire Lock
	// TODO: Fix defer order!
	defer addLog("Lock_Released")

	// 2. Acquire DB
	defer addLog("DB_Closed")

	// 3. Acquire File
	filename := "temp_file.txt"
	defer func() {
		// BUG: This captures `filename` by reference, but filename is mutated below!
		// It will log "wrong_file.txt_Closed".
		// TODO: Fix the capture so it logs the correct filename at the time of deferring.
		addLog(fmt.Sprintf("%s_Closed", filename))
	}()

	// Simulating other logic that changes the variable...
	filename = "wrong_file.txt"

	return cleanupLog
}
