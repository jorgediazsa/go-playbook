package reflection

import (
	"testing"
)

type TargetUser struct {
	ID         int
	Name       string
	unexported bool // Reflection cannot set this!
}

func TestCopyMapToStruct(t *testing.T) {
	src := map[string]any{
		"ID":         42,
		"Name":       "Deepmind",
		"unexported": true,
		"FakeField":  "ignore me",
	}

	dst := TargetUser{}

	// Test 1: Fail if not a pointer
	if err := CopyMapToStruct(src, dst); err == nil {
		t.Fatalf("FAILED: Expected error when destination is passed by value")
	}

	// Test 2: Happy Path
	if err := CopyMapToStruct(src, &dst); err != nil {
		t.Fatalf("FAILED: Expected success, got error: %v", err)
	}

	if dst.ID != 42 || dst.Name != "Deepmind" {
		t.Fatalf("FAILED: Copy failed. Expected {42, Deepmind}, got %v", dst)
	}

	if dst.unexported == true {
		t.Fatalf("FAILED: You illegally bypassed CanSet() and altered an unexported field!")
	}
}
