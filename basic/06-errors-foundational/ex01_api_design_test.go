package errorsfunc

import "testing"

func TestLoadConfig(t *testing.T) {
	// 1. Test Empty
	_, err := LoadConfig("")
	if err != ErrEmptyPayload {
		t.Errorf("Expected ErrEmptyPayload, got %v", err)
	}

	// 2. Test Invalid
	_, err = LoadConfig("invalid_string")
	if err != ErrInvalidFormat {
		t.Errorf("Expected ErrInvalidFormat, got %v", err)
	}

	// 3. Test Valid
	cfg, err := LoadConfig("valid")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cfg.Timeout != 30 || cfg.Retries != 3 {
		t.Errorf("Config was not populated correctly: %+v", cfg)
	}
}
