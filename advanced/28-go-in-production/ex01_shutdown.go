package production

import (
	"errors"
	"net/http"
	"os"
	"time"
)

// Context: Graceful Server Shutdown
// You are deploying an API to Kubernetes. When K8s rolls out a new version,
// it sends a SIGTERM to the old pod.
//
// Why this matters: If your `RunServer` function immediately exits the program
// when it receives a signal, any user currently uploading a photo will have
// their TCP connection violently severed.
//
// Requirements:
// 1. Refactor `RunServer` to start `server.ListenAndServe()` in a background goroutine.
// 2. The main function should block waiting for a signal on the `sigCh`.
// 3. When a signal arrives, create a `context.WithTimeout` (e.g., 5 seconds).
// 4. Call `server.Shutdown(ctx)`. This tells Go to stop accepting new requests,
//    wait for active requests to finish, and return an error if the 5s timeout hits.

func RunServer(addr string, sigCh <-chan os.Signal) error {
	server := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Simulate a slow upload
			time.Sleep(2 * time.Second)
			w.WriteHeader(http.StatusOK)
		}),
	}

	// BUG: This blocks forever, ignoring the `sigCh`, meaning the app never shuts down!
	// And if K8s kills it forcefully later, active connections drop.
	// TODO: Move ListenAndServe to a goroutine.
	// TODO: Block on `<-sigCh`.
	// TODO: Call `server.Shutdown(ctx)` with a 5-second timeout.

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
