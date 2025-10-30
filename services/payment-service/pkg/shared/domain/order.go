package domain

import (
	"time"
)

type PaymentOrder struct {
	ID                int            `gorm:"column:id;primary_key" json:"id"`
	OrderID           string         `gorm:"column:order_id" json:"order_id"`
	MethodID          int            `gorm:"column:method_id" json:"method_id"`
	Amount            float64        `gorm:"column:amount;type:decimal(20,2)" json:"amount"`
	Status            string         `gorm:"column:status;type:varchar(50)" json:"status"`
	FraudStatus       string         `gorm:"column:fraud_status;type:varchar(50)" json:"fraud_status"`
	TransactionStatus string         `gorm:"column:transaction_status;type:varchar(50)" json:"transaction_status"`
	Channel           string         `gorm:"column:channel;type:channel_type" json:"channel"`
	OrderData         string         `gorm:"column:order_data;type:text" json:"order_data"`
	Method            *PaymentMethod `gorm:"foreignKey:MethodID;references:ID" json:"method,omitempty"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

func (PaymentOrder) TableName() string {
	return "payment_order"
}
