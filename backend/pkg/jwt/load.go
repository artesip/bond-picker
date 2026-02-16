package jwt

import (
	"crypto"
	"log/slog"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func LoadKey(path string, logger *slog.Logger) crypto.PrivateKey {
	file, err := os.ReadFile(path)
	if err != nil {
		logger.Error("jwt load key error", err.Error())
		return nil
	}

	key, err := jwt.ParseEdPrivateKeyFromPEM(file)
	if err != nil {
		return nil
	}

	return key
}
