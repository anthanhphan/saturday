package utils

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
)

// RandString generates a random hex string of the specified length.
// It uses crypto/rand for secure random number generation.
//
// Parameters:
//   - length: The desired length of the output hex string (actual output will be 2x this length)
//
// Returns:
//   - string: The generated random hex string
//   - error: Any error that occurred during random number generation
//
// Examples:
//
//	RandString(4)    // returns "a1b2c3d4", nil
//	RandString(2)    // returns "f5e9", nil
//	RandString(0)    // returns "", nil
func RandString(length int) (string, error) {
	randStr := make([]byte, length)
	if _, err := rand.Read(randStr); err != nil {
		return "", err
	}
	return hex.EncodeToString(randStr), nil
}

// RandInt64 generates a cryptographically secure random 64-bit integer.
// It uses crypto/rand to generate random bytes which are then converted to an int64.
//
// Returns:
//   - int64: A random 64-bit integer
//   - error: Any error that occurred during random number generation
//
// Examples:
//
//	RandInt64()    // returns 8674665223082153551, nil
//	RandInt64()    // returns -6129484611666145821, nil
func RandInt64() (int64, error) {
	// Create an 8-byte buffer for a 64-bit number
	var buf [8]byte

	// Fill buffer with random bytes
	if _, err := rand.Read(buf[:]); err != nil {
		return 0, err
	}

	// Convert bytes to int64 using binary.BigEndian
	return int64(binary.BigEndian.Uint64(buf[:])), nil
}
