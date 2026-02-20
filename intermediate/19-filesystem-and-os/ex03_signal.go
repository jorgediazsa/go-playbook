package filesystemos

import (
	"context"
	"fmt"
	"time"
)

// Context: Graceful Shutdown (OS Signals)
// You have an HTTP server running inside a Kubernetes cluster. When Kubernetes
// decides to deploy a new version, it sends a `SIGTERM` (terminate) OS signal
// to your process.
//
// Why this matters: If you don't trap this signal, Linux instantly kills your
// process, terminating in-flight HTTP requests and causing database corruption.
// Idiomatic Go intercepts these signals, flips the server into "Shutdown" mode,
// waits for active connections to finish, and exits cleanly.
//
// Requirements:
// 1. You are given a `RunServer` function. It currently loops endlessly.
// 2. We inject a `ctx`. Your orchestrator (`main.go` or tests) uses `signal.NotifyContext`
//    to automatically cancel this context when a SIGTERM or SIGINT is received!
// 3. Refactor the endless loop. At every "tick", use a `select` statement:
//    - If the `ticker` fires, simulate work.
//    - If the `ctx.Done()` channel is closed, break the loop completely to shut down cleanly.

func RunServer(ctx context.Context) string {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	workDone := 0

	// BUG: This loop runs FOREVER, ignoring OS signals completely!
	// When Docker sends SIGTERM, this process will just ignore it until the OS
	// resorts to violently SIGKILL-ing it 30 seconds later.
	// TODO: Replace this `for` and `sleep(ticker)` with a `for` enclosing a `select`.

	for {
		<-ticker.C
		workDone++
	}

	return fmt.Sprintf("Shutdown Cleanly. Processed %d tasks.", workDone)
}
