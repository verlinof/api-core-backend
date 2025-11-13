package domain

import (
	shareddomain "payment-service/pkg/shared/domain"
)

// RequestCheckout model
type RequestCheckout struct {
	ID    int `json:"id"`
	Field string `json:"field"`
}

// Deserialize to db model
func (r *RequestCheckout) Deserialize() (res shareddomain.Checkout) {
	res.Field = r.Field
	return
}
