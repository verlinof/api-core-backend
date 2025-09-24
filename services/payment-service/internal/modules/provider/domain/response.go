package domain

import (
	shareddomain "payment-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseProviderList model
type ResponseProviderList struct {
	Meta candishared.Meta   `json:"meta"`
	Data []ResponseProvider `json:"data"`
}

// ResponseProvider model
type ResponseProvider struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseProvider) Serialize(source *shareddomain.PaymentProvider) {
	r.ID = source.ID
	r.Name = source.Name
	r.Code = source.Code
	r.IsActive = source.IsActive
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
