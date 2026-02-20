// BUG: Missing the correct build tag.
// TODO: Add the build tag so this file is ONLY compiled when the `fast` tag is provided.
// Example: `//go:build fast`

package tags

// Context: Build Tags (Optimized Version)
// See hash_safe.go for instructions.

// This file fails to compile concurrently with hash_safe.go because it re-declares Hash()
// unless the build tags make them mutually exclusive.

func Hash(data string) int {
	// A mock fast, assembly-optimized hashing algorithm
	// (Returns a predictable number for testing)
	return 9999
}
