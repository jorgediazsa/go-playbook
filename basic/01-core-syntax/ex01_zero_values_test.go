package core

import (
	"testing"
)

func TestApplyPatch(t *testing.T) {
	base := UserSettings{
		ID:            "user-123",
		Notifications: true,
		Retries:       5,
	}

	tests := []struct {
		name     string
		patch    UserSettingsPatch
		expected UserSettings
	}{
		{
			name:     "Empty patch should not change anything",
			patch:    UserSettingsPatch{},
			expected: base,
		},
		{
			name: "Update Notifications to false (zero value)",
			patch: UserSettingsPatch{
				Notifications: OptionalBool{Value: false, Valid: true},
			},
			expected: UserSettings{ID: "user-123", Notifications: false, Retries: 5},
		},
		{
			name: "Update Retries to 0 (zero value) but leave notifications alone",
			patch: UserSettingsPatch{
				Retries: OptionalInt{Value: 0, Valid: true},
			},
			expected: UserSettings{ID: "user-123", Notifications: true, Retries: 0},
		},
		{
			name: "Update both to new values",
			patch: UserSettingsPatch{
				Notifications: OptionalBool{Value: false, Valid: true},
				Retries:       OptionalInt{Value: 10, Valid: true},
			},
			expected: UserSettings{ID: "user-123", Notifications: false, Retries: 10},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ApplyPatch(base, tc.patch)
			if result.ID != tc.expected.ID {
				t.Errorf("ID: expected %v, got %v", tc.expected.ID, result.ID)
			}
			if result.Notifications != tc.expected.Notifications {
				t.Errorf("Notifications: expected %v, got %v", tc.expected.Notifications, result.Notifications)
			}
			if result.Retries != tc.expected.Retries {
				t.Errorf("Retries: expected %v, got %v", tc.expected.Retries, result.Retries)
			}
		})
	}
}
