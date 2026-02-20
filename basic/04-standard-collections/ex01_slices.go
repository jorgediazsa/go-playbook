package collections

// Context: Slice aliasing and reallocation
// You are writing an auditing function that appends an "AUDIT" tag to a list of log entries.
// The list of entries is passed in from various parts of the codebase.
//
// Why this matters: Slices are passed by value (the 24-byte header is copied).
// If `append` doesn't exceed the capacity, it modifies the existing backing array (meaning
// the caller sees the new element if they change their slice length, or it overwrites data).
// If `append` DOES exceed capacity, it allocates a new array. The caller never sees the added element.
// Both scenarios lead to severe bugs.
//
// Requirements:
// 1. the `AppendAuditLog` function currently accepts a slice and tries to append to it.
// 2. Fix the function signature and implementation so it safely returns the updated slice.
// 3. Guarantee that appending the audit log does NOT overwrite the caller's future capacity
//    if the caller had a slice with `len < cap`. (Hint: Force reallocation or return a clean copy).

func AppendAuditLog(logs []string) []string {
	// BUG: If len(logs) < cap(logs), appending here overwrites the memory adjacent to the slice,
	// which the caller might be relying on.
	// If it reallocates, the caller never sees the change unless we return it.

	// TODO: Create a safe copy of the slice, append "AUDIT_ENTRY", and return it.
	// Do not mutate the backing array passed in!

	logs = append(logs, "AUDIT_ENTRY") // This mutation is unsafe.
	return logs
}
