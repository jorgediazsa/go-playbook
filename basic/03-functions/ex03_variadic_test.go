package functions

import "testing"

func TestSanitizeAndSum(t *testing.T) {
	callerSlice := []int{50, 150, 20}

	// Copy the original to verify it hasn't been corrupted
	original := make([]int, len(callerSlice))
	copy(original, callerSlice)

	// Caller passes slice using unpack operator
	sum := SanitizeAndSum(callerSlice...)

	expectedSum := 50 + 100 + 20 // 170
	if sum != expectedSum {
		t.Errorf("Expected sum %d, got %d", expectedSum, sum)
	}

	// Verify the caller's slice was NOT mutated
	for i, v := range callerSlice {
		if v != original[i] {
			t.Fatalf("CRITICAL: Variadic function mutated the caller's backing array! \nIndex %d changed from %d to %d", i, original[i], v)
		}
	}
}
