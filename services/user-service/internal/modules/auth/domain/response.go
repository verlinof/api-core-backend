package domain

import "github.com/golang-jwt/jwt/v5"

type ResponseClaims struct {
	UserID string `json:"userId" validate:"required"`
	jwt.RegisteredClaims
}

type ResponseGenerateToken struct {
	Token        string          `json:"token"`
	RefreshToken string          `json:"refreshToken"`
	Claim        *ResponseClaims `json:"claim"`
}

type ResponseLogin struct {
	UserID       string `json:"userId"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
