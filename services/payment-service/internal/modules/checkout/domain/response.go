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
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsActive  bool   `json:"isActive"`
	Sequence  int    `json:"sequence"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Method []shareddomain.PaymentMethod `json:"method,omitempty"`
}

// Serialize from db model
func (r *ResponseMethod) Serialize(source *shareddomain.PaymentCategory) {
	r.ID = source.ID
	r.Name = source.Name
	r.Code = source.Code
	r.IsActive = source.IsActive
	r.Sequence = source.Sequence
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)

	// Method Relation
	r.Method = source.Method
}
