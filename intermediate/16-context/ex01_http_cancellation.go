package ctxexercises

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Context: HTTP Cancellation
// You are building an expensive reporting endpoint (`/report`). Generating
// the report takes 3 seconds of heavy DB compute.
//
// Why this matters: If a user clicks "Generate 10 times" and closes the browser 9
// times, your server is currently computing 10 massive reports in the background
// wasting resources, even though 9 users have already disconnected!
// Every `http.Request` gives you a Context that cancels when the client disconnects.
//
// Requirements:
// 1. Refactor `GenerateReportHandler` to use the request's context.
// 2. You must pass that context into the `MockDatabaseQuery`.
// 3. Inside `MockDatabaseQuery`, use a `select` statement to either:
//    - Return the result after 3 simul-seconds.
//    - Return exactly `ctx.Err()` immediately if the context cancels.

func MockDatabaseQuery(ctx context.Context) (string, error) {
	// BUG: This function ignores the context. It burns CPU for 3 seconds guarantees.
	// TODO: Replace this sleep with a `select` on `time.After` and `ctx.Done()`.

	time.Sleep(3 * time.Second)

	return "Massive Report Data", nil
}

func GenerateReportHandler(w http.ResponseWriter, r *http.Request) {
	// BUG: We are passing context.Background() instead of the request context!
	// TODO: Capture the request context and pass it down.
	ctx := context.Background()

	result, err := MockDatabaseQuery(ctx)
	if err != nil {
		// If it was canceled, return a 499 (Client Closed Request) or just let it drop.
		http.Error(w, err.Error(), 499)
		return
	}

	fmt.Fprint(w, result)
}
