package validate

import (
	"fmt"
	"reflect"

	"github.com/anthanhphan/saturday/tag"
)

// IsZero determines whether the given value is considered "zero" in Go.
// It checks if a slice is empty or if a value is the zero value for its type.
//
// Parameters:
//   - in: The value to check
//
// Returns:
//   - bool: true if the value is zero; false otherwise
//
// Examples:
//
//	IsZero("")        // returns true
//	IsZero(0)         // returns true
//	IsZero([]int{})   // returns true
//	IsZero("hello")   // returns false
//	IsZero([]int{1})  // returns false
func IsZero(in any) bool {
	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		return v.Len() == 0
	default:
		return !v.IsValid() || v.IsZero()
	}
}

// IsPrimaryKeyNonZero checks if the primary key of a struct is non-zero.
// It inspects the struct's fields to find a field tagged as a primary key ("gorm:primaryKey")
// and validates its value.
//
// Parameters:
//   - in: The struct to check
//
// Returns:
//   - error: An error if the primary key is missing or its value is zero, nil otherwise
//
// Examples:
//
//	type User struct {
//	    ID   int    `gorm:"primaryKey"`
//	    Name string
//	}
//
//	user1 := User{ID: 0, Name: "John"}
//	IsPrimaryKeyNonZero(user1) // returns error: "primary key cannot be zero"
//
//	user2 := User{ID: 1, Name: "John"}
//	IsPrimaryKeyNonZero(user2) // returns nil
//
//	type InvalidStruct struct {
//	    Name string
//	}
//	invalid := InvalidStruct{Name: "John"}
//	IsPrimaryKeyNonZero(invalid) // returns error: "no primary key found"
func IsPrimaryKeyNonZero(in any) error {
	v := reflect.ValueOf(in)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagParsed := tag.Parse(field.Tag.Get("gorm"), ";")
		if _, ok := tagParsed["PRIMARYKEY"]; ok {
			if IsZero(v.Field(i).Interface()) {
				return fmt.Errorf("primary key cannot be zero")
			}
			return nil
		}
	}
	return fmt.Errorf("no primary key found")
}

// HasNonZeroExcludingKeys checks if a map contains any non-zero value excluding specified keys to skip.
//
// Parameters:
//   - m: The map to check
//   - skipKeys: A set of keys to skip during the check
//
// Returns:
//   - bool: true if any key has a non-zero value and is not in the skip list, false otherwise
//
// Examples:
//
//	m := map[string]any{
//	    "id":    1,
//	    "name":  "",
//	    "age":   0,
//	    "email": "user@example.com",
//	}
//
//	skip := map[string]any{"age": nil}
//	HasNonZeroExcludingKeys(m, skip) // returns true (due to "id" and "email")
//
//	skip = map[string]any{"id": nil, "email": nil}
//	HasNonZeroExcludingKeys(m, skip) // returns false
func HasNonZeroExcludingKeys(m map[string]any, skipKeys map[string]any) bool {
	for k, v := range m {
		if _, skip := skipKeys[k]; !skip && !IsZero(v) {
			return true
		}
	}
	return false
}
