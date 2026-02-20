package testingadv

import "testing"

func TestParseMetric(t *testing.T) {
	// BUG: The table is empty!
	// TODO: Fill the table with the following scenarios:
	// 1. name: "Empty String", input: "",      wantName: "",     wantVal: 0, wantErr: true
	// 2. name: "Valid Simple", input: "cpu:5",  wantName: "cpu",  wantVal: 1, wantErr: false
	// 3. name: "Empty Value",  input: "mem:",   wantName: "mem",  wantVal: 0, wantErr: false
	// 4. name: "No Colon",     input: "disk",   wantName: "",     wantVal: 0, wantErr: true

	tests := []struct {
		name     string
		input    string
		wantName string
		wantVal  int
		wantErr  bool
	}{
		// TODO: Add test cases here
	}

	// This is defensive to ensure the exercise compiles and fails loudly if not implemented.
	if len(tests) == 0 {
		t.Fatal("FAILED: Test table is empty. Add the 4 test cases!")
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotName, gotVal, err := ParseMetric(tc.input)

			if (err != nil) != tc.wantErr {
				t.Fatalf("ParseMetric(%q) error = %v, wantErr %v", tc.input, err, tc.wantErr)
			}
			if gotName != tc.wantName {
				t.Errorf("ParseMetric(%q) gotName = %v, want %v", tc.input, gotName, tc.wantName)
			}
			if gotVal != tc.wantVal {
				t.Errorf("ParseMetric(%q) gotVal = %v, want %v", tc.input, gotVal, tc.wantVal)
			}
		})
	}
}
