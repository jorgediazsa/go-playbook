package channels

import "testing"

func TestDirectionalChannels(t *testing.T) {
	// The real test here is whether ProvideStream and ConsumeAnalytics compile
	// with the restricted read-only types.

	stream := ProvideStream()
	count := ConsumeAnalytics(stream)

	// The internal producer writes 3 legitimate logs.
	if count != 3 {
		t.Fatalf("Expected 3 valid logs to be consumed, got %d. Did a fake log sneak in?", count)
	}
}
