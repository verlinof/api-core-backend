package domain

import (
	shareddomain "payment-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseBankList model
type ResponseBankList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseBank   `json:"data"`
}

// ResponseBank model
type ResponseBank struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseBank) Serialize(source *shareddomain.PaymentBank) {
	r.ID = source.ID
	r.Name = source.Name
	r.Code = source.Code
	r.IsActive = source.IsActive
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
