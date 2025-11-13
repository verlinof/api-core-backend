package domain

import (
	shareddomain "monorepo/services/user-service/pkg/shared/domain"
)

// RequestAuth model
type RequestAuth struct {
	ID    int    `json:"id"`
	Field string `json:"field"`
}

// Deserialize to db model
func (r *RequestAuth) Deserialize() (res shareddomain.Auth) {
	res.Field = r.Field
	return
}
