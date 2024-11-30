package jwt

import (
	"testing"
	"time"

	"github.com/anthanhphan/saturday/utils"
	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAndValidate(t *testing.T) {
	// Generate RSA keys for the test
	rsaKeys, err := utils.GenerateRsaKeyPair()
	assert.NoError(t, err)

	jwtService := NewJwt(rsaKeys.PrivateKey, rsaKeys.PublicKey)

	// Test data for generating a token
	payload := &Payload{
		UserId: 123,
	}

	// Generate token
	tokenString, err := jwtService.Generate(payload, 3600) // Expire in 1 hour
	assert.NoError(t, err)
	assert.NotNil(t, tokenString)

	// Validate token after a short delay (to avoid time issues)
	time.Sleep(1 * time.Second)
	validatedPayload, err := jwtService.Validate(*tokenString)
	assert.NoError(t, err)
	assert.NotNil(t, validatedPayload)

	// Ensure the payload matches the original
	assert.Equal(t, payload.UserId, validatedPayload.UserId)
}

func TestValidateInvalidToken(t *testing.T) {
	// Generate RSA keys for the test
	rsaKeys, err := utils.GenerateRsaKeyPair()
	assert.NoError(t, err)

	jwtService := NewJwt(rsaKeys.PrivateKey, rsaKeys.PublicKey)

	// Invalid token test case
	invalidToken := "invalid.token.string"

	_, err = jwtService.Validate(invalidToken)
	assert.Error(t, err)
}

func TestGenerateWithInvalidPrivateKey(t *testing.T) {
	// Use invalid private key
	jwtService := NewJwt("invalid-private-key", "valid-public-key")

	payload := &Payload{
		UserId: 123,
	}

	// Try to generate token with an invalid private key
	_, err := jwtService.Generate(payload, 3600)
	assert.Error(t, err)
}

func TestValidateWithInvalidPublicKey(t *testing.T) {
	// Generate RSA keys for the test
	rsaKeys, err := utils.GenerateRsaKeyPair()
	assert.NoError(t, err)

	jwtService := NewJwt(rsaKeys.PrivateKey, "invalid-public-key")

	payload := &Payload{
		UserId: 123,
	}

	// Generate a valid token
	tokenString, err := jwtService.Generate(payload, 3600)
	assert.NoError(t, err)
	assert.NotNil(t, tokenString)

	// Validate token with an invalid public key
	_, err = jwtService.Validate(*tokenString)
	assert.Error(t, err)
}

func TestValidateTokenWithExpiredClaims(t *testing.T) {
	// Generate RSA keys for the test
	rsaKeys, err := utils.GenerateRsaKeyPair()
	assert.NoError(t, err)

	jwtService := NewJwt(rsaKeys.PrivateKey, rsaKeys.PublicKey)

	payload := &Payload{
		UserId: 123,
	}

	// Generate token with immediate expiry
	tokenString, err := jwtService.Generate(payload, -1) // Expired token
	assert.NoError(t, err)
	assert.NotNil(t, tokenString)

	// Validate token after expiry
	time.Sleep(1 * time.Second)
	_, err = jwtService.Validate(*tokenString)
	assert.Error(t, err, "expected error for expired token")
}

func TestValidateMissingUserIDClaim(t *testing.T) {
	// Generate RSA keys for the test
	rsaKeys, err := utils.GenerateRsaKeyPair()
	assert.NoError(t, err)

	jwtService := NewJwt(rsaKeys.PrivateKey, rsaKeys.PublicKey)

	// Create a token with missing user_id claim
	token := gojwt.NewWithClaims(gojwt.SigningMethodRS256, gojwt.MapClaims{
		"role": "admin",
		"exp":  time.Now().Add(time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	})

	// Sign the token
	key, _ := gojwt.ParseRSAPrivateKeyFromPEM([]byte(rsaKeys.PrivateKey))
	tokenString, err := token.SignedString(key)
	assert.NoError(t, err)

	// Validate the token with missing user_id claim
	_, err = jwtService.Validate(tokenString)
	assert.Error(t, err, "expected error for missing user_id claim")
}

func TestValidateTokenWithWrongSigningMethod(t *testing.T) {
	// Generate RSA keys for the test
	rsaKeys, err := utils.GenerateRsaKeyPair()
	assert.NoError(t, err)

	jwtService := NewJwt(rsaKeys.PrivateKey, rsaKeys.PublicKey)

	payload := &Payload{
		UserId: 123,
	}

	// Generate a token
	tokenString, err := jwtService.Generate(payload, 3600)
	assert.NoError(t, err)

	// Dereference the pointer to get the actual string value and tamper with the token
	tamperedToken := (*tokenString)[:len(*tokenString)-1] + "not-correct" // Modify the last character

	// Validate tampered token
	_, err = jwtService.Validate(tamperedToken)
	assert.Error(t, err, "expected error for tampered token")
}

func TestGenerateWithInvalidPrivateKeyFormat(t *testing.T) {
	jwtService := NewJwt("invalid-private-key-format", "valid-public-key")

	payload := &Payload{
		UserId: 123,
	}

	_, err := jwtService.Generate(payload, 3600)
	assert.Error(t, err, "expected error due to invalid private key format")
}

func TestValidateWithInvalidPublicKeyFormat(t *testing.T) {
	jwtService := NewJwt("valid-private-key", "invalid-public-key-format")

	payload := &Payload{
		UserId: 123,
	}

	_, err := jwtService.Generate(payload, 3600)
	assert.Error(t, err)
}

func TestGenerateWithNilPayload(t *testing.T) {
	rsaKeys, err := utils.GenerateRsaKeyPair()
	assert.NoError(t, err)

	jwtService := NewJwt(rsaKeys.PrivateKey, rsaKeys.PublicKey)

	_, err = jwtService.Generate(nil, 3600) // Nil payload
	assert.Error(t, err, "payload cannot be nil")
}

func TestGenerateWithEmptyClaims(t *testing.T) {
	rsaKeys, err := utils.GenerateRsaKeyPair()
	assert.NoError(t, err)

	jwtService := NewJwt(rsaKeys.PrivateKey, rsaKeys.PublicKey)

	// Create a payload with empty claims
	payload := &Payload{} // Empty claims

	_, err = jwtService.Generate(payload, 3600)
	assert.Error(t, err)
}
