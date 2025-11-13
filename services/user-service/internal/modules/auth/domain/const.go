package domain

import "errors"

var (
	// ErrUnexpectedSigningMethod var
	ErrUnexpectedSigningMethod = errors.New("Unexpected signing method")
	// ErrTokenSignatureInvalid var
	ErrTokenSignatureInvalid = errors.New("Token signature is invalid")
	// ErrTokenFormat var
	ErrTokenFormat = errors.New("Invalid token format")
	// ErrTokenExpired var
	ErrTokenExpired = errors.New("Token is expired")
	// ErrTokenInvalid var
	ErrTokenInvalid = errors.New("Token is invalid")
	// ErrTokenNotValidYet var
	ErrTokenNotValidYet = errors.New("Token not valid yet")
	// ErrUserNotFound var
	ErrUserNotFound = errors.New("User not found")
	// ErrInvalidCredentials var
	ErrInvalidCredentials = errors.New("Invalid credentials")
	// ErrTokenNotFound var
	ErrTokenRevoked = errors.New("Token revoked")
)

const (
	RedisTokenConst = "user-service-token"
)
