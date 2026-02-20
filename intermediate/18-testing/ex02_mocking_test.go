package testingadv

import "testing"

// TODO: Create a Mock struct here that implements your new interface!
// It should return the phone number "+19999999" for any user,
// and record the Number of times it was called in a struct field.
type MockDB struct {
	// Add fields here
}

// TODO: Implement the required method(s) on MockDB

func TestSendUrgentAlertWithMock(t *testing.T) {
	// BUG: We instantiate the real DB here, which will supposedly fail in CI!
	// TODO: Replace this with your Mock struct.
	// db := &LegacyUserDB{}

	var db *MockDB // Ensure you instantiate this correctly once implemented

	if db == nil {
		t.Fatal("FAILED: You must replace LegacyUserDB with your MockDB struct!")
	}

	// If you didn't refactor the signature of SendUrgentAlert, passing a mock
	// here will fail to compile.
	err := SendUrgentAlert(db, "usr_xyz", "Server is down!")

	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}

	// Verify the mock was actually called.
	// if db.Calls != 1 { ... t.Fatalf("Mock wasn't called!") }
}
