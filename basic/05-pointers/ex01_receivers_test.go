package pointers

import (
	"sync"
	"testing"
)

func TestAccountMutationAndSafety(t *testing.T) {
	// Let's create an account with $1000.
	// We MUST pass it by pointer, or else the first method call operates on a copy.
	acc := &Account{Balance: 1000}

	// 1. Check basic mutation (did the receiver bug get fixed?)
	acc.Withdraw(100)
	if acc.Balance != 900 {
		t.Fatalf("Expected balance 900, got %d. Mutation failed. Did you change the receiver from 'a Account' to 'a *Account'?", acc.Balance)
	}

	acc.Deposit(200)
	if acc.Balance != 1100 {
		t.Fatalf("Expected balance 1100, got %d. Mutation failed.", acc.Balance)
	}

	// 2. Check concurrency safety (did the mutex copy bug get fixed?)
	// If the user left `a Account`, the race detector will scream here,
	// or the final balance will be wildly off.
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			acc.Deposit(1)
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			acc.Withdraw(1)
		}()
	}

	wg.Wait()

	if acc.Balance != 1100 {
		t.Fatalf("Expected final balance to be 1100 (net zero change), got %d. Concurrency bug exists.", acc.Balance)
	}

	// Ensure the program is run with -race!
}
