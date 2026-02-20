package control

// Context: Range Semantics (Map non-determinism & String decoding)
// You have a map of priority user IDs that need to be grouped into batches.
// You also have an encoded priority string that contains multi-byte characters.
//
// Why this matters:
// 1. Map iteration order in Go is intentionally randomized to prevent developers
//    from depending on insertion order (which hash maps don't guarantee).
// 2. Iterating a string with a classic `for i := 0; i < len(n); i++` loops over BYTES,
//    not RUNES (characters). `for i, c := range s` loops over runes.
//
// Requirements:
// 1. Return the userIDs in a *deterministic*, alphabetically sorted order.
// 2. Count the number of actual characters (runes) in the string, not bytes.

func GetSortedVIPs(vips map[string]bool) []string {
	var result []string

	// BUG: Direct range over map is non-deterministic.
	// TODO: Fix this to always return a sorted slice of keys.
	for k := range vips {
		result = append(result, k)
	}

	return result
}

func CountCharacters(encoded string) int {
	// BUG: len(string) returns the number of bytes, not characters.
	// A string like "こんにちは" has 5 characters but 15 bytes.
	// TODO: Use a range loop (or utf8 package) to correctly return character count.
	return len(encoded)
}
