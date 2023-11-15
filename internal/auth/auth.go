package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from
// the Headers of an HTTP reques
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	authorization := headers.Get("Authorization")
	if authorization == "" {
		return "", errors.New("No Authentication info found")
	}

	authorizations := strings.Split(authorization, " ")
	if len(authorizations) != 2 {
		return "", errors.New("No Authentication info found")
	}
	if authorizations[0] != "ApiKey" {
		return "", errors.New("No Authentication info found")
	}
	return authorizations[1], nil
}
