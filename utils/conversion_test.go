package utils_test

import (
	"reflect"
	"testing"

	"github.com/anthanhphan/saturday/utils"
)

func TestArrayInt64ToString(t *testing.T) {
	tests := []struct {
		name     string
		input    []int64
		delim    string
		expected string
	}{
		{
			name:     "normal sequence",
			input:    []int64{1, 2, 3},
			delim:    ",",
			expected: "1,2,3",
		},
		{
			name:     "single number",
			input:    []int64{42},
			delim:    ",",
			expected: "42",
		},
		{
			name:     "empty array",
			input:    []int64{},
			delim:    ",",
			expected: "",
		},
		{
			name:     "different delimiter",
			input:    []int64{1, 2, 3},
			delim:    "-",
			expected: "1-2-3",
		},
		{
			name:     "large numbers",
			input:    []int64{9223372036854775807, -9223372036854775808},
			delim:    ",",
			expected: "9223372036854775807,-9223372036854775808",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ArrayInt64ToString(tt.input, tt.delim)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestArrayStringToString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		delim    string
		expected string
	}{
		{
			name:     "normal sequence",
			input:    []string{"a", "b", "c"},
			delim:    ",",
			expected: "a,b,c",
		},
		{
			name:     "single string",
			input:    []string{"hello"},
			delim:    ",",
			expected: "hello",
		},
		{
			name:     "empty array",
			input:    []string{},
			delim:    ",",
			expected: "",
		},
		{
			name:     "empty strings",
			input:    []string{"", "", ""},
			delim:    ",",
			expected: ",,",
		},
		{
			name:     "different delimiter",
			input:    []string{"a", "b", "c"},
			delim:    " | ",
			expected: "a | b | c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ArrayStringToString(tt.input, tt.delim)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStringToArrayInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		delim    string
		expected []int64
	}{
		{
			name:     "normal sequence",
			input:    "1,2,3",
			delim:    ",",
			expected: []int64{1, 2, 3},
		},
		{
			name:     "with spaces",
			input:    "1, 2, 3",
			delim:    ",",
			expected: []int64{1, 2, 3},
		},
		{
			name:     "invalid numbers",
			input:    "1,invalid,3",
			delim:    ",",
			expected: []int64{1, 3},
		},
		{
			name:     "empty string",
			input:    "",
			delim:    ",",
			expected: []int64{},
		},
		{
			name:     "only spaces",
			input:    "  ,  ,  ",
			delim:    ",",
			expected: []int64{},
		},
		{
			name:     "large numbers",
			input:    "9223372036854775807,-9223372036854775808",
			delim:    ",",
			expected: []int64{9223372036854775807, -9223372036854775808},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.StringToArrayInt64(tt.input, tt.delim)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStringToArrayString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		delim    string
		trim     bool
		expected []string
	}{
		{
			name:     "normal sequence no trim",
			input:    "a,b,c",
			delim:    ",",
			trim:     false,
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "normal sequence with trim",
			input:    "a, b, c",
			delim:    ",",
			trim:     true,
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "empty entries no trim",
			input:    "a,,c",
			delim:    ",",
			trim:     false,
			expected: []string{"a", "", "c"},
		},
		{
			name:     "empty entries with trim",
			input:    "a,,c",
			delim:    ",",
			trim:     true,
			expected: []string{"a", "c"},
		},
		{
			name:     "empty string",
			input:    "",
			delim:    ",",
			trim:     true,
			expected: []string{},
		},
		{
			name:     "only spaces with trim",
			input:    "  ,  ,  ",
			delim:    ",",
			trim:     true,
			expected: []string{},
		},
		{
			name:     "only spaces no trim",
			input:    "  ,  ,  ",
			delim:    ",",
			trim:     false,
			expected: []string{"  ", "  ", "  "},
		},
		{
			name:     "different delimiter",
			input:    "a|b|c",
			delim:    "|",
			trim:     false,
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.StringToArrayString(tt.input, tt.delim, tt.trim)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
