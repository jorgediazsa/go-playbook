package defers

// Context: Panic Boundaries in Goroutines
// You are building a scatter-gather fan-out system.
// A main coordinator spawns workers to process chunks of an array.
// What happens if one of the workers encounters a nil pointer and panics?
//
// Why this matters: An unrecovered panic in ANY goroutine crashes the entire process.
// The main function cannot `recover()` panics that occur in spawned goroutines.
// Every single goroutine you spawn MUST have its own panic recovery boundary if
// it executes potentially unsafe or third-party code.
//
// Requirements:
// 1. `RunWorkers` spawns a goroutine to process a job.
// 2. Wrap the goroutine execution in a panic recovery boundary.
// 3. When a panic is recovered, return the panic message wrapped as an error
//    to the `errCh` channel, so the coordinator can log it safely.

func RunWorkers(jobID int, errCh chan error) {
	go func() {
		// BUG: No panic recovery boundary!
		// If ProcessJob panics, the entire program dies instantly.
		// TODO: Add a defer/recover block here. If a panic happens, capture the value
		// and send `fmt.Errorf("worker panicked: %v", recovered_val)` down the errCh.
		// If no panic, send nil.

		ProcessJob(jobID)

		errCh <- nil
	}()
}

// Simulated chaotic third-party library
func ProcessJob(id int) {
	if id == 42 { // UNLUCKY NUMBER
		panic("nil pointer dereference")
	}
}
