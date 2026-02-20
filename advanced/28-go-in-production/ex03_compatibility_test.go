package production

import (
	"strings"
	"testing"
)

func TestAPICompatibility(t *testing.T) {
	dbUser := DBUser{FirstName: "Ada", LastName: "Lovelace"}

	payload, err := FormatUser(dbUser)
	if err != nil {
		t.Fatalf("FAILED: FormatUser error: %v", err)
	}

	jsonStr := string(payload)

	// Test Backward Compatibility (Old Mobile Apps)
	if !strings.Contains(jsonStr, `"name":"Ada Lovelace"`) && !strings.Contains(jsonStr, `"name": "Ada Lovelace"`) {
		t.Fatalf("FAILED (BREAKING CHANGE): Old clients expecting the 'name' field will crash! You must include the concatenated 'name' field in your JSON response to maintain backward compatibility. Output: %v", jsonStr)
	}

	// Test Forward Evolution (New Web Apps)
	if !strings.Contains(jsonStr, `"first_name":"Ada"`) {
		t.Fatalf("FAILED: Missing new struct fields!")
	}
}
