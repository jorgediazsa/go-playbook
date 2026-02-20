package errorsadv

// Context: Error Classification (Retriable vs Fatal)
// You have a generic retry loop that takes a function and runs it 3 times.
// But some errors are FATAL (e.g., "invalid password", 401 Unauthorized), meaning
// retrying them is utterly pointless. Others are TEMPORARY (e.g., "server down",
// 503 Unavailable), meaning a retry might work.
//
// Why this matters: You shouldn't hardcode `if err == ErrServerDown` inside your
// generic retry logic, because then that package has to know about every possible
// error type from every other package.
// Idiomatic Go uses small interfaces to ask the error about its *behavior*, not its *identity*.
//
// Requirements:
// 1. Define a `Temporary` interface with a single method `Temporary() bool`.
// 2. Refactor `ExecuteWithRetry` to check if the error implements `Temporary`.
//    - If it DOES implement `Temporary` and returns `true`, continue the retry loop.
//    - If it DOES NOT implement `Temporary` (or returns `false`), return the fatal error immediately.

import (
	"errors"
	"time"
)

var ErrInvalidCredentials = errors.New("invalid credentials, do not retry")
var ErrNetworkTimeout = errors.New("network timeout, please retry")

// TODO: Define the Temporary interface

func ExecuteWithRetry(operation func() error) error {
	// BUG: This retries EVERYTHING, even if the password was wrong!
	// TODO: Use `errors.As` or type-assertion to check if the error
	// implements your `Temporary()` interface. If not, return it immediately.

	for i := 0; i < 3; i++ {
		err := operation()
		if err == nil {
			return nil
		}

		// If the error is fatal, return err immediately!

		time.Sleep(10 * time.Millisecond)
	}
	return errors.New("max retries exceeded")
}
