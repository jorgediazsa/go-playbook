package syncprims

import "sync"

// Context: Bounded Buffer using sync.Cond
// You are building an ultra-low latency event router. You need a fast,
// bounded buffer (maximum 10 items).
//
// Why this matters: You could use a buffered channel `make(chan string, 10)`,
// but channels carry some allocation overhead. For sheer millions-per-second
// throughput, a statically allocated slice protected by a `Mutex` and
// coordinated via a `sync.Cond` is often faster.
//
// A `sync.Cond` has a Mutex (`L`). You lock the Mutex, check a condition in a
// `for` loop (e.g. `for isFull`), and if true, you call `Cond.Wait()`. When
// another goroutine changes the state, it calls `Cond.Signal()` to wake you.
//
// Requirements:
// 1. Refactor `BoundedBuffer` to use `sync.Cond` for signaling.
// 2. `Produce`: If the buffer is full (count == capacity), it MUST block and wait.
// 3. `Consume`: If the buffer is empty (count == 0), it MUST block and wait it.
// 4. Critical: Produce must `Signal()` after adding an item, and Consume must
//    `Signal()` after removing an item, to wake up waiting goroutines!

type BoundedBuffer struct {
	// TODO: Replace these basic Mutexes with a single Mutex and two Conditions
	// (e.g., `notFull` and `notEmpty`, or just one generic `cond := sync.NewCond(&mu)`).
	mu       sync.Mutex
	items    []string
	count    int
	capacity int
}

func NewBoundedBuffer(capacity int) *BoundedBuffer {
	return &BoundedBuffer{
		items:    make([]string, capacity),
		capacity: capacity,
	}
}

func (b *BoundedBuffer) Produce(item string) {
	// BUG: This silently overwrites items or panics if overflowing,
	// rather than blocking and waiting gracefully.
	// TODO: Lock. *Wait* in a loop while buffer is full. Insert item. Signal! Unlock.

	b.mu.Lock()
	defer b.mu.Unlock()
	b.items[b.count] = item
	b.count++
}

func (b *BoundedBuffer) Consume() string {
	// BUG: This panics if called when empty!
	// TODO: Lock. *Wait* in a loop while buffer is empty. Remove item. Signal! Unlock.

	b.mu.Lock()
	defer b.mu.Unlock()
	b.count--
	return b.items[b.count]
}
