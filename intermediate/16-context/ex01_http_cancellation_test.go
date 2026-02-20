package ctxexercises

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGenerateReportHandlerCancellation(t *testing.T) {
	// 1. Setup a request with a context that we can MANUALLY cancel early.
	ctx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequestWithContext(ctx, "GET", "/report", nil)
	rr := httptest.NewRecorder()

	// 2. We run the handler in a goroutine
	done := make(chan struct{})
	go func() {
		GenerateReportHandler(rr, req)
		close(done)
	}()

	// 3. The user instantly disconnects (cancels) after 100ms!
	time.Sleep(100 * time.Millisecond)
	cancel()

	// 4. Wait for the handler to finish
	start := time.Now()
	<-done
	duration := time.Since(start)

	// If they didn't fix `MockDatabaseQuery` or didn't pass `r.Context()`,
	// this will have taken the full 3 seconds.
	if duration > 1*time.Second {
		t.Fatalf("FAILED: The DB query did not abort! It took %v. You must select on ctx.Done()!", duration)
	}

	if rr.Code != 499 {
		t.Fatalf("Expected HTTP 499 (Client Closed), got %d", rr.Code)
	}
}
