package dbio

import (
	"encoding/json"
	"io"
)

// Context: Streaming JSON (OOM Prevention)
// You are downloading a 5GB JSON file containing millions of `{"id": int, "email": string}`
// objects.
//
// Why this matters: `json.Unmarshal(data, &slice)` reads the ENTIRE 5GB string into Ram,
// then builds a 5GB Array of structs in Ram. This instantly crashes (OOMs) most
// smaller pods or containers.
//
// Requirements:
// 1. Refactor `CountFastmailUsers` to NOT use `json.Unmarshal` or `io.ReadAll`.
// 2. Use `json.NewDecoder(reader)`.
// 3. Keep memory footprint near zero by decoding records one by one in a loop,
//    counting them, and letting the garbage collector sweep them instantly.
//    (Hint: Use `decoder.Decode(&u)` inside a loop until `err == io.EOF`).

type JSONUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func CountFastmailUsers(reader io.Reader) (int, error) {
	// BUG: This approach buffers the entire payload into memory twice!
	// TODO: Replace this entirely with `json.NewDecoder`.

	data, err := io.ReadAll(reader)
	if err != nil {
		return 0, err
	}

	var users []JSONUser
	err = json.Unmarshal(data, &users)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, u := range users {
		if u.Email == "fastmail.com" { // Simplified domain check for exercise
			count++
		}
	}

	return count, nil
}
