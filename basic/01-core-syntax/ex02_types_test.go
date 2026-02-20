package core

import (
	"reflect"
	"testing"
)

func TestTypeAliasRefactor(t *testing.T) {
	// Create sample types
	var u UserID = "user-123"
	var p ProductID = "prod-abc"

	uType := reflect.TypeOf(u)
	pType := reflect.TypeOf(p)
	strType := reflect.TypeOf("")

	// The core requirement: UserID and ProductID must NOT be identical to each other or to string.
	// In the initial buggy version, all three types are identical string types.
	if uType == pType {
		t.Errorf("Security Flaw: UserID and ProductID evaluate to the exact same type. Did you forget to remove the '=' in the type definition?")
	}

	if uType == strType {
		t.Errorf("Security Flaw: UserID is just an alias for string.")
	}

	result := ProcessOrder(u, p)
	expected := "Order: user-123 bought prod-abc"
	if result != expected {
		t.Errorf("ProcessOrder() = %v, want %v", result, expected)
	}

	userData := FetchUser(u)
	expectedData := "Data for user-123"
	if userData != expectedData {
		t.Errorf("FetchUser() = %v, want %v", userData, expectedData)
	}
}
