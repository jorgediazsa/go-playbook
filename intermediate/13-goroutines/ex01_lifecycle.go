package goroutines

import (
	"time"
)

// Context: Goroutine Lifecycle & Leaks
// You are building an in-memory Cache. Every 10 seconds, it needs to scan
// its memory and evict expired items.
//
// Why this matters: If we start a `StartEvictionLoop()` goroutine, we MUST
// provide a way to stop it. If the user creates 5 Cache instances in tests
// and drops them, 5 background goroutines will leak and run forever, consuming CPU.
//
// Requirements:
// 1. Implement `Start()` to launch the background eviction loop.
// 2. Implement `Stop()` to cleanly signal the loop to exit.
// 3. Ensure `Stop()` blocks until the background loop has COMPLETELY exited.
//    (Hint: Use a `sync.WaitGroup`).
// 4. Do not use Context here (we will cover it in Topic 16). Use a `done` channel
//    or a boolean flag protected by a Mutex (a channel is usually cleaner combined with `select`).

type Cache struct {
	// TODO: Add fields here to manage the lifecycle (e.g., done channel, WaitGroup)
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Start() {
	// BUG: This goroutine leaks forever! It has no way to be stopped,
	// and we have no way to wait for it to finish gracefully.

	// TODO: Fix this lifecycle.
	go func() {
		for {
			time.Sleep(10 * time.Millisecond) // Simulated ticker
			c.evict()
		}
	}()
}

func (c *Cache) Stop() {
	// TODO: Signal the goroutine to stop, and block until it actually returns.
}

func (c *Cache) evict() {
	// Simulated work
}
