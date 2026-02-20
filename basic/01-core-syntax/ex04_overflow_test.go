package core

import (
	"math"
	"testing"
)

func TestSafeConvert(t *testing.T) {
	tests := []struct {
		name      string
		input     int64
		want      uint32
		wantError error
	}{
		{"Zero", 0, 0, nil},
		{"Small positive", 42, 42, nil},
		{"Max uint32 exact", math.MaxUint32, math.MaxUint32, nil},
		{"Max uint32 + 1 overflows", math.MaxUint32 + 1, 0, ErrOverflow},
		{"Huge positive overflows", math.MaxInt64, 0, ErrOverflow},
		{"Negative underflows", -1, 0, ErrOverflow},
		{"Huge negative underflows", math.MinInt64, 0, ErrOverflow},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SafeConvertInt64ToUint32(tc.input)
			if tc.wantError != nil {
				if err != tc.wantError {
					t.Fatalf("Expected error %v, got %v", tc.wantError, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Expected no error, got %v", err)
				}
				if got != tc.want {
					t.Fatalf("Expected %v, got %v", tc.want, got)
				}
			}
		})
	}
}
