package dbio

import (
	"context"
	"errors"
	"testing"
)

// MockTx implements our Tx interface
type MockTx struct {
	RollbackCalled bool
	CommitCalled   bool
	ShouldFailExec bool
}

func (m *MockTx) Exec(query string, args ...any) (any, error) {
	if m.ShouldFailExec && query == "INSERT INTO orders (user, item) VALUES (?, ?)" {
		return nil, errors.New("db crash")
	}
	return nil, nil
}
func (m *MockTx) Commit() error   { m.CommitCalled = true; return nil }
func (m *MockTx) Rollback() error { m.RollbackCalled = true; return nil }

type MockPool struct {
	Tx *MockTx
}

func (p *MockPool) BeginTx(ctx context.Context) (Tx, error) {
	return p.Tx, nil
}

func TestCheckoutTransactionSafety(t *testing.T) {
	// Test 1: Happy Path
	mockTx := &MockTx{}
	pool := &MockPool{Tx: mockTx}

	err := Checkout(context.Background(), pool, "item1", "user1")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if !mockTx.CommitCalled {
		t.Fatalf("FAILED: Happy path did not call Commit()!")
	}

	// Test 2: Failure Path (Rollback required!)
	mockTxFail := &MockTx{ShouldFailExec: true}
	poolFail := &MockPool{Tx: mockTxFail}

	errFail := Checkout(context.Background(), poolFail, "item1", "user1")
	if errFail == nil {
		t.Fatal("Expected an error from checkout, got nil")
	}

	if !mockTxFail.RollbackCalled {
		t.Fatalf("CRITICAL: The transaction failed, but Rollback() was never called! You leaked a database connection! Did you use `defer tx.Rollback()`?")
	}
	if mockTxFail.CommitCalled {
		t.Fatalf("FAILED: You called Commit() even though an error occurred!")
	}
}
