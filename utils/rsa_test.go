package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"reflect"
	"testing"
)

func TestGenerateRSAKeys(t *testing.T) {
	rsaKey, err := GenerateRsaKeyPair()
	if err != nil {
		t.Errorf("GenerateRSAKeys returned an error: %v", err)
	}

	// Decode private key PEM block
	privateKeyBlock, _ := pem.Decode([]byte(rsaKey.PrivateKey))
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		t.Errorf("Invalid private key PEM block")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		t.Errorf("Failed to parse private key: %v", err)
	}

	// Decode public key PEM block
	publicKeyBlock, _ := pem.Decode([]byte(rsaKey.PublicKey))
	if publicKeyBlock == nil || publicKeyBlock.Type != "RSA PUBLIC KEY" {
		t.Errorf("Invalid public key PEM block")
	}
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		t.Errorf("Failed to parse public key: %v", err)
	}
	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		t.Error("Failed to convert public key to *rsa.PublicKey")
	}

	// Check if the public key corresponds to the private key
	if !reflect.DeepEqual(&privateKey.PublicKey, publicKey) {
		t.Error("Public key does not correspond to the private key")
	}
}
