package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenSha() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(randomBytes)
	return hex.EncodeToString(hash[:])
}
