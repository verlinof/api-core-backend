package domain

import (
	shareddomain "payment-service/pkg/shared/domain"
)

// RequestProvider model
type RequestProvider struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	IsActive bool   `json:"isActive"`
}

// Deserialize to db model
func (r *RequestProvider) Deserialize() (res shareddomain.PaymentProvider) {
	res.Name = r.Name
	res.Code = r.Code
	res.IsActive = r.IsActive
	return
}
