package timeutil

import (
	"testing"
	"time"
)

func TestUptimeProtectsAgainstNegative(t *testing.T) {
	// Let's simulate a wildly incorrect future date (an NTP rollback scenario)
	// Server thinks it started TOMORROW.
	futureStart := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")

	duration := Uptime(futureStart)

	if duration < 0 {
		t.Fatalf("CRITICAL: Uptime resulted in a negative duration (%v). You must protect against Wall Clock jumps!", duration)
	}

	if duration != 0 {
		t.Fatalf("Expected exactly 0 for negative durations, got %v", duration)
	}
}
