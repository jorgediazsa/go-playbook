package composition

import (
	"testing"
)

// MockClient is used to test OnboardUser without hitting real APIs.
type MockClient struct {
	SentTo string
}

// TODO: Ensure MockClient satisfies the EmailSender interface you defined in ex01_interfaces.go.
// Implement the required method(s) here.
// Simulate a failure if the email is "fail@test.com", otherwise record the email and return nil.

func TestOnboardUser(t *testing.T) {
	mock := &MockClient{}

	// If OnboardUser is still requiring the concrete `*ThirdPartyEmailClientImpl`,
	// this next line will not compile.

	// Test happy path
	err := OnboardUser(mock, "new@user.com")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if mock.SentTo != "new@user.com" {
		t.Errorf("Mock did not record the correct email. Got %q", mock.SentTo)
	}

	// Test failure path
	err = OnboardUser(mock, "fail@test.com")
	if err == nil {
		t.Fatal("Expected an error for the failure scenario, got nil")
	}
}
