package security

import "encoding/base64"

const sessionKeyLength uint32 = 64

// Generates a random session key for use with the user store
func GenerateSessionKey() (string, error) {
	byteKey, err := generateSalt(sessionKeyLength)

	if err != nil {
		return "", err
	}

	// Casting to a string doesnt work here
	return base64.StdEncoding.EncodeToString(byteKey), nil
}
