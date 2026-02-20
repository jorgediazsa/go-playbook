package composition

import (
	"bytes"
	"testing"
)

func TestQuotaWriter(t *testing.T) {
	// A standard bytes.Buffer implements io.Writer.
	var underlying bytes.Buffer

	// We wrap it in our decorator.
	// If the user's struct doesn't implement Write, this test won't compile.
	qw := NewQuotaWriter(&underlying, 10)

	// Write 1 (Valid)
	n, err := qw.Write([]byte("Hello")) // 5 bytes
	if err != nil {
		t.Fatalf("Unexpected error on valid write: %v", err)
	}
	if n != 5 {
		t.Fatalf("Expected 5 bytes written, got %d", n)
	}
	if qw.BytesWritten != 5 {
		t.Fatalf("Decorator state not updated. Expected 5, got %d", qw.BytesWritten)
	}

	// Write 2 (Exceeds limit)
	n, err = qw.Write([]byte(" World!")) // 7 bytes. Total would be 12 > 10.

	if err != ErrQuotaExceeded {
		t.Fatalf("Expected ErrQuotaExceeded, got %v", err)
	}

	// Did it write anyway through the wrapped writer?
	if underlying.String() != "Hello" {
		t.Fatalf("CRITICAL: The decorator leaked the payload to the underlying writer despite exceeding quota! Wrote: %q", underlying.String())
	}

	// Ensure the state didn't increment on a failed write
	if qw.BytesWritten != 5 {
		t.Fatalf("Decorator state corrupted on failed write. Expected 5, got %d", qw.BytesWritten)
	}
}
