package domain

import "github.com/golang-jwt/jwt/v5"

const (

	// HS256 const
	HS256 = "HS256"

	// RS256 const
	RS256 = "RS256"
)

// Claim for token claim data
type CustomClaims struct {
	UserID string `json:"id_user"`
	jwt.RegisteredClaims
}
