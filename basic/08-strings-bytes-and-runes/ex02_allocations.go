package stringsbytes

import (
	"strconv"
)

// Context: Allocation Costs
// You are generating a massive CSV row from a million data points.
//
// Why this matters: Strings in Go are immutable.
// Using `+=` inside a loop creates a new string every time, copying the old string
// and adding the new piece. This is an O(N^2) operation that causes massive CPU
// and GC overhead in production.
//
// Requirements:
// 1. Refactor `GenerateCSVRow` to execute without doing `+=` string concatenations.
// 2. Use `strings.Builder`.
// 3. (Optional but recommended) Estimate the capacity of the builder upfront.

func GenerateCSVRow(data []int) string {
	// BUG: High allocation loop
	// TODO: Replace with strings.Builder

	row := ""
	for i, val := range data {
		row += strconv.Itoa(val)
		if i < len(data)-1 {
			row += ","
		}
	}

	return row
}
