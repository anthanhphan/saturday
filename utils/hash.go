package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword securely hashes a plain text password using bcrypt.
// It uses the default cost factor for the hashing algorithm.
//
// Parameters:
//   - password: The plain text password to hash
//
// Returns:
//   - *string: A pointer to the hashed password string
//   - error: Any error that occurred during hashing, or nil if successful
//
// Examples:
//
//	hashedPwd, err := HashPassword("mypassword123")
//	if err != nil {
//	    // Handle error
//	}
func HashPassword(password string) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hashedPasswordStr := string(hashedPassword)
	return &hashedPasswordStr, nil
}

// CompareHashAndPassword checks if a plain text password matches its hashed version.
// It uses bcrypt's comparison function to safely verify the password.
//
// Parameters:
//   - hashedPassword: The hashed password to compare against
//   - password: The plain text password to verify
//
// Returns:
//   - bool: true if passwords match, false otherwise
//
// Examples:
//
//	isValid := CompareHashAndPassword(hashedPassword, "mypassword123")
//	if isValid {
//	    // Password is correct
//	}
func CompareHashAndPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
