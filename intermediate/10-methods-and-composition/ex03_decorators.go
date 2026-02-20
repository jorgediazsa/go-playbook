package composition

// Title: Stateful Writer Decorator
//
// Context: You are designing an HTTP proxy that forwards large payload bodies
// to a filesystem. But before writing to the file, security requirements mandate
// that we count every byte written to ensure the user doesn't exceed their 1MB quota.
//
// Why this matters: Instead of building a "QuotaAwareFileStruct", idiomatic Go uses
// decorators wrapping standard interfaces like `io.Writer`. A decorator accepts an `io.Writer`,
// implements the `io.Writer` interface itself, does its logic, and forwards the data.
// It can be chained infinitely.
//
// Requirements:
// 1. Define a struct `QuotaWriter` that wraps an `io.Writer`.
// 2. It must track the `BytesWritten`, and hold a `Limit`.
// 3. Implement the `Write(p []byte) (n int, err error)` method on `QuotaWriter`:
//    - If the incoming payload + BytesWritten exceeds `Limit`, return an `ErrQuotaExceeded`
//      and DO NOT pass the data to the underlying writer at all.
//    - Otherwise, pass the data to the underlying writer, add the written bytes to `BytesWritten`,
//      and return the result.

import (
	"errors"
	"io"
)

var ErrQuotaExceeded = errors.New("quota exceeded")

// TODO: Define the QuotaWriter struct.
// It must hold the underlying writer, the limit, and the current bytes written.

// TODO: Implement the io.Writer interface on QuotaWriter.
// func (qw *QuotaWriter) Write(p []byte) (n int, err error) { ... }
// Hint: Be sure to use a pointer receiver so BytesWritten actually increments!

// NewQuotaWriter is a helper constructor.
func NewQuotaWriter(w io.Writer, limit int) *QuotaWriter {
	// TODO: Return a properly initialized QuotaWriter pointer.
	return nil
}
