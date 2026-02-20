package dbio

import (
	"context"
	"errors"
)

// Context: Transaction Management & Interfaces
// You are building an e-commerce checkout. You must deduct stock and create an order.
// Both must succeed, or both must fail (rollback).
//
// Why this matters: `database/sql`'s `*sql.Tx` struct is concrete, making it hard
// to mock cleanly without huge libraries. By defining an interface representing
// your "Unit of Work," you decouple your business logic from `database/sql` entirely.
// Additionally, if a transaction is started, it MUST be rolled back if an error occurs.
//
// Requirements:
// 1. Define an interface `Tx` containing the methods `Exec`, `Commit`, and `Rollback`.
// 2. Refactor `Checkout` to accept this `Tx` interface instead of a raw struct.
// 3. Ensure that if ANY error occurs, `tx.Rollback()` is called BEFORE returning the error.
//    (Hint: The most bulletproof way is to `defer tx.Rollback()` as soon as the Tx begins).
// 4. Return `tx.Commit()` at the end of success.

// TODO: Define the Tx interface.

// We mock a DB connection pool that just returns our interface.
type DBPool interface {
	BeginTx(ctx context.Context) (Tx, error)
}

func Checkout(ctx context.Context, db DBPool, itemID string, userID string) error {
	tx, err := db.BeginTx(ctx)
	if err != nil {
		return err
	}

	// BUG: If Exec fails, or if a panic happens, the transaction leaks!
	// TODO: Safely guarantee a rollback using `defer`.
	// (Note: `tx.Rollback()` safely returns an error if the transaction was already committed).

	_, err = tx.Exec("UPDATE stock SET count = count - 1 WHERE id = ?", itemID)
	if err != nil {
		// BUG: User forgot to rollback!
		return errors.New("stock update failed")
	}

	_, err = tx.Exec("INSERT INTO orders (user, item) VALUES (?, ?)", userID, itemID)
	if err != nil {
		// BUG: User forgot to rollback!
		return errors.New("order creation failed")
	}

	// BUG: User forgot to commit!
	// TODO: Commit the transaction and return its error.
	return nil
}
