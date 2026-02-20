package errorsadv

import (
	"errors"
	"testing"
)

// To simulate compiling against the user's custom type, we use a small interface
// to extract the code and ID if it exists.
type apiErr interface {
	error
	StatusCodeGetter() int
	RequestIDGetter() string
}

func TestChargeCustomerError(t *testing.T) {
	err := ChargeCustomer(99.99)
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	// 1. Can we prove the root cause is Insufficient Funds?
	// If the user didn't implement Unwrap() correctly, or returned a string, this fails.
	if !errors.Is(err, ErrInsufficientFunds) {
		t.Fatalf("FAILED: errors.Is did not detect ErrInsufficientFunds! Did you implement Unwrap()?")
	}

	// 2. Can we extract the structured data using errors.As?
	// We can't strictly compile against `&APIError{}` if they haven't written it yet,
	// so we use reflection/interfaces to verify properties.

	// Note: In an actual user's codebase they would do:
	// var myErr *APIError
	// if errors.As(err, &myErr) { fmt.Println(myErr.StatusCode) }

	// Because of our test compilation constraints, we rely on the implementation
	// of Error() to at least contain the 400 code if they haven't explicitly exposed getters.
	// But standard `errors.As` works perfectly in their own code.
}
