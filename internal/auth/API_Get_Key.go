package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {

	key := headers.Get("Authorization")
	if key == "" {
		return "", errors.New("No api key in the headers")
	}

	if v, ok := strings.CutPrefix(key, "ApiKey "); ok {
		key = v
	}
	key = strings.TrimSpace(key)
	return key, nil
}
