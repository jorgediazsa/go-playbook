package httpjson

import (
	"encoding/json"
	"errors"
)

// Context: Dynamic JSON Parsing
// You are building an Event Router for an analytics system. It receives
// generic "Events" where the `payload` changes depending on the `event_type`.
//
// Why this matters: If you unmarshal `{"event_type": "click", "payload": {"x": 5}}`
// into `GenericEvent`, the standard `json.Unmarshal` doesn't know what Go struct
// to map `"payload"` into, so it turns it into a `map[string]interface{}`.
// This destroys type-safety!
//
// Requirements:
// 1. Change the `Payload` field inside `GenericEvent` to a `json.RawMessage`.
//    This tells the decoder "Leave this as raw bytes for now."
// 2. Refactor `ProcessEvent`. First, unmarshal the bytes into `GenericEvent`.
// 3. Check `event.EventType`.
//    - If it's "click", unmarshal `event.Payload` into a `ClickPayload` struct, and return the `X` coordinate.
//    - If it's "scroll", unmarshal `event.Payload` into a `ScrollPayload` struct, and return the `Depth`.
//    - Otherwise, return -1.

// TODO: Change Payload type to json.RawMessage
type GenericEvent struct {
	EventType string `json:"event_type"`
	Payload   any    `json:"payload"`
}

type ClickPayload struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ScrollPayload struct {
	Depth int `json:"depth"`
}

func ProcessEvent(data []byte) (int, error) {
	// BUG: We are ignoring the dynamic nature of the payload.
	// TODO: Unmarshal into GenericEvent. Then switch on EventType.
	// Then Unmarshal the RawMessage into the specific payload struct!

	var event GenericEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return 0, err
	}

	// Because Payload was `any`, if you tried to access event.Payload.X it wouldn't compile!
	return -1, errors.New("unimplemented dynamic parsing")
}
