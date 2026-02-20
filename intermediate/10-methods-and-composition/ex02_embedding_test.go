package composition

import (
	"testing"
	"time"
)

func TestCounterEmbedding(t *testing.T) {
	counter := NewCounter("c-100")

	// If the user didn't initialize the embedded pointer, this will panic with a nil pointer dereference.
	var _ = counter.BaseMetric

	// 1. Does it start clean?
	origTime := counter.LastEvent
	if counter.Count != 0 {
		t.Fatalf("Expected count 0, got %d", counter.Count)
	}

	// Wait 1 ms to ensure time ticks
	time.Sleep(1 * time.Millisecond)

	// 2. Test the overridden method.
	// In the absence of an overriding `func (c *CounterMetric) Record()`,
	// this calls the promoted `(*BaseMetric).Record()`, which updates time but NOT the count!
	// If the user wrote the override, it will do both.

	counter.Record()

	if counter.Count != 1 {
		t.Fatalf("\nExpected count 1, got %d. \nYou are probably calling the promoted BaseMetric.Record() instead of shadowing it. \nDid you define `func (c *CounterMetric) Record()`?", counter.Count)
	}

	if !counter.LastEvent.After(origTime) {
		t.Fatalf("\nThe time was not updated. \nDid your overriding Record() method forget to explicitly call `c.BaseMetric.Record()`?")
	}
}
