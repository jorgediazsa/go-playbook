package syncprims

import (
	"sync"
	"testing"
	"time"
)

func TestBoundedBufferCond(t *testing.T) {
	buffer := NewBoundedBuffer(2)

	var wg sync.WaitGroup

	// Launch a slow consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond) // Wait before consuming

		val1 := buffer.Consume()
		if val1 != "item1" {
			t.Errorf("Expected item1, got %v", val1)
		}

		val2 := buffer.Consume()
		if val2 != "item2" {
			t.Errorf("Expected item2, got %v", val2)
		}

		val3 := buffer.Consume()
		if val3 != "item3" {
			t.Errorf("Expected item3, got %v", val3)
		}
	}()

	// Producer
	start := time.Now()

	// Fills immediately
	buffer.Produce("item1")
	buffer.Produce("item2")

	// Should BLOCK here until the consumer wakes up at 100ms and pulls an item!
	buffer.Produce("item3")

	duration := time.Since(start)
	if duration < 50*time.Millisecond {
		t.Fatalf("FAILED: Produce() did not block when the buffer was full! It returned instantly! Did you use `Cond.Wait()` correctly in a loop?")
	}

	wg.Wait()
}
