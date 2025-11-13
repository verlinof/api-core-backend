package userservice

import "time"

type ClaimResponse struct {
	UserId    string        `json:"id_user"`
	ExpiresAt time.Duration `json:"exp"`
	IssuedAt  time.Duration `json:"iat"`
}
