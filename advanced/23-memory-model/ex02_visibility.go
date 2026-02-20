package memorymodel

// Context: Visibility Guarantees and Data Races
// You are building an orchestrator that can be manually paused by an admin via
// a `/pause` HTTP endpoint. The orchestrator checks this boolean natively in its
// hot loop.
//
// Why this matters: `adminPaused` is written aggressively by the HTTP handler
// (Core 1) and read aggressively by the orchestrator (Core 2).
// The Go Memory Model explicitly states that without synchronization
// (a "happens-before" edge), Core 2 is mathematically NOT GUARANTEED to ever
// see the write made by Core 1. Core 2 might cache `adminPaused == false`
// in its L1 cache and spin forever, ignoring the admin's pause command!
//
// Requirements:
// 1. Run the test with the race detector: `go test -race ./...`
// 2. Refactor `adminPaused` into an `atomic.Bool` to guarantee hardware-level
//    visibility and synchronization across CPU cores.
// 3. Alternatively, use a `sync.RWMutex`, which mathematically establishes a
//    happens-before edge when unlocking -> locking.
// 4. Update the `Pause` and `IsPaused` methods to use you chosen synchronization.

import (
	"time"
)

type Orchestrator struct {
	// BUG: A raw boolean read/written concurrently causes a Data Race and
	// visibility failures.
	// TODO: Replace with an `atomic.Bool` or a `sync.Mutex`/`RWMutex`.

	adminPaused bool
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{}
}

func (o *Orchestrator) Pause(state bool) {
	// BUG: Unsafe concurrent write.
	o.adminPaused = state
}

func (o *Orchestrator) IsPaused() bool {
	// BUG: Unsafe concurrent read.
	return o.adminPaused
}

func (o *Orchestrator) StartHotLoop() {
	go func() {
		for {
			if o.IsPaused() {
				time.Sleep(10 * time.Millisecond)
				continue
			}
			// Do heavy work...
			time.Sleep(1 * time.Millisecond)
		}
	}()
}
