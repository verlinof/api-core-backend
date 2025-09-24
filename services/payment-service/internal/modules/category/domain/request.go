package domain

import (
	shareddomain "payment-service/pkg/shared/domain"
)

// RequestCategory model
type RequestCategory struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Sequence int    `json:"sequence"`
	IsActive bool   `json:"isActive"`
}

// Deserialize to db model
func (r *RequestCategory) Deserialize() (res shareddomain.PaymentCategory) {
	res.Name = r.Name
	res.Code = r.Code
	res.Sequence = r.Sequence
	res.IsActive = r.IsActive
	return
}
