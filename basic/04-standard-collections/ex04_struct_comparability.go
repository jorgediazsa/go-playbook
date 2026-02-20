package collections

// Context: Struct Comparability
// You are building a deduplication filter. You want to store incoming HTTP request
// signatures in a map to quickly check if you've seen them before in the last 5 seconds.
//
// Why this matters: Structs are only comparable (usable as map keys or with ==)
// if ALL of their fields are comparable. Slices, maps, and functions are NOT comparable.
// If you use a struct containing a slice/map as a map key, the compiler rejects it.
// If you use an interface that happens to hold a slice/map at runtime, it PANICS.
//
// Requirements:
// 1. `RequestSignature` currently contains a `[]string` for headers, making it incomparable.
// 2. Refactor `RequestSignature` so it IS comparable and can be used as a map key.
//    (Hint: How can you safely represent a list of strings in a comparable way?
//    Maybe a comma-separated string, or an array if the size is strictly fixed.
//    For this exercise, converting the headers slice to a single joined string is best).
// 3. Update `IsDuplicate` to compile and work.

// BUG: This struct cannot be used as a map key because `Headers` is a slice.
// TODO: Refactor this to be comparable.
type RequestSignature struct {
	Method  string
	Path    string
	Headers []string
}

// TODO: Fix the signature of this map and the logic below
// after you fix `RequestSignature`.
var seenRequests = make(map[interface{}]bool) // using interface{} to mock the failure

func IsDuplicate(req RequestSignature) bool {
	// BUG: If seenRequests was map[RequestSignature]bool, this wouldn't compile
	// because RequestSignature is not comparable.
	// Since we mocked it with map[interface{}]bool, this will actually panic at runtime
	// when passing an incomparable struct!

	// TODO: After fixing `RequestSignature` to be comparable, change `seenRequests` to
	// `map[RequestSignature]bool` and fix this logic.

	if seenRequests[req] {
		return true
	}
	seenRequests[req] = true
	return false
}

func ResetCache() {
	// seenRequests = make(map[RequestSignature]bool)
	seenRequests = make(map[interface{}]bool)
}
