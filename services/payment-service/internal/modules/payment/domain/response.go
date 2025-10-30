package domain

import (
	shareddomain "payment-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponsePaymentList model
type ResponsePaymentList struct {
	Meta candishared.Meta  `json:"meta"`
	Data []ResponsePayment `json:"data"`
}

// ResponsePayment model
type ResponsePayment struct {
	ID        int    `json:"id"`
	Field     string `json:"field"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponsePayment) Serialize(source *shareddomain.Payment) {
	r.ID = source.ID
	r.Field = source.Field
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}

type CreateTransactionResponse struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
	OrderID     string `json:"order_id"`
}
