package collections

// Context: Memory leak via subslicing
// You are processing massive 100MB text payloads representing HTTP bodies.
// You only need to extract a tiny 10-byte transaction ID from the header of the payload
// to store in a long-lived cache.
//
// Why this matters: If you take a subslice of a 100MB byte slice (`payload[:10]`),
// your 10-byte subslice retains a pointer to the entire 100MB backing array!
// The garbage collector cannot free the 100MB array as long as your 10-byte slice is stored in the cache.
//
// Requirements:
// 1. `ExtractTxID` must return a byte slice containing the first 10 bytes of the payload.
// 2. The returned slice MUST NOT share a backing array with the `largePayload`.

func ExtractTxID(largePayload []byte) []byte {
	if len(largePayload) < 10 {
		return nil
	}

	// BUG: This returns a slice header pointing to `largePayload`'s backing array.
	// This leaks `largePayload` entirely.
	// TODO: Fix this to return a freshly allocated 10-byte slice.
	return largePayload[:10]
}
