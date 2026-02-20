package httpjson

import "testing"

func TestProcessEventDynamic(t *testing.T) {
	clickJSON := []byte(`{"event_type": "click", "payload": {"x": 105, "y": 200}}`)
	scrollJSON := []byte(`{"event_type": "scroll", "payload": {"depth": 900}}`)
	unknownJSON := []byte(`{"event_type": "hover", "payload": {}}`)

	// Test Click Extract
	val, err := ProcessEvent(clickJSON)
	if err != nil {
		t.Fatalf("FAILED: ProcessEvent returned error on click: %v", err)
	}
	if val != 105 {
		t.Fatalf("Expected click X to be 105, got %d. Did you successfully unmarshal the RawMessage?", val)
	}

	// Test Scroll Extract
	val, err = ProcessEvent(scrollJSON)
	if err != nil {
		t.Fatalf("FAILED: ProcessEvent returned error on scroll: %v", err)
	}
	if val != 900 {
		t.Fatalf("Expected scroll Depth to be 900, got %d", val)
	}

	// Test Unknown
	val, _ = ProcessEvent(unknownJSON)
	if val != -1 {
		t.Fatalf("Expected -1 for unknown event types, got %d", val)
	}
}
