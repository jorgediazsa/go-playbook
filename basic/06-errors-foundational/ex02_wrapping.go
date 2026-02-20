package errorsfunc

import (
	"errors"
	"fmt"
)

// Context: Wrapping context around errors
// You are writing an orchestration layer that fetches user data from a DB
// and then publishes a message to a queue.
//
// Why this matters: When a sub-component fails (e.g., the DB driver returns "connection lost"),
// returning that raw error to the HTTP layer gives no context to the caller.
// "connection lost" ... while doing what?
// However, if you create a brand new error (`fmt.Errorf("db fetch failed")`), you destroy
// the original error, so the caller can't do programatic checks against `ErrConnectionLost`.
//
// Requirements:
// 1. In `Orchestrate`, if `fetchUser` fails, wrap the error with the message "failed to fetch user: %w".
// 2. If `publishEvent` fails, wrap the error with the message "failed to publish event: %w".
// 3. Use the `%w` verb in `fmt.Errorf` to ensure `errors.Is` will still work.

var ErrConnectionLost = errors.New("connection lost")
var ErrQueueFull = errors.New("queue full")

func fetchUser(id string) error {
	return ErrConnectionLost // simulating a failure
}

func publishEvent(data string) error {
	return ErrQueueFull // simulating a failure
}

func Orchestrate(userID string) error {
	err := fetchUser(userID)
	if err != nil {
		// BUG: Returning raw error deletes context.
		// TODO: Wrap it!
		return err
	}

	err = publishEvent("user_created")
	if err != nil {
		// BUG: Creating a new error destroys the original type (`%v` instead of `%w`).
		// TODO: Wrap it correctly using `%w`.
		return fmt.Errorf("failed to publish event: %v", err)
	}

	return nil
}
