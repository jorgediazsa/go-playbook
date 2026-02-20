package reflection

import (
	"testing"
)

func TestRefactorToInterface(t *testing.T) {
	/* TODO: Uncomment this test after implementing the Identifiable interface and LogID function!

	o := Order{ID: "ORD-123", Amount: 50}
	d := Device{ID: "DEV-ABC", Mac: "00:00"}

	// LogID must accept both via the Identifiable interface
	id1 := LogID(o)
	id2 := LogID(d)

	if id1 != "ORD-123" {
		t.Fatalf("Failed to log order ID")
	}

	if id2 != "DEV-ABC" {
		t.Fatalf("Failed to log device ID")
	}

	// The beauty of this: if someone passes a plain string `LogID("hello")`,
	// it fails at COMPILE TIME! No reflection, no runtime panics, 100x faster.
	*/
}
