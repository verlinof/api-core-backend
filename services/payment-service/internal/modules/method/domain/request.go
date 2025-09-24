package domain

import (
	shareddomain "payment-service/pkg/shared/domain"
)

// RequestMethod model
type RequestMethod struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	CategoryID  int    `json:"categoryId"`
	BankID      int    `json:"bankId"`
	ProviderID  int    `json:"providerId"`
	IsActive    bool   `json:"isActive"`
}

// Deserialize to db model
func (r *RequestMethod) Deserialize() (res shareddomain.PaymentMethod) {
	res.Code = r.Code
	res.Name = r.Name
	res.Description = r.Description
	res.IconURL = r.IconURL
	res.CategoryID = r.CategoryID
	res.BankID = r.BankID
	res.ProviderID = r.ProviderID
	res.IsActive = r.IsActive
	return
}
