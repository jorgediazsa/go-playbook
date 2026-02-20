package pointers

import "sync"

// Context: Receivers and Copied Mutexes
// You are building an in-memory datastore representing an organization's bank accounts.
// The datastore must be concurrency-safe, so it embeds a sync.Mutex.
//
// Why this matters: Using a value receiver on a struct that contains a `sync.Mutex`
// creates a COPY of the mutex on every method call. This completely neutralizes
// the lock, leading to catastrophic race conditions. Furthermore, any mutations
// applied to the struct inside the method will be discarded when the method returns,
// since they were applied to the copy!
//
// Requirements:
// 1. `Withdraw` must subtract the amount from the balance, entirely safely.
// 2. `Deposit` must add the amount to the balance, entirely safely.
// 3. Fix the receiver types on both methods.

type Account struct {
	mu      sync.Mutex
	Balance int
}

// BUG: Value receiver copies the struct (and the mutex).
// Any lock acquired here only locks the copy. Any mutation to Balance only mutates the copy.
// TODO: Fix the receiver type so mutations apply to the actual account, and locks work.
func (a Account) Withdraw(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.Balance >= amount {
		a.Balance -= amount
	}
}

// BUG: Same as above.
// TODO: Fix the receiver.
func (a Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance += amount
}
