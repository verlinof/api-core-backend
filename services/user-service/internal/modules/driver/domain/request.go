package domain

import (
	shareddomain "monorepo/services/user-service/pkg/shared/domain"
)

// RequestDriver model
type RequestDriver struct {
	ID    int    `json:"id"`
	Field string `json:"field"`
}

// Deserialize to db model
func (r *RequestDriver) Deserialize() (res shareddomain.Driver) {
	res.Field = r.Field
	return
}
