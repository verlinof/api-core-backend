package globalshared

import (
	"errors"
	"net/http"
	"strings"
)

// ExtractBearerToken extracts a bearer token from the Authorization header of an HTTP request.
func ExtractBearerToken(r *http.Request) (string, error) {
	// Get the Authorization header value
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header is missing")
	}

	// Split the header to separate the "Bearer" scheme from the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("Authorization header format must be Bearer <token>")
	}

	// Return the token part, trimming any extra whitespace
	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", errors.New("token is missing from Authorization header")
	}

	return token, nil
}
