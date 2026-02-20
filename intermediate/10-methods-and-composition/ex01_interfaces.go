package composition

// Title: Interface-Driven Design & Minimal Contracts
//
// Context: You are reviewing a PR for a new User Onboarding service.
// The author tightly coupled the logic to a massive `ThirdPartyEmailClientImpl` struct.
// Now, the team wants to write unit tests, but they can't without actually sending real emails
// to the external API, because the concrete struct is hardcoded into the business logic.
//
// Why this matters: In Go, producers (the email client) don't declare interfaces.
// Consumers (the onboarding service) declare small, implicit interfaces outlining *exactly*
// what they need. This makes dependency injection and mocking trivial without huge mocking frameworks.
//
// Requirements:
// 1. Refactor `OnboardUser` to accept an interface, not the concrete `*ThirdPartyEmailClientImpl`.
// 2. Define a minimal interface `EmailSender` right here in this package that
//    `ThirdPartyEmailClientImpl` implicitly satisfies.
// 3. Fix the `MockClient` in the `_test.go` file to satisfy your new interface so tests compile.

// --- EXTERNAL PACKAGE MOCK ---
// Imagine this comes from `github.com/enterprise/email`
// You cannot modify this struct!
type ThirdPartyEmailClientImpl struct {
	APIKey string
}

func (c *ThirdPartyEmailClientImpl) SendEmail(to, subject, body string) error {
	// ... Does complex HTTP calls ...
	return nil
}

func (c *ThirdPartyEmailClientImpl) PingServer() error {
	return nil
}

// -----------------------------

// TODO: Define the Consumer Interface `EmailSender` here.
// It should only include the methods `OnboardUser` actually uses!

// BUG: Tightly coupled to a concrete implementation.
// TODO: Refactor `client` to use your new interface.
func OnboardUser(client *ThirdPartyEmailClientImpl, email string) error {

	err := client.SendEmail(email, "Welcome!", "Thanks for signing up.")
	if err != nil {
		return err
	}

	return nil
}
