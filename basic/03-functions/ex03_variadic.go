package functions

// Context: Variadic Functions and Slice Unpacking
// You are building a metrics aggregation system.
// Other teams pass dynamically sized slices of latency metrics to your `Record` function.
//
// Why this matters: When passing an existing slice to a variadic function using the
// spread operator (`Record(mySlice...)`), Go does NOT copy the backing array.
// If your variadic function modifies the slice (e.g., sorting it, capping values, or
// appending which might or might not reallocate), you corrupt the caller's slice!
//
// Requirements:
// 1. `SanitizeAndSum` takes variadic integers, caps any value > 100 to 100, and returns the sum.
// 2. Ensure that `SanitizeAndSum` does NOT permanently alter the slice passed in by the caller.

func SanitizeAndSum(metrics ...int) int {
	// BUG:
	// We are modifying the `metrics` slice directly.
	// If the caller used `...` to pass an existing slice, this iterates over their
	// backing array and modifies their memory!
	// TODO: Prevent leaking mutations back to the caller.

	sum := 0
	for i, m := range metrics {
		if m > 100 {
			metrics[i] = 100
		}
		sum += metrics[i]
	}

	return sum
}
