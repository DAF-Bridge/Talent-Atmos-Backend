package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateStateString generates a secure random state string for OAuth.
func GenerateStateString() string {
	// Define the length of the random bytes
	const byteLength = 16 // 16 bytes = 128 bits

	// Generate random bytes
	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("Failed to generate random state string: " + err.Error())
	}

	// Encode the random bytes to a URL-safe string
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(randomBytes)
}
