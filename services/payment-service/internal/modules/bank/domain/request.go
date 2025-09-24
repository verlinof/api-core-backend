package domain

import (
	shareddomain "payment-service/pkg/shared/domain"
)

// RequestBank model
type RequestBank struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	IsActive bool   `json:"isActive"`
}

// Deserialize to db model
func (r *RequestBank) Deserialize() (res shareddomain.PaymentBank) {
	res.Name = r.Name
	res.Code = r.Code
	res.IsActive = r.IsActive
	return
}
