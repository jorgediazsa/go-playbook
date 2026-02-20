package production

// Context: API Backward Compatibility
// You are upgrading your User API. The mobile app team built the app expecting
// `User.Name` to be a single string.
// You just refactored the database to split names into `FirstName` and `LastName`.
//
// Why this matters: If you deploy v2 of the API and return `{"first_name": "John", "last_name": "Doe"}`,
// the mobile app will crash because it cannot find the `"name"` field.
// You MUST maintain backward compatibility for older clients.
//
// Requirements:
// 1. Refactor the `UserResponseV2` struct so that it fulfills BOTH old and new
//    client requirements simultaneously.
// 2. Or, better yet, explicitly define `UserResponseV1` and `UserResponseV2`
//    and route them based on the request (simulated below).
// 3. For this exercise, fix `FormatUser` so it returns a struct containing
//    `"name"` (the concatenated full name) AND the new `"first_name"`, `"last_name"`.
//    This allows old clients to read `"name"` safely, while new clients migrate to the new fields.

import "encoding/json"

type DBUser struct {
	FirstName string
	LastName  string
}

// BUG: We deleted the `Name` field! 100,000 mobile clients will instantly break
// when we deploy this JSON shape!
// TODO: Restore the `"name"` JSON field while keeping the new fields, avoiding a breaking change.
type UserResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func FormatUser(u DBUser) ([]byte, error) {
	resp := UserResponse{
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}

	return json.Marshal(resp)
}
