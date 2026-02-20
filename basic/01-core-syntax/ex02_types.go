package core

// Title: Alias vs Defined Type Behavior
// Context: You're working on a core domain model for an e-commerce platform.
// In the legacy code, all identifiers (UserID, ProductID) were created as type aliases (= string).
//
// Why this matters: Type aliases (`type UserID = string`) do not create a new type.
// They are completely interchangeable with strings and with each other. This causes
// massive bugs where developers accidentally pass a ProductID into a function taking a UserID.
//
// Requirements:
// 1. Refactor `UserID` and `ProductID` to be strict defined types (not aliases).
// 2. Refactor `ProcessOrder` to securely take a UserID and ProductID.
// 3. Complete `FetchUser` which mimics fetching data from a legacy database driver
//    that only understands raw strings. You must safely cast to/from the defined types.
//
// Note: When you fix requirement 1, your code might initially fail to compile due to strict typing.
// Fix the compilation errors by applying correct type conversions.

// TODO: Change these from type aliases to defined types.
type UserID = string
type ProductID = string

// ProcessOrder matches a user and product.
func ProcessOrder(u UserID, p ProductID) string {
	// TODO: implement logic that returns formatted string: "Order: <user> bought <product>"
	// Ensure you are safely coercing them back to strings if needed for formatting.
	return "Order: " + string(u) + " bought " + string(p) // string cast works for both aliases and defined
}

// LegacyFetch fetches from DB using string ID. Do not change this function's signature.
func LegacyFetch(id string) string {
	return "Data for " + id
}

// FetchUser fetches user data. It must use LegacyFetch internally, doing correct conversion.
func FetchUser(u UserID) string {
	// TODO: If you changed UserID to a defined type, passing `u` directly to LegacyFetch
	// will fail to compile. Fix the conversion!
	return LegacyFetch(u)
}
