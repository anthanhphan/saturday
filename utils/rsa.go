package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type RsaKey struct {
	PrivateKey string
	PublicKey  string
}

// GenerateRsaKeyPair generates RSA private and public keys and returns them as PEM-encoded strings.
// It uses a 2048-bit key size which provides a good balance of security and performance.
//
// Parameters:
//   - none
//
// Returns:
//   - RsaKey: Struct containing the generated key pair as PEM-encoded strings
//   - error: Any error that occurred during key generation or encoding
//
// Examples:
//
//	keys, err := GenerateRsaKeyPair()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(keys.PrivateKey) // prints PEM-encoded private key
//	fmt.Println(keys.PublicKey)  // prints PEM-encoded public key
func GenerateRsaKeyPair() (RsaKey, error) {
	var rsaKey RsaKey

	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return rsaKey, err
	}

	// Encode private key to PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	privateKeyBytes := pem.EncodeToMemory(privateKeyPEM)
	rsaKey.PrivateKey = string(privateKeyBytes)

	// Generate and encode public key to PEM format
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return rsaKey, err
	}
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicKeyBytes = pem.EncodeToMemory(publicKeyPEM)
	rsaKey.PublicKey = string(publicKeyBytes)

	return rsaKey, nil
}
