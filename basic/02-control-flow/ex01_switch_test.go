package control

import (
	"testing"
)

func TestRoutePriority(t *testing.T) {
	tests := []struct {
		name     string
		req      Request
		expected string
	}{
		{
			name:     "VIP small payload",
			req:      Request{Path: "/api/users", Method: "GET", IsVIP: true, PayloadSize: 50},
			expected: "Priority-High",
		},
		{
			name:     "VIP large payload - falls through to normal rules",
			req:      Request{Path: "/admin/settings", Method: "POST", IsVIP: true, PayloadSize: 500},
			expected: "Priority-Critical",
		},
		{
			name:     "Admin pathway is critical",
			req:      Request{Path: "/admin/users", Method: "GET", IsVIP: false, PayloadSize: 10},
			expected: "Priority-Critical",
		},
		{
			name:     "API POST is medium",
			req:      Request{Path: "/api/users", Method: "POST", IsVIP: false, PayloadSize: 200},
			expected: "Priority-Medium",
		},
		{
			name:     "API GET is low",
			req:      Request{Path: "/api/users", Method: "GET", IsVIP: false, PayloadSize: 200},
			expected: "Priority-Low",
		},
		{
			name:     "Random endpoint is low",
			req:      Request{Path: "/health", Method: "GET", IsVIP: false, PayloadSize: 10},
			expected: "Priority-Low",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RoutePriority(tc.req)
			if got != tc.expected {
				t.Fatalf("RoutePriority() = %v, want %v", got, tc.expected)
			}
		})
	}
}
