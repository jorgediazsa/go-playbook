package channels

// Context: Bounded Buffering & Deadlocks
// You are designing an ingestion queue. External networks dump logs into this queue.
// An internal processing engine reads them off.
//
// Why this matters: If the queue is unbuffered (`make(chan string)`), the ingestion
// network instantly blocks if the processor is busy, causing timeouts to clients.
// If the queue has infinite capacity (not possible natively, but via wrappers),
// a slow processor will cause an Out-Of-Memory panic.
// An idiomatic Go bounded queue is simply a buffered channel: (`make(chan string, 100)`).
//
// Requirements:
// 1. Refactor `NewQueue` to create a buffered channel of size `capacity`.
// 2. Implement `Enqueue` so it DOES NOT block indefinitely. Use a `select` statement.
//    - If the channel is full, immediately return `ErrQueueFull`.
// 3. Implement `Dequeue` to pull an item.

import (
	"errors"
)

var ErrQueueFull = errors.New("queue is full")

type LogQueue struct {
	// TODO: Define the buffered channel
	stream chan string
}

func NewQueue(capacity int) *LogQueue {
	// BUG: Unbuffered channel
	return &LogQueue{
		stream: make(chan string),
	}
}

func (q *LogQueue) Enqueue(logMsg string) error {
	// BUG: This blocks forever if the processor is busy.
	// TODO: Use select with a default case to drop the message if the buffer is full!

	q.stream <- logMsg

	return nil
}

func (q *LogQueue) Dequeue() string {
	// Blocks until a log is available
	return <-q.stream
}
