package kafka

import (
	"crypto/sha256"

	"github.com/xdg-go/scram"
)

// SHA256 provides a hash generator function using SHA-256 algorithm.
//
// Returns:
//   - scram.HashGeneratorFcn: A SHA-256 hash generator for SCRAM authentication
var SHA256 scram.HashGeneratorFcn = sha256.New

// scramClient implements SCRAM authentication mechanism.
//
// Fields:
//   - Client: The underlying SCRAM client
//   - ClientConversation: The current SCRAM conversation
//   - HashGeneratorFcn: The hash function generator to use
type scramClient struct {
	*scram.Client
	*scram.ClientConversation
	scram.HashGeneratorFcn
}

// Begin initializes a new SCRAM authentication conversation.
//
// Parameters:
//   - user: The username for authentication
//   - password: The password for authentication
//   - authId: The authorization ID (optional)
//
// Returns:
//   - error: An error if client initialization fails, nil otherwise
//
// Examples:
//
//	client.Begin("user", "pass", "")  // Basic auth
//	client.Begin("user", "pass", "authz-id")  // With authorization ID
func (s *scramClient) Begin(user, password, authId string) error {
	var err error
	s.Client, err = s.HashGeneratorFcn.NewClient(user, password, authId)
	if err != nil {
		return err
	}
	s.ClientConversation = s.Client.NewConversation()

	return nil
}

// Step advances the SCRAM authentication conversation.
//
// Parameters:
//   - challenge: The challenge string from the server
//
// Returns:
//   - string: The client's response to the challenge
//   - error: An error if the step fails, nil otherwise
func (s *scramClient) Step(challenge string) (string, error) {
	return s.ClientConversation.Step(challenge)
}

// Done checks if the SCRAM conversation is complete.
//
// Returns:
//   - bool: true if the conversation is complete, false otherwise
func (s *scramClient) Done() bool {
	return s.ClientConversation.Done()
}
