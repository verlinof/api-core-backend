package domain

import (
	"time"
)

type PaymentLog struct {
	ID         int              `gorm:"column:id;primary_key" json:"id"`
	OrderID    int              `gorm:"column:order_id" json:"order_id"`
	ProviderID int              `gorm:"column:provider_id" json:"provider_id"`
	StatusCode int              `gorm:"column:status_code" json:"status_code"`
	PaymentURL string           `gorm:"column:payment_url;type:text" json:"payment_url"`
	Request    string           `gorm:"column:request_data;type:text" json:"request_data"`
	Response   string           `gorm:"column:response_data;type:text" json:"response_data"`
	CreatedAt  time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time        `gorm:"column:updated_at" json:"updated_at"`
	Provider   *PaymentProvider `gorm:"foreignKey:ProviderID;references:ID" json:"provider,omitempty"`
}

func (PaymentLog) TableName() string {
	return "payment_log"
}
