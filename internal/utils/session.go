package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSessionID() string {
	b := make([]byte, 16) // 128 бит
	rand.Read(b)
	return hex.EncodeToString(b)
}
