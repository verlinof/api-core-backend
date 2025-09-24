package domain

import (
	"time"
)

// Payment Bank Model
type PaymentBank struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Code      string    `gorm:"column:code;type:varchar(100)" json:"code"`
	IsActive  bool      `gorm:"column:is_active" json:"is_active"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (PaymentBank) TableName() string {
	return "payment_bank"
}
