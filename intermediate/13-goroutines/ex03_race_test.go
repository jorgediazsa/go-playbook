package goroutines

import (
	"sync"
	"testing"
)

func TestAggregatorRace(t *testing.T) {
	agg := &Aggregator{}
	var wg sync.WaitGroup

	// We spawn 1,000 goroutines hitting the aggregator simultaneously.
	// Without a Mutex, the final result will be highly unpredictable (usually much lower than 10,000),
	// and `go test -race` will SCREAM at you.

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			agg.ReceiveChunk(10) // 10 bytes per chunk
		}()
	}

	wg.Wait()

	total := agg.GetTotal()
	if total != 10000 {
		t.Fatalf("RACE DETECTED: Expected total 10,000 bytes, but got %d! Check your Mutexes.", total)
	}
}
