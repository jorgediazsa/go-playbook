package control

import (
	"reflect"
	"testing"
)

func TestGetSortedVIPs(t *testing.T) {
	vips := map[string]bool{
		"Zack": true,
		"Abby": true,
		"Mick": true,
		"Carl": true,
	}

	expected := []string{"Abby", "Carl", "Mick", "Zack"}

	// Run multiple times to verify map randomization fix
	for i := 0; i < 20; i++ {
		result := GetSortedVIPs(vips)
		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("Iteration %d: Expected %v, got %v. Ensure output is sorted.", i, expected, result)
		}
	}
}

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"ASCII", "Hello", 5},
		{"Emojis", "👋🌍", 2},
		{"Japanese", "こんにちは", 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountCharacters(tc.input)
			if got != tc.expected {
				t.Errorf("CountCharacters(%q) = %d, want %d", tc.input, got, tc.expected)
			}
		})
	}
}
