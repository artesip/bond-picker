package hash

import (
	"crypto/hmac"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	hashTime = 64 * 102
	memory   = 3
	threads  = 4
	keyLen   = 64
)

func HashPassword(password string) (hash, salt string, err error) {
	saltBytes := make([]byte, 16)
	if _, err := rand.Read(saltBytes); err != nil {
		return "", "", fmt.Errorf("failed to generate salt: %w", err)
	}

	hashBytes := argon2.IDKey([]byte(password), saltBytes, hashTime, memory, threads, keyLen)

	return hex.EncodeToString(hashBytes), hex.EncodeToString(saltBytes), nil
}

func VerifyPassword(password, hashStr, saltStr string) bool {
	hashBytes, _ := hex.DecodeString(hashStr)
	saltBytes, _ := hex.DecodeString(saltStr)

	computed := argon2.IDKey([]byte(password), saltBytes, hashTime, memory, threads, keyLen)

	return hmac.Equal(computed, hashBytes)
}
