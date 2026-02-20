package errorsadv

import (
	"errors"
	"fmt"
)

// Context: Custom Error Types and Wrapping
// You are building an HTTP client for a billing API. When a request fails,
// the API returns a JSON error containing a specific HTTP Status Code and
// a Request ID for debugging.
//
// Why this matters: Returning `fmt.Errorf("API failed: %d, ID: %s", code, id)`
// destroys the machine-readability of the error. The caller can't extract the
// Status Code to decide if it should retry without doing horrific string parsing.
//
// Requirements:
// 1. Define a custom struct `APIError` that implements the `error` interface.
// 2. It must hold the `StatusCode` (int), `RequestID` (string), and `Err` (the underlying error).
// 3. Implement the `Unwrap() error` method on `APIError` so it returns the underlying `Err`.
// 4. Refactor `ChargeCustomer` to return your `APIError` wrapped around `ErrInsufficientFunds`.

var ErrInsufficientFunds = errors.New("insufficient funds")

// TODO: Define APIError here. Remember to implement `Error() string` and `Unwrap() error`.

func ChargeCustomer(amount float64) error {
	// BUG: We are returning an opaque string error. The caller can't use `errors.As`
	// to get the Status Code, and `errors.Is` will fail to detect ErrInsufficientFunds.

	// TODO: Return an *APIError with StatusCode 400, RequestID "req-123",
	// wrapping the ErrInsufficientFunds.

	return fmt.Errorf("API failed with status 400 (req-123): %w", ErrInsufficientFunds)
}
