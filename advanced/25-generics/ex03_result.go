package generics

// Context: The Result Pattern (Error Handling in Channels)
// You are building a concurrent web scraper. 100 goroutines are fetching URLs
// and sending the HTML back to a single aggregator over a channel.
//
// Why this matters: You cannot send `(string, error)` through a channel.
// A channel only takes one type. Historically, developers created custom structs
// like `type FetchResult struct { body string; err error }` for every single job.
// With generics, we can build a universal `Result[T]` struct that works for ANY
// type sent across a channel.
//
// Requirements:
// 1. Define a generic struct `Result[T any]` containing `Value T` and `Err error`.
// 2. Refactor `ScrapeWorker` so that the `results` channel accepts `Result[string]`.
// 3. Send successful creations or errors properly.

import (
	"context"
	"time"
)

// TODO: Define the generic `Result[T any]` struct here:
// type Result[T any] struct { ... }

// BUG: The channel is currently just `chan string`, meaning errors are silently
// dropped or cause a panic!
// TODO: Refactor `chan string` to `chan Result[string]`
func ScrapeWorker(ctx context.Context, urls []string, results chan string) {
	for _, url := range urls {
		select {
		case <-ctx.Done():
			return
		default:
		}

		// Simulate a network failure on http://bad.com
		if url == "http://bad.com" {
			// BUG: How do we send this error over the channel?
			// TODO: Send `Result[string]{Err: errors.New("network timeout")}`
			continue
		}

		time.Sleep(10 * time.Millisecond)

		// TODO: Send `Result[string]{Value: "<html>..."}`
		results <- "<html>Body of " + url + "</html>"
	}
}
