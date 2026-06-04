package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header not found")
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("authorization header is not a Bearer token")
	}

	token := strings.TrimSpace(authHeader[len(prefix):])
	if token == "" {
		return "", errors.New("token is empty")
	}

	return token, nil
}

func MakeRefreshToken() (string, error) {

	b := make([]byte, 256)

	if _, err := rand.Read(b); err != nil {
		return "", errors.New("Something went wrong with rand.Read")
	}

	token := hex.EncodeToString(b)

	return token, nil
}
