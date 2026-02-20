package syncprims

import (
	"sync"
	"testing"
)

func TestDBManagerOnceValues(t *testing.T) {
	// Enable mock failure
	mockFail = true
	manager := NewDBManager()

	var wg sync.WaitGroup
	var errs []error
	var mu sync.Mutex

	// Hammer it with 50 concurrent requests.
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := manager.GetConnection()

			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
		}()
	}

	wg.Wait()

	// 1. Every single caller should have received ErrTimeout safely.
	for _, e := range errs {
		if e != ErrTimeout {
			t.Fatalf("FAILED: Expected ErrTimeout for all callers, got %v", e)
		}
	}

	// Reset to success
	mockFail = false
	manager2 := NewDBManager()
	conn, err := manager2.GetConnection()

	if err != nil || conn == nil || conn.Status != "OK" {
		t.Fatalf("Expected successful connection on second manager, got err: %v", err)
	}

	// Another fast call shouldn't block.
	conn2, _ := manager2.GetConnection()
	if conn2 != conn {
		t.Fatalf("sync.Once should return exactly the same pointer!")
	}
}
