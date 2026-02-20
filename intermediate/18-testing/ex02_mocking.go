package testingadv

import "fmt"

// Context: Mocking via Interfaces
// You are building an Alerting service that queries a legacy User Database
// to find a user's phone number, then sends an SMS.
//
// Why this matters: The original author hardcoded the concrete `*LegacyUserDB`
// struct into the function. Now, it is IMPOSSIBLE to write a fast, deterministic
// unit test without spinning up a real Postgres database in Docker.
//
// Requirements:
// 1. Define a small, local interface (e.g., `UserFetcher`) that contains ONLY
//    the method(s) that `SendUrgentAlert` actually uses.
// 2. Refactor `SendUrgentAlert` to accept your interface instead of the concrete DB.
// 3. Open `ex02_mocking_test.go` and implement a mock struct.

// --- EXTERNAL PACKAGE ---
// Imagine this is in a completely different repository/package.
type LegacyUserDB struct{}

func (db *LegacyUserDB) GetPhone(userID string) (string, error) {
	// ... connects to Postgres ...
	return "+15550000", nil
}
func (db *LegacyUserDB) GetAddress(userID string) string { return "123 Main St" }

// ------------------------

// TODO: Define your minimal interface here.

// BUG: This function takes a concrete DB connection, making testing a nightmare.
// TODO: Change the parameter `db` to use your new interface.
func SendUrgentAlert(db *LegacyUserDB, userID string, msg string) error {

	phone, err := db.GetPhone(userID)
	if err != nil {
		return err
	}

	// Complex SMS logic here...
	fmt.Printf("Sending SMS to %s: %s\n", phone, msg)
	return nil
}
