package errorsadv

import (
	"errors"
	"strings"
	"testing"
)

func TestValidatePayloadAggregate(t *testing.T) {
	// Send a completely blank payload
	err := ValidatePayload("", "short")

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	// Because `errors.Join` returns an unexported struct that iterates over the slice
	// of internal errors, `errors.Is` should mathematically prove BOTH errors happened at once!

	hasMissingEmail := errors.Is(err, ErrMissingEmail)
	hasShortPassword := errors.Is(err, ErrPasswordShort)

	if !hasMissingEmail || !hasShortPassword {
		t.Fatalf(
			"FAILED: Validation stopped on the first error!\nhasMissingEmail: %v\nhasShortPassword: %v\nAggregated Error Output: %v\n\nDid you use errors.Join?",
			hasMissingEmail, hasShortPassword, err,
		)
	}

	// Print exactly how `errors.Join` formats them to the console.
	// (Usually joining with newlines).
	if !strings.Contains(err.Error(), "email is required") || !strings.Contains(err.Error(), "password too short") {
		t.Fatalf("The error string did not contain both failure messages: %q", err.Error())
	}
}
