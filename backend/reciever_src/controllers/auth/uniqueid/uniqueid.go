package uniqueid

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// URL-safe charset: only alphanumeric characters and hyphens/underscores
// Removed: ! @ # $ % ^ & * ( ) + = < > ? / { } [ ] ~
// These characters cause issues in URLs and query parameters
const charset = "abcdefghijklmnopqrstuvwxyz0123456789-_"
const length = 20

// Generate creates a URL-safe unique identifier
func Generate() (string, error) {
	b := make([]byte, length)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[num.Int64()]
	}
	return string(b), nil
}

// Validate checks if the unique ID is valid
// Requirements: exactly 20 chars, only allowed chars (a-z, 0-9, -, _), no spaces
func Validate(id string) error {
	if len(id) != 20 {
		return errors.New("unique id must be exactly 20 characters")
	}

	for _, r := range id {
		// Check for spaces
		if r == ' ' {
			return errors.New("unique id cannot contain spaces")
		}

		// Check if character is in allowed set
		if (r >= 'a' && r <= 'z') ||
			(r >= '0' && r <= '9') ||
			r == '-' ||
			r == '_' {
			continue
		}

		return errors.New("unique id contains invalid characters (only a-z, 0-9, -, _ allowed)")
	}

	return nil
}
