package tag

import (
	"strings"
)

// Parse converts a delimited string of key-value pairs into a map.
// Each pair can be separated by the specified separator, and keys are automatically
// converted to uppercase. Keys and values are separated by a colon (':').
//
// The function handles escaped separators using backslashes ('\'), allowing the separator
// character to be included within values.
//
// Parameters:
//   - input: The string containing the key-value pairs
//   - separator: The string used to separate different key-value pairs
//
// Returns:
//   - map[string]string: A map containing the parsed key-value pairs where:
//   - Keys are converted to uppercase
//   - If a key has no value (no colon), the key is used as the value
//   - Empty keys are ignored
//
// Format: "KEY1:value1{separator}KEY2:value2{separator}KEY3"
//
// Examples (using ';' as separator):
//
//	"FOO:bar;BAZ:qux"     -> {"FOO": "bar", "BAZ": "qux"}
//	"FOO:bar;BAZ"         -> {"FOO": "bar", "BAZ": "BAZ"}
//	"FOO:bar\\;baz;QUX"   -> {"FOO": "bar;baz", "QUX": "QUX"}
func Parse(input string, separator string) map[string]string {
	capacity := strings.Count(input, separator) + 1
	settings := make(map[string]string, capacity)
	parts := splitEscaped(input, separator)

	for _, part := range parts {
		key, value := extractKeyValue(part)
		if key != "" {
			settings[key] = value
		}
	}

	return settings
}

func splitEscaped(input string, separator string) []string {
	// Preallocate slice with estimated capacity
	parts := strings.Split(input, separator)
	result := make([]string, 0, len(parts))
	var builder strings.Builder

	for i := 0; i < len(parts); i++ {
		current := parts[i]

		// Handle escaped separators
		if len(current) > 0 && current[len(current)-1] == '\\' {
			builder.Reset()
			builder.WriteString(current[:len(current)-1])

			for i+1 < len(parts) && len(parts[i+1]) > 0 && parts[i+1][len(parts[i+1])-1] == '\\' {
				builder.WriteString(separator)
				builder.WriteString(parts[i+1][:len(parts[i+1])-1])
				i++
			}

			if i+1 < len(parts) {
				builder.WriteString(separator)
				builder.WriteString(parts[i+1])
				i++
			}

			current = builder.String()
		}
		result = append(result, current)
	}

	return result
}

func extractKeyValue(part string) (string, string) {
	if idx := strings.IndexByte(part, ':'); idx >= 0 {
		key := strings.TrimSpace(strings.ToUpper(part[:idx]))
		value := strings.TrimSpace(part[idx+1:])
		if value == "" {
			return key, key
		}
		return key, value
	}
	// No colon found
	key := strings.TrimSpace(strings.ToUpper(part))
	return key, key
}
