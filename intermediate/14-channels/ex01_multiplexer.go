package channels

// Context: Multiplexing with Select
// You are building a fast, highly-available Search endpoint. To guarantee
// a 50ms response, you query 3 different backend Search services simultaneously.
// Whichever one responds first, you return. If none respond within 50ms,
// you return a timeout error.
//
// Why this matters: `select` allows a goroutine to wait on multiple channel
// operations. Combined with unbuffered channels or buffered 1 channels, it is
// the cornerstone of cancellation, timeouts, and multiplexing in Go.
//
// Requirements:
// 1. Refactor `SearchFastest` to use a `select` block.
// 2. Query all three providers concurrently (spin up 3 goroutines).
// 3. The first one to send a result on a shared channel wins.
// 4. If 50ms passes, return "timeout" immediately.
// 5. Ensure you don't leak the slower goroutines! (Hint: use a buffered channel of size 3,
//    or pass them a context they can check, though a buffered channel is simpler here).

type Provider func(query string) string

func SearchFastest(query string, p1, p2, p3 Provider) string {
	// BUG: This queries them sequentially, taking the sum of all their times!
	// TODO: Query them concurrently and return the FASTEST result.
	// Add a 50ms timeout.

	res := p1(query)
	if res != "" {
		return res
	}

	res = p2(query)
	if res != "" {
		return res
	}

	return p3(query)
}
