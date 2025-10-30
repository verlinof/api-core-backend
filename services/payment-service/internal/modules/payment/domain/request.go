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
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
}

type ItemDetail struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Qty   int32  `json:"qty"`
}

type CreateOrderRequest struct {
	OrderID  string         `json:"order_id"`
	Amount   int64          `json:"amount"`
	Channel  string         `json:"channel"`
	MethodID int            `json:"method_id"`
	Customer CustomerDetail `json:"customer"`
	Items    []ItemDetail   `json:"items"`
}
