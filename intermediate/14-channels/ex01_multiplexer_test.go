package channels

import (
	"runtime"
	"testing"
	"time"
)

func TestSearchFastest(t *testing.T) {
	baselineRoutines := runtime.NumGoroutine()

	// Provider 1: Slow (100ms)
	p1 := func(q string) string {
		time.Sleep(100 * time.Millisecond)
		return "P1"
	}

	// Provider 2: Fast (10ms)
	p2 := func(q string) string {
		time.Sleep(10 * time.Millisecond)
		return "P2"
	}

	// Provider 3: Normal (30ms)
	p3 := func(q string) string {
		time.Sleep(30 * time.Millisecond)
		return "P3"
	}

	start := time.Now()
	res := SearchFastest("query", p1, p2, p3)
	duration := time.Since(start)

	if res != "P2" {
		t.Fatalf("Expected the fastest result 'P2', got '%s'", res)
	}

	if duration > 50*time.Millisecond {
		t.Fatalf("Failed to run concurrently! Expected < 50ms, took %v", duration)
	}

	// Wait for the slow ones to "finish".
	time.Sleep(150 * time.Millisecond)

	if runtime.NumGoroutine() > baselineRoutines {
		t.Fatalf("LEAK DETECTED: You leaked the slower goroutines. Did you use an unbuffered channel that they got stuck sending to after the function returned?")
	}

	// Test Timeout Logic
	slowProvider := func(q string) string {
		time.Sleep(100 * time.Millisecond)
		return "SLOW"
	}

	res = SearchFastest("query", slowProvider, slowProvider, slowProvider)
	if res != "timeout" {
		t.Fatalf("Expected 'timeout', got '%s'", res)
	}
}
