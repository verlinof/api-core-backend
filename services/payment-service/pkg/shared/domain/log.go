package domain

import (
	"time"
)

type PaymentLog struct {
	ID         int              `gorm:"column:id;primary_key" json:"id"`
	OrderID    int              `gorm:"column:order_id" json:"order_id"`
	ProviderID int              `gorm:"column:provider_id" json:"provider_id"`
	URL        string           `gorm:"column:url;type:text" json:"url"`
	Request    string           `gorm:"column:request;type:text" json:"request"`
	Response   string           `gorm:"column:response;type:text" json:"response"`
	CreatedAt  time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time        `gorm:"column:updated_at" json:"updated_at"`
	Provider   *PaymentProvider `gorm:"foreignKey:ProviderID;references:ID" json:"provider,omitempty"`
}

func (PaymentLog) TableName() string {
	return "payment_log"
}
