package utils

import (
	"strconv"
	"strings"
)

// ArrayInt64ToString converts a slice of int64 to a delimited string.
// It efficiently joins all numbers with the specified delimiter using a string builder.
//
// Parameters:
//   - arr: The slice of int64 numbers to convert
//   - delim: The delimiter to insert between numbers
//
// Returns:
//   - string: The resulting delimited string, or empty string if input is empty
//
// Examples:
//
//	ArrayInt64ToString([]int64{1, 2, 3}, ",")    // returns "1,2,3"
//	ArrayInt64ToString([]int64{}, ",")           // returns ""
//	ArrayInt64ToString([]int64{42}, "-")         // returns "42"
func ArrayInt64ToString(arr []int64, delim string) string {
	if len(arr) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(arr) * 8)

	builder.WriteString(strconv.FormatInt(arr[0], 10))
	for _, v := range arr[1:] {
		builder.WriteString(delim)
		builder.WriteString(strconv.FormatInt(v, 10))
	}
	return builder.String()
}

// ArrayStringToString converts a slice of strings to a delimited string.
// It efficiently joins all strings with the specified delimiter using a string builder.
//
// Parameters:
//   - arr: The slice of strings to convert
//   - delim: The delimiter to insert between strings
//
// Returns:
//   - string: The resulting delimited string, or empty string if input is empty
//
// Examples:
//
//	ArrayStringToString([]string{"a", "b", "c"}, ",")    // returns "a,b,c"
//	ArrayStringToString([]string{}, ",")                 // returns ""
//	ArrayStringToString([]string{"hello"}, " ")          // returns "hello"
func ArrayStringToString(arr []string, delim string) string {
	if len(arr) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(arr) * 8)

	builder.WriteString(arr[0])
	for _, v := range arr[1:] {
		builder.WriteString(delim)
		builder.WriteString(v)
	}
	return builder.String()
}

// StringToArrayInt64 converts a delimited string to a slice of int64 numbers.
// It splits the string by the delimiter, trims whitespace, and converts valid numbers.
// Invalid numbers and empty strings are skipped.
//
// Parameters:
//   - str: The input string to split and convert
//   - delim: The delimiter that separates the numbers
//
// Returns:
//   - []int64: A slice containing the successfully parsed numbers
//
// Examples:
//
//	StringToArrayInt64("1,2,3", ",")           // returns []int64{1, 2, 3}
//	StringToArrayInt64("1, 2, 3", ",")         // returns []int64{1, 2, 3}
//	StringToArrayInt64("1,invalid,3", ",")     // returns []int64{1, 3}
//	StringToArrayInt64("", ",")                // returns []int64{}
func StringToArrayInt64(str string, delim string) []int64 {
	parts := strings.Split(str, delim)
	result := make([]int64, 0, len(parts))
	for _, v := range parts {
		v = strings.TrimSpace(v)
		if v != "" {
			if i, err := strconv.ParseInt(v, 10, 64); err == nil {
				result = append(result, i)
			}
		}
	}
	return result
}

// StringToArrayString converts a delimited string to a slice of strings.
// It splits the string by the delimiter and optionally trims whitespace and empty entries.
//
// Parameters:
//   - str: The input string to split
//   - delim: The delimiter to split on
//   - trim: If true, trims whitespace from each part and removes empty strings
//
// Returns:
//   - []string: The resulting slice of strings
//
// Examples:
//
//	StringToArrayString("a,b,c", ",", false)      // returns []string{"a", "b", "c"}
//	StringToArrayString("a, b, c", ",", true)     // returns []string{"a", "b", "c"}
//	StringToArrayString("a,,c", ",", false)       // returns []string{"a", "", "c"}
//	StringToArrayString("a,,c", ",", true)        // returns []string{"a", "c"}
//	StringToArrayString("", ",", true)            // returns []string{}
func StringToArrayString(str, delim string, trim bool) []string {
	if str == "" {
		return []string{}
	}

	parts := strings.Split(str, delim)
	if !trim && delim != "" {
		return parts
	}

	result := make([]string, 0, len(parts))
	for _, v := range parts {
		if trim {
			v = strings.TrimSpace(v)
		}
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}
