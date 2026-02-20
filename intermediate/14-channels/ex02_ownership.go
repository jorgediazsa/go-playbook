package channels

// Context: Channel Ownership and Closing
// You are building an event processing pipeline. You have a central event
// channel, and 3 producers (sensors) writing to it concurrently.
// A consumer reads from the channel until it's closed.
//
// Why this matters: The Golden Rule of Go channels is: "The sender closes the channel."
// But what happens when you have MULTIPLE senders?
// If Sensor 1 finishes and calls `close(ch)`, Sensor 2 will immediately PANIC
// with "send on closed channel" when it tries to write its next event.
//
// Requirements:
// 1. Refactor `StartPipeline` to spawn the 3 producers concurrently.
// 2. Safely close the `events` channel ONLY after ALL 3 producers have finished.
//    (Hint: Use a `sync.WaitGroup` to track the producers, and a dedicated
//     goroutine to `Wait()` on them and then close the channel).
// 3. Return the `count` of events processed by the consumer.

func StartPipeline(sensorData [][]string) int {
	events := make(chan string)

	// Create 3 producers (one for each slice in sensorData)
	// BUG: The current code runs sequentially, and the single sender closes it.
	// TODO: Spawn them concurrently. Coordinate the close() safely without panicking.

	for _, dataChunk := range sensorData {
		for _, e := range dataChunk {
			events <- e
		}
	}
	close(events) // WRONG place for a multi-producer setup!

	// Consumer
	count := 0
	for _ = range events {
		count++
	}

	return count
}
