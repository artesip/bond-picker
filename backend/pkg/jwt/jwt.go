package jwt

import (
	"backend/internal/domain"
	"backend/pkg/cookie"
	"crypto"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const iss = "bond-picker-auth"

func GenerateToken(key crypto.PrivateKey, userId domain.UUID) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodEdDSA,
		jwt.MapClaims{
			"iss": iss,
			"sub": userId,
			"exp": time.Now().Add(time.Duration(cookie.TokenAge) * time.Second).Unix(),
		})

	token, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}
	return token, nil
}
