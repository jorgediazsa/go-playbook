package dbio

import (
	"context"
)

// Context: Handling SQL Nulls
// You are building an API that fetches a user profile. In the database,
// the `bio` column is nullable (it can literally be NULL).
//
// Why this matters: If your Go struct uses `Bio string`, and the database
// tries to scan a NULL into it, `sql.Row.Scan()` will return an error
// because Go strings cannot be nil.
//
// Requirements:
// 1. Refactor the `User` struct so that `Bio` can safely accept a SQL NULL.
//    (Hint: You can use `sql.NullString`, OR idiomatically in modern Go,
//    a pointer `*string`).
// 2. Refactor `FetchUser` to return the new struct format safely.

// We mock a row scanner
type RowScanner interface {
	Scan(dest ...any) error
}

// BUG: Go strings cannot be NULL. This will panic or error on Scan!
// TODO: Change Bio so it can safely represent both a string and a SQL NULL.
type User struct {
	ID   int
	Name string
	Bio  string
}

func FetchUser(ctx context.Context, scanner RowScanner) (User, error) {
	var u User

	// If Bio in the DB is NULL, this Scan() will throw an error because `&u.Bio` is a `*string`.
	err := scanner.Scan(&u.ID, &u.Name, &u.Bio)

	if err != nil {
		return User{}, err
	}

	return u, nil
}
