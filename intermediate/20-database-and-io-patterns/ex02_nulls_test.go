package dbio

import (
	"context"
	"database/sql"
	"testing"
)

// MockRow simulates a DB row returning a NULL for the 3rd column (Bio).
type MockNullRow struct{}

func (m MockNullRow) Scan(dest ...any) error {
	// The first two columns are fine
	*dest[0].(*int) = 1
	*dest[1].(*string) = "Alice"

	// The third column is NULL.
	// If the user passed a `*string`, it panics/errors without sql.Scanner logic,
	// EXCEPT that the standard library's DB driver handles `**string` explicitly!
	// Here we simulate the driver's strictness.

	// If the dest[2] is a `*sql.NullString`:
	if ns, ok := dest[2].(*sql.NullString); ok {
		ns.Valid = false
		return nil
	}

	// If the dest[2] is a `**string` (pointer to a pointer approach):
	if ps, ok := dest[2].(**string); ok {
		*ps = nil
		return nil
	}

	return sql.ErrNoRows // A generic error replacing the actual panic/error for this test
}

func TestFetchUserNullBio(t *testing.T) {
	row := MockNullRow{}

	// The user's code must compile and run against this mock.
	user, err := FetchUser(context.Background(), row)

	if err != nil {
		t.Fatalf("FAILED: Scan returned an error! You must change `Bio` to `sql.NullString` or `*string` to handle SQL NULLs. Error: %v", err)
	}

	// We type-check the struct lightly depending on what they chose.
	// Since we can't strictly compile against two different possible correct answers
	// perfectly in a static test, the fact that `FetchUser` returned no error
	// against `MockNullRow` proves they fixed the struct and `Scan()` signature!

	_ = user
}
