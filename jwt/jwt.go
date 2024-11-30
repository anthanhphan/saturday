package jwt

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	gojwt "github.com/golang-jwt/jwt/v5"
)

type Jwt interface {
	Generate(payload *Payload, expiry int64) (*string, error)
	Validate(tokenString string) (*Payload, error)
}

type jwt struct {
	privateKey []byte
	publicKey  []byte
}

var _ Jwt = (*jwt)(nil)

// NewJwt is a constructor function to initialize a new JWT struct
func NewJwt(privateKey string, publicKey string) Jwt {
	return &jwt{
		privateKey: []byte(privateKey),
		publicKey:  []byte(publicKey),
	}
}

// Generate creates a JWT token with the given payload and expiry time
func (j *jwt) Generate(payload *Payload, expiry int64) (*string, error) {
	if payload == nil {
		return nil, fmt.Errorf("payload cannot be nil")
	}

	// Validate the payload
	if err := validator.New().Struct(payload); err != nil {
		return nil, fmt.Errorf("invalid payload: %v", err)
	}

	// Create a new JWT token with the specified claims
	token := gojwt.NewWithClaims(gojwt.SigningMethodRS256, gojwt.MapClaims{
		"user_id": payload.UserId,
		"exp":     time.Now().Add(time.Second * time.Duration(expiry)).Unix(),
		"iat":     time.Now().Unix(),
	})

	// Parse the RSA private key for signing the token
	key, err := gojwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Sign the token using the private key
	tokenString, err := token.SignedString(key)
	if err != nil {
		return nil, fmt.Errorf("failed to sign token: %w", err)
	}

	return &tokenString, nil
}

// Validate verifies the JWT token and extracts the payload
func (j *jwt) Validate(tokenString string) (*Payload, error) {
	// Parse the RSA private key for verifying the token
	key, err := gojwt.ParseRSAPublicKeyFromPEM(j.publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	// Parse and validate the JWT token
	jwtToken, err := gojwt.Parse(tokenString, func(jwtToken *gojwt.Token) (interface{}, error) {
		// Ensure the token is signed with the correct method
		if _, ok := jwtToken.Method.(*gojwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	// Check if the token is valid
	if !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract and verify the token claims
	claims, ok := jwtToken.Claims.(gojwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse claims: %w", err)
	}

	// Missing user_id or role in the token claims
	if claims["user_id"] == nil {
		return nil, fmt.Errorf("missing user_id in token claims")
	}
	// Return the payload extracted from the claims
	return &Payload{
		UserId: int64(claims["user_id"].(float64)),
	}, nil
}
