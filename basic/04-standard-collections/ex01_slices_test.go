package collections

import "testing"

func TestAppendAuditLog(t *testing.T) {
	// We simulate a caller who created a slice with extra capacity for their own use.
	callerArray := make([]string, 2, 5)
	callerArray[0] = "Log1"
	callerArray[1] = "Log2"

	// The caller plans to put something at index 2 eventually.
	callerArray = append(callerArray, "CallerFutureLog")

	// Revert the length back to 2 to simulate passing the slice before they appended.
	passedSlice := callerArray[:2]

	result := AppendAuditLog(passedSlice)

	// Did the function return the expected result?
	if len(result) != 3 || result[2] != "AUDIT_ENTRY" {
		t.Fatalf("Expected result to have AUDIT_ENTRY at index 2, got %v", result)
	}

	// The critical test: Did `AppendAuditLog` overwrite `callerArray[2]` ?
	// If `AppendAuditLog` mutated the backing array, `callerArray[2]` is now "AUDIT_ENTRY"
	// instead of "CallerFutureLog".
	if callerArray[2] != "CallerFutureLog" {
		t.Fatalf("CRITICAL SECURITY BUG: AppendAuditLog overwrote the caller's backing array! callerArray[2] is now %q", callerArray[2])
	}
}
