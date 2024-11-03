package validate

import (
	"testing"
)

func TestIsZero(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"empty string", "", true},
		{"non-empty string", "hello", false},
		{"zero int", 0, true},
		{"non-zero int", 42, false},
		{"empty slice", []int{}, true},
		{"non-empty slice", []int{1, 2, 3}, false},
		{"nil slice", []int(nil), true},
		{"empty array", [0]int{}, true},
		{"non-empty array", [3]int{1, 2, 3}, false},
		{"zero struct", struct{}{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZero(tt.input); got != tt.expected {
				t.Errorf("IsZero() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsPrimaryKeyNonZero(t *testing.T) {
	type TestStruct struct {
		ID   int `gorm:"primaryKey"`
		Name string
	}

	type NoPKStruct struct {
		Name string
	}

	tests := []struct {
		name        string
		input       any
		wantErr     bool
		errContains string
	}{
		{
			name:        "zero primary key",
			input:       TestStruct{ID: 0, Name: "test"},
			wantErr:     true,
			errContains: "primary key cannot be zero",
		},
		{
			name:        "non-zero primary key",
			input:       TestStruct{ID: 1, Name: "test"},
			wantErr:     false,
			errContains: "",
		},
		{
			name:        "struct without primary key",
			input:       NoPKStruct{Name: "test"},
			wantErr:     true,
			errContains: "no primary key found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsPrimaryKeyNonZero(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrimaryKeyNonZero() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.errContains != "" && err.Error() != tt.errContains {
				t.Errorf("IsPrimaryKeyNonZero() error = %v, want error containing %v", err, tt.errContains)
			}
		})
	}
}

func TestHasNonZeroExcludingKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]any
		skipKeys map[string]any
		expected bool
	}{
		{
			name: "has non-zero values",
			input: map[string]any{
				"id":    1,
				"name":  "",
				"email": "test@example.com",
			},
			skipKeys: map[string]any{
				"name": nil,
			},
			expected: true,
		},
		{
			name: "all zero or skipped",
			input: map[string]any{
				"id":   0,
				"name": "",
			},
			skipKeys: map[string]any{
				"id": nil,
			},
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[string]any{},
			skipKeys: map[string]any{},
			expected: false,
		},
		{
			name: "all values skipped",
			input: map[string]any{
				"id":   1,
				"name": "test",
			},
			skipKeys: map[string]any{
				"id":   nil,
				"name": nil,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasNonZeroExcludingKeys(tt.input, tt.skipKeys); got != tt.expected {
				t.Errorf("HasNonZeroExcludingKeys() = %v, want %v", got, tt.expected)
			}
		})
	}
}
