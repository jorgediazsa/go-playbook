package errorsadv

// Context: Aggregating Errors (`errors.Join`)
// You are validating a massive JSON payload with 3 independent fields.
// If all 3 fields are invalid, you want to return ALL 3 errors to the client
// at once, rather than dropping the request on the very first failure, forcing
// them to fix it, submit again, and hit the second failure.
//
// Why this matters: Historically, developers used third-party packages like
// `hashicorp/go-multierror`. In Go 1.20, `errors.Join(err1, err2...)` was added
// directly to the standard library.
// The magic of `Join` is that the resulting aggregate error perfectly satisfies
// `errors.Is` for EVERY error inside the join array!
//
// Requirements:
// 1. Refactor `ValidatePayload` to validate all three fields using `errors.Join`.
// 2. Return a single joined error. If no fields are broken, return `nil`.

import "errors"

var ErrMissingEmail = errors.New("email is required")
var ErrEmailInvalid = errors.New("email format invalid")
var ErrPasswordShort = errors.New("password too short")

func ValidatePayload(email string, password string) error {
	// BUG: We exit on the first failure! The client won't know the password
	// is also short if they forgot the email!
	// TODO: Collect the errors and return them combined with `errors.Join`.

	if email == "" {
		return ErrMissingEmail
	}

	if password == "" || len(password) < 8 {
		return ErrPasswordShort
	}

	return nil
}
