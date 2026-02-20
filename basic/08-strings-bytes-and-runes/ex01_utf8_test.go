package stringsbytes

import "testing"

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxChars int
		expected string
	}{
		{"ASCII Short", "Hello", 10, "Hello"},
		{"ASCII Long", "Hello World", 5, "Hello..."},
		{"Multibyte Exact", "こんにちは", 5, "こんにちは"},
		{"Multibyte Truncate", "こんにちは世界", 5, "こんにちは..."},
		{"Emojis Truncate", "🚀🌍🚢🚗", 2, "🚀🌍..."},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Truncate(tc.input, tc.maxChars)
			if got != tc.expected {
				t.Fatalf("Truncate(%q, %d) = %q, want %q", tc.input, tc.maxChars, got, tc.expected)
			}
		})
	}
}
