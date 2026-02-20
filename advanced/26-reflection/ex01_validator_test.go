package reflection

import (
	"testing"
)

type ValidUser struct {
	ID   int    `validate:"required"`
	Name string `validate:"required"`
	Age  int    // Optional
}

func TestValidateStruct(t *testing.T) {
	// Test 1: Happy Path
	user1 := ValidUser{ID: 1, Name: "Alice", Age: 0}
	if err := ValidateStruct(user1); err != nil {
		t.Fatalf("FAILED: Expected valid user to pass, got error: %v", err)
	}

	// Test 2: Missing Int
	user2 := ValidUser{ID: 0, Name: "Bob", Age: 30}
	if err := ValidateStruct(user2); err == nil {
		t.Fatalf("FAILED: Expected error because ID is 0 and required.")
	}

	// Test 3: Missing String
	user3 := ValidUser{ID: 3, Name: "", Age: 40}
	if err := ValidateStruct(user3); err == nil {
		t.Fatalf("FAILED: Expected error because Name is empty and required.")
	}

	// Test 4: Passed a Pointer
	user4 := &ValidUser{ID: 4, Name: "Charlie"}
	if err := ValidateStruct(user4); err != nil {
		t.Fatalf("FAILED: Must handle pointers to structs too! Error: %v", err)
	}

	// Test 5: Not a Struct
	if err := ValidateStruct("just a string"); err == nil {
		t.Fatalf("FAILED: Expected error when passing a non-struct")
	}
}
