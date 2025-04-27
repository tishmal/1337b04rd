package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

type UUID [16]byte

// NewUUID generates a new UUID and returns it as a shorter base64 string.
func NewUUID() (string, error) {
	var uuid UUID
	_, err := rand.Read(uuid[:])
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %w", err)
	}
	uuid[6] &= 0x0F
	uuid[6] |= 0x40
	uuid[8] &= 0x3F
	uuid[8] |= 0x80

	// Return a Base64-encoded string representation of the UUID
	return base64.RawURLEncoding.EncodeToString(uuid[:]), nil
}

func ParseUUID(s string) (UUID, error) {
	// Decode the Base64-encoded string back to UUID
	bytes, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return UUID{}, fmt.Errorf("failed to decode UUID: %w", err)
	}

	var uuid UUID
	copy(uuid[:], bytes)
	return uuid, nil
}
