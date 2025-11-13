package userservice

import "context"

// Userservice client abstract interface
type Userservice interface {
	// Add service client method
	VerifyToken(ctx context.Context, token string) (resp ClaimResponse, err error)
}
