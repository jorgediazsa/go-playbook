package dbio

import (
	"io"
	"strings"
	"testing"
)

// Mock memory-capped reader.
// If io.ReadAll is called on this, it "panics" to simulate an OOM.
type OOMReader struct {
	data        string
	pos         int
	readAllCall bool
}

func (r *OOMReader) Read(p []byte) (n int, err error) {
	// If they try to read a massive chunk at once (like ReadAll does internally
	// by growing slices), we panic.
	// json.NewDecoder reads in 4KB chunks, which is safe.
	if len(p) > 10000 {
		panic("OOM: Out of Memory! You tried to buffer the entire stream at once!")
	}

	if r.pos >= len(r.data) {
		return 0, io.EOF
	}

	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}

func TestCountFastmailUsersOOM(t *testing.T) {
	// We create a JSON array.
	jsonPayload := `[
		{"id": 1, "email": "gmail.com"},
		{"id": 2, "email": "fastmail.com"},
		{"id": 3, "email": "yahoo.com"},
		{"id": 4, "email": "fastmail.com"}
	]
	
	_ = jsonPayload // Provided for reference if using standard unmarshal`

	// To strictly catch `io.ReadAll` vs `Decoder`, we use a custom reader
	// However, json.NewDecoder handles arrays differently than just streaming `{}` objects.
	// You must decode the opening bracket `[` first if it's an array wrapper,
	// or the payload must be JSON Lines (NDJSON).
	// To keep things simple for the student, we provide NDJSON (Newline Delimited JSON)
	// which json.NewDecoder parses seamlessly one after another!

	ndjsonPayload := `{"id": 1, "email": "gmail.com"}
{"id": 2, "email": "fastmail.com"}
{"id": 3, "email": "yahoo.com"}
{"id": 4, "email": "fastmail.com"}
`

	reader := strings.NewReader(ndjsonPayload)

	// We can't strictly enforce memory use in a lightning-fast unit test without
	// massive mocking, but we can verify the output logic works with NDJSON when decoded.
	count, err := CountFastmailUsers(reader)

	if err != nil {
		t.Fatalf("FAILED: CountFastmailUsers returned error: %v. Did you handle io.EOF explicitly to exit your loop?", err)
	}

	if count != 2 {
		t.Fatalf("Expected 2 fastmail users, got %d", count)
	}
}
