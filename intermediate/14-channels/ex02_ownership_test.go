package channels

import "testing"

func TestStartPipeline(t *testing.T) {
	data := [][]string{
		{"s1_a", "s1_b", "s1_c"},
		{"s2_a", "s2_b"},
		{"s3_a", "s3_b", "s3_c", "s3_d"},
	}

	// If the user's refactor incorrectly closes the channel while other
	// producers are still sending, this will throw a fatal runtime Panic ("send on closed channel").
	// We catch that via standard panic trace in the test environment.

	totalProcessed := StartPipeline(data)

	// 3 + 2 + 4 = 9
	if totalProcessed != 9 {
		t.Fatalf("Expected 9 processed events, got %d. Did your consumer exit early or block forever?", totalProcessed)
	}
}
