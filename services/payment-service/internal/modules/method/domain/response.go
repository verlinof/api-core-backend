package domain

import (
	shareddomain "payment-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseMethodList model
type ResponseMethodList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseMethod `json:"data"`
}

// ResponseMethod model
type ResponseMethod struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	AdminFee    string `json:"adminFee"`
	CategoryID  int    `json:"categoryID"`
	BankID      int    `json:"bankID"`
	ProviderID  int    `json:"providerID"`
	IsActive    bool   `json:"isActive"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseMethod) Serialize(source *shareddomain.PaymentMethod) {
	r.ID = source.ID
	r.Code = source.Code
	r.Name = source.Name
	r.Description = source.Description
	r.Icon = source.Icon
	r.CategoryID = source.CategoryID
	r.BankID = source.BankID
	r.ProviderID = source.ProviderID
	r.IsActive = source.IsActive
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
