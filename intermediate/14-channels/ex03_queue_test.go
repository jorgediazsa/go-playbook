package channels

import "testing"

func TestLogQueue(t *testing.T) {
	q := NewQueue(2)

	// Since capacity is 2, first two should succeed instantly.
	err1 := q.Enqueue("msg1")
	err2 := q.Enqueue("msg2")

	if err1 != nil || err2 != nil {
		t.Fatalf("Failed to enqueue logs into queue: %v, %v", err1, err2)
	}

	// 3rd should fail instantly! It must not block!
	// If the user didn't implement the select{ default: } block,
	// this next line will deadlock the test permanently (until test timeout).
	err3 := q.Enqueue("msg3")

	if err3 != ErrQueueFull {
		t.Fatalf("Expected ErrQueueFull, got %v", err3)
	}

	val1 := q.Dequeue()
	if val1 != "msg1" {
		t.Fatalf("Expected msg1, got %s", val1)
	}
}
