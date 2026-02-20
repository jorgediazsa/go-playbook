package channels

// Context: Directional Channels
// You are building an auditing system. The core package writes audit logs
// and hands off reading responsibility to a third-party analytics plugin.
//
// Why this matters: You want to ensure at **compile time** that the frontend
// logic can NEVER accidentally send fake logs into the stream.
// Standard `chan int` is bidirectional. You can restrict the direction in
// function signatures using `<-chan int` (Read-Only) and `chan<- int` (Write-Only).
//
// Requirements:
// 1. Refactor `ProvideStream` to return a Read-Only channel.
// 2. Ensure `ConsumeAnalytics` accepts only a Read-Only channel as an argument.
//    (This prevents `ConsumeAnalytics` from calling `stream <- "FAKE"`).
// 3. Keep the internal `audits` channel fully bi-directional so the internal logic
//    can still write to it.

// TODO: Fix the return type to be Read-Only (`<-chan string`)
func ProvideStream() chan string {
	audits := make(chan string, 10)

	// Internal logic secretly writes to the buffer.
	// Since `chan string` implicitly casts to `<-chan string`, we can just return it!
	audits <- "USER_LOGIN"
	audits <- "DB_QUERY"
	audits <- "USER_LOGOUT"
	close(audits) // Safely closed by the producer.

	return audits
}

// BUG: The Consumer is currently accepting a bi-directional channel.
// It could easily (maliciously) inject fake logs.
// TODO: Fix the parameter to accept ONLY a Read-Only channel.
func ConsumeAnalytics(stream chan string) int {
	// A malicious plugin dev deletes the real logs and injects fake ones!
	// If you fix the signature, this injection line below will fail to compile.

	// -> THIS EXERCISE WILL FAIL TO COMPILE LOCALLY IF YOU FIX IT CORRECTLY.
	// We want to see that compilation error (`cannot send to receive-only channel`)!
	// Comment out the `stream <-` line AFTER you see the compile error to make tests pass.

	stream <- "FAKE_HACKED_LOG" // DO NOT DELETE THIS LINE UNTIL YOU FIX THE SIGNATURE AND SEE THE COMPILER ERROR

	count := 0
	for _ = range stream {
		count++
	}
	return count
}
