package domain

import (
	shareddomain "payment-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseCategoryList model
type ResponseCategoryList struct {
	Meta candishared.Meta   `json:"meta"`
	Data []ResponseCategory `json:"data"`
}

// ResponseCategory model
type ResponseCategory struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseCategory) Serialize(source *shareddomain.PaymentCategory) {
	r.ID = source.ID
	r.Name = source.Name
	r.Code = source.Code
	r.IsActive = source.IsActive
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
