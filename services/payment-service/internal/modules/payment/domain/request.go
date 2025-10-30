package domain

import (
	shareddomain "payment-service/pkg/shared/domain"
)

// RequestPayment model
type RequestPayment struct {
	ID    int    `json:"id"`
	Field string `json:"field"`
}

// Deserialize to db model
func (r *RequestPayment) Deserialize() (res shareddomain.Payment) {
	res.Field = r.Field
	return
}

type CustomerDetail struct {
	FirstName string `json:"first_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
}

type CreateOrderRequest struct {
	OrderID  string         `json:"order_id"`
	Amount   int64          `json:"amount"`
	Customer CustomerDetail `json:"customer_detail"`
}
