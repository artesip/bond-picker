package jwt

import (
	"crypto/ed25519"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func LoadKeys(path string) (ed25519.PrivateKey, ed25519.PublicKey, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("jwt load key error: %w", err)
	}

	key, err := jwt.ParseEdPrivateKeyFromPEM(file)
	if err != nil {
		return nil, nil, fmt.Errorf("parse key error: %w", err)
	}

	privateKey, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, nil, fmt.Errorf("parse key error: %w", err)
	}

	return privateKey, privateKey.Public().(ed25519.PublicKey), nil
}
