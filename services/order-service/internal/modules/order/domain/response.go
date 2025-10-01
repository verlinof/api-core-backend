package domain

import (
	shareddomain "monorepo/services/order-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseOrderList model
type ResponseOrderList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseOrder   `json:"data"`
}

// ResponseOrder model
type ResponseOrder struct {
	ID        int `json:"id"`
	Field     string `json:"field"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseOrder) Serialize(source *shareddomain.Order) {
	r.ID = source.ID
	r.Field = source.Field
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
