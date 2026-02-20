package filesystemos

import (
	"context"
	"testing"
	"time"
)

func TestGracefulShutdownSignal(t *testing.T) {
	// Usually, we would do:
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	// defer cancel()
	// But sending OS signals to unit-test sub-processes is wildly flaky across
	// Mac/Linux/Windows. Because Go flawlessly unifies Context and OS Signals,
	// we can perfectly simulated a SIGTERM by simply cancelling a background context!

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan string)

	// 1. Start the server
	go func() {
		res := RunServer(ctx)
		done <- res
	}()

	// 2. Wait 200 ms to let it process a few ticks.
	time.Sleep(200 * time.Millisecond)

	// 3. SIMULATE SIGTERM: Send the shutdown signal
	cancel()

	// 4. Did the server break the loop gracefully?
	select {
	case res := <-done:
		// Success! The server noticed the signal and exited the loop.
		if res == "" {
			t.Fatalf("Server returned empty string, expected graceful msg")
		}
	case <-time.After(1 * time.Second):
		// FAILED! The server ignored the signal and kept running.
		t.Fatalf("FAILED: A SIGTERM signal was sent (ctx cancelled), but the server kept looping forever. You must select on ctx.Done()!")
	}
}
