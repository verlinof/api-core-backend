package domain

import (
	"time"
)

// Category model
type PaymentCategory struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Code      string    `gorm:"column:code;type:varchar(100)" json:"code"`
	Sequence  int       `gorm:"column:sequence" json:"sequence"`
	IsActive  bool      `gorm:"column:is_active" json:"is_active"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	Method []PaymentMethod `gorm:"foreignKey:CategoryID;references:ID" json:"method,omitempty"`
}

// TableName return table name of Category model
func (PaymentCategory) TableName() string {
	return "payment_category"
}
