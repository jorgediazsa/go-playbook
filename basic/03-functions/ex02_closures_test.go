package functions

import "testing"

func TestCreateHandlers(t *testing.T) {
	handlers := CreateHandlers()

	if len(handlers) != 3 {
		t.Fatalf("Expected 3 handlers, got %d", len(handlers))
	}

	// Wait, we call them now. If the bug exists, they all return 3.
	// If fixed, they return 0, 1, 2.

	for i, h := range handlers {
		val := h()
		if val != i {
			t.Errorf("Handler %d returned %d, likely due to closure capture bug.", i, val)
		}
	}
}
