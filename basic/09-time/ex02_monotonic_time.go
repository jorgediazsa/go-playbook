package timeutil

import (
	"time"
)

// Context: Monotonic Time vs Wall Clock formatting
// You are logging the exact startup time of your service to a JSON file.
// Later, you read that file and calculate how long the service has been running.
//
// Why this matters: `time.Since` relies on the Monotonic clock reading embedded
// inside a `time.Time` object. When you call `time.Now().Format(...)`, or serialize
// it to JSON, the monotonic clock reading is STRIPPED.
// If you deserialize it and call `time.Since()`, it relies entirely on the Wall clock.
// If an NTP sync reversed the server's clock by 1 minute, `time.Since()` could
// return a NEGATIVE duration (e.g., -59 seconds).
//
// Requirements:
// 1. Fix `Uptime` so it does not result in logic bugs if the system clock
//    has been slightly shifted backwards by NTP.
// 2. Ensure that durations are positive or zero.

func FormattedStartTime() (string, time.Time) {
	now := time.Now()
	// Formats as: 2006-01-02 15:04:05
	return now.Format("2006-01-02 15:04:05"), now
}

func Uptime(serializedStart string) time.Duration {
	// BUG: The serialized date has no monotonic clock.
	// We parse it back into a time.Time object using only the Wall clock.
	parsedStart, _ := time.Parse("2006-01-02 15:04:05", serializedStart)

	// If the server clock jumped backwards by 1 minute between FormattedStartTime
	// and now, `time.Since` will produce a negative duration!
	// TODO: Fix this function to at least protect against returning negative durations.
	// If the duration is negative, return 0 (meaning we can't reliably detect uptime).

	uptime := time.Since(parsedStart)
	return uptime
}
