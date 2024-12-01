package utils

import (
	"testing"
)

func TestRandString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"small length", 8},
		{"medium length", 16},
		{"large length", 32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test multiple times to ensure randomness
			seen := make(map[string]bool)
			for i := 0; i < 100; i++ {
				got, err := RandString(tt.length)
				if err != nil {
					t.Errorf("RandString(%d) error = %v", tt.length, err)
					return
				}

				// Check length
				if len(got) != tt.length*2 { // *2 because hex encoding doubles length
					t.Errorf("RandString(%d) length = %d, want %d", tt.length, len(got), tt.length*2)
				}

				// Check for duplicates
				if seen[got] {
					t.Errorf("RandString(%d) generated duplicate value: %s", tt.length, got)
				}
				seen[got] = true
			}
		})
	}
}

func TestRandInt64(t *testing.T) {
	// Test multiple times to ensure randomness
	seen := make(map[int64]bool)
	for i := 0; i < 1000; i++ {
		got, err := RandInt64()
		if err != nil {
			t.Errorf("RandInt64() error = %v", err)
			return
		}

		// Check for unlikely duplicates
		if seen[got] {
			t.Errorf("RandInt64() generated duplicate value: %d", got)
		}
		seen[got] = true
	}
}
