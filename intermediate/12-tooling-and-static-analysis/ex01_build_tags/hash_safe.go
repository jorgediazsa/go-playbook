//go:build !fast

package tags

// Context: Build Tags (Safe Fallback)
// You have an optimized assembly algorithm for hashing, but it only runs on
// specific OS/Arch combinations or when explicitly requested via a build flag.
//
// Why this matters: Build tags (`//go:build tagname`) allow the Go compiler
// to include or exclude files at compile time. This is how the standard library
// supports Windows, Mac, and Linux seamlessly.
//
// Requirements:
// 1. You have two files: `hash_safe.go` and `hash_fast.go`.
// 2. Add the correct build tags so that:
//    - `hash_safe.go` is compiled BY DEFAULT (when no tags are passed).
//    - `hash_fast.go` is compiled ONLY when `go build -tags fast` is run.
// 3. Make sure they both declare the same function signature for `Hash(data string) int`.
//
// Note: The build tag must be at the VERY top of the file, followed by a blank line.
//       This file currently has `//go:build !fast`.

func Hash(data string) int {
	// A mock slow, safe hashing algorithm
	sum := 0
	for _, char := range data {
		sum += int(char)
	}
	return sum
}
