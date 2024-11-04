package utils

import (
	"testing"
)

func TestDefaultIfEmpty(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		fallback any
		want     any
	}{
		{
			name:     "empty string returns fallback",
			value:    "",
			fallback: "default",
			want:     "default",
		},
		{
			name:     "non-empty string returns original",
			value:    "hello",
			fallback: "default",
			want:     "hello",
		},
		{
			name:     "zero int returns fallback",
			value:    0,
			fallback: 42,
			want:     42,
		},
		{
			name:     "non-zero int returns original",
			value:    10,
			fallback: 42,
			want:     10,
		},
		{
			name:     "nil slice returns fallback",
			value:    []string(nil),
			fallback: []string{"default"},
			want:     []string{"default"},
		},
		{
			name:     "empty slice returns fallback",
			value:    []string{},
			fallback: []string{"default"},
			want:     []string{"default"},
		},
		{
			name:     "non-empty slice returns original",
			value:    []string{"hello"},
			fallback: []string{"default"},
			want:     []string{"hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.value.(type) {
			case string:
				got := DefaultIfEmpty[string](v, tt.fallback.(string))
				if got != tt.want {
					t.Errorf("DefaultIfEmpty() = %v, want %v", got, tt.want)
				}
			case int:
				got := DefaultIfEmpty[int](v, tt.fallback.(int))
				if got != tt.want {
					t.Errorf("DefaultIfEmpty() = %v, want %v", got, tt.want)
				}
			case []string:
				got := DefaultIfEmpty[[]string](v, tt.fallback.([]string))
				if !sliceEqual(got, tt.want.([]string)) {
					t.Errorf("DefaultIfEmpty() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestGetFromInterface(t *testing.T) {
	testMap := map[string]any{
		"string":     "value",
		"empty":      "",
		"int":        42,
		"zero":       0,
		"stringList": []string{"a", "b"},
		"emptyList":  []string{},
	}

	tests := []struct {
		name           string
		src            map[string]any
		key            string
		defaultValue   any
		checkZeroValue bool
		want           any
	}{
		{
			name:         "returns existing string",
			src:          testMap,
			key:          "string",
			defaultValue: "default",
			want:         "value",
		},
		{
			name:           "returns default for empty string with zero check",
			src:            testMap,
			key:            "empty",
			defaultValue:   "default",
			checkZeroValue: true,
			want:           "default",
		},
		{
			name:         "returns empty string without zero check",
			src:          testMap,
			key:          "empty",
			defaultValue: "default",
			want:         "",
		},
		{
			name:         "returns existing int",
			src:          testMap,
			key:          "int",
			defaultValue: 0,
			want:         42,
		},
		{
			name:           "returns default for zero int with zero check",
			src:            testMap,
			key:            "zero",
			defaultValue:   99,
			checkZeroValue: true,
			want:           99,
		},
		{
			name:         "returns zero int without zero check",
			src:          testMap,
			key:          "zero",
			defaultValue: 99,
			want:         0,
		},
		{
			name:         "returns default for missing key",
			src:          testMap,
			key:          "missing",
			defaultValue: "default",
			want:         "default",
		},
		{
			name:         "returns existing slice",
			src:          testMap,
			key:          "stringList",
			defaultValue: []string{"default"},
			want:         []string{"a", "b"},
		},
		{
			name:           "returns default for empty slice with zero check",
			src:            testMap,
			key:            "emptyList",
			defaultValue:   []string{"default"},
			checkZeroValue: true,
			want:           []string{"default"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.defaultValue.(type) {
			case string:
				got := GetFromInterface[string](tt.src, tt.key, tt.defaultValue.(string), tt.checkZeroValue)
				if got != tt.want {
					t.Errorf("GetFromInterface() = %v, want %v", got, tt.want)
				}
			case int:
				got := GetFromInterface[int](tt.src, tt.key, tt.defaultValue.(int), tt.checkZeroValue)
				if got != tt.want {
					t.Errorf("GetFromInterface() = %v, want %v", got, tt.want)
				}
			case []string:
				got := GetFromInterface[[]string](tt.src, tt.key, tt.defaultValue.([]string), tt.checkZeroValue)
				if !sliceEqual(got, tt.want.([]string)) {
					t.Errorf("GetFromInterface() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

// Helper function for comparing slices
func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
