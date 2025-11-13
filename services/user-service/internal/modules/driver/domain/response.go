package domain

import (
	shareddomain "monorepo/services/user-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseDriverList model
type ResponseDriverList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseDriver `json:"data"`
}

// ResponseDriver model
type ResponseDriver struct {
	ID        int    `json:"id"`
	Field     string `json:"field"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseDriver) Serialize(source *shareddomain.Driver) {
	r.ID = source.ID
	r.Field = source.Field
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
