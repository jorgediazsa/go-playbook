package production

import (
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestGracefulShutdown(t *testing.T) {
	sigCh := make(chan os.Signal, 1)

	errCh := make(chan error, 1)

	// Since they must start ListenAndServe in a background goroutine internally,
	// RunServer should now block on the sigCh.
	go func() {
		errCh <- RunServer("localhost:9191", sigCh)
	}()

	// Wait for server to boot
	time.Sleep(100 * time.Millisecond)

	// Fire an active request that takes 2 seconds to finish.
	reqStarted := time.Now()
	go http.Get("http://localhost:9191/")

	time.Sleep(100 * time.Millisecond)

	// Send the shutdown signal!
	sigCh <- syscall.SIGTERM

	// Wait for RunServer to return.
	// If graceful shutdown is implemented, `Shutdown` will pause and wait
	// for the 2-second HTTP request above to finish before returning.

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("FAILED: Server returned error: %v", err)
		}

		runDuration := time.Since(reqStarted)

		// If duration < 2s, the server dropped the connection immediately.
		// If duration > 2s, the server gracefully waited for the connection to finish!
		if runDuration < 1900*time.Millisecond {
			t.Fatalf("FAILED: Server exited too fast (%v). It dropped active connections instead of calling Shutdown()!", runDuration)
		}

	case <-time.After(6 * time.Second):
		t.Fatalf("FAILED: Server did not shut down or ignored the SIGTERM!")
	}
}
