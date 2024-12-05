package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extract api key from the headers of an HTTP Request
// Exmple:
// Authorization: ApiKey {insert apiKey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No Authorization header found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Invalid Authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Invalid Authorization first part of auth header")
	}
	return vals[1], nil
	
}