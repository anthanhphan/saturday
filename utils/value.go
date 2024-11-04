package utils

import "github.com/anthanhphan/saturday/validate"

// DefaultIfEmpty determines whether to return the original value or a fallback.
// It uses validate.IsZero to check if the input value is considered "zero".
//
// Parameters:
//   - value: The value to check
//   - fallback: The value to return if empty
//
// Returns:
//   - T: The original value if not empty, fallback otherwise
//
// Examples:
//
//	DefaultIfEmpty[string]("", "fallback")      // returns "fallback"
//	DefaultIfEmpty[int](0, 42)                  // returns 42
//	DefaultIfEmpty[string]("hello", "fallback") // returns "hello"
func DefaultIfEmpty[T any](value any, fallback T) T {
	if validate.IsZero(value) {
		return fallback
	}

	return value.(T)
}

// GetFromInterface safely retrieves a typed value from a map with a default fallback.
// Supports optional zero-value checking for determining when to use the default.
//
// Parameters:
//   - src: Source map to retrieve value from
//   - key: Key to look up in the map
//   - defaultValue: Value to return if key missing or value is zero
//   - checkZeroValue: Optional flag to enable zero-value checking
//
// Returns:
//   - T: The found value (type-asserted to T) or defaultValue
//
// Examples:
//
//	m := map[string]any{"count": 5, "name": ""}
//	GetFromInterface[int](m, "count", 0)                // returns 5
//	GetFromInterface[string](m, "missing", "default")   // returns "default"
//	GetFromInterface[string](m, "name", "default", true) // returns "default"
func GetFromInterface[T any](src map[string]any, key string, defaultValue T, checkZeroValue ...bool) T {
	value, exists := src[key]
	if !exists || (validate.IsZero(value) && len(checkZeroValue) > 0 && checkZeroValue[0]) {
		return defaultValue
	}
	return value.(T)
}
