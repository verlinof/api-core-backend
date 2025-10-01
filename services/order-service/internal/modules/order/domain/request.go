package domain

import (
	shareddomain "monorepo/services/order-service/pkg/shared/domain"
)

// RequestOrder model
type RequestOrder struct {
	ID    int `json:"id"`
	Field string `json:"field"`
}

// Deserialize to db model
func (r *RequestOrder) Deserialize() (res shareddomain.Order) {
	res.Field = r.Field
	return
}
