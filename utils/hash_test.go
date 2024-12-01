package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "mySecretPassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if hashedPassword == nil || *hashedPassword == "" {
		t.Fatalf("Expected a hashed password, got nil or empty string")
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "mySecretPassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !CompareHashAndPassword(*hashedPassword, password) {
		t.Fatalf("Expected password to be verified successfully")
	}

	wrongPassword := "wrongPassword"
	if CompareHashAndPassword(*hashedPassword, wrongPassword) {
		t.Fatalf("Expected password verification to fail with wrong password")
	}
}
