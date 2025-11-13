package domain

import (
	"time"
)

// Checkout model
type Checkout struct {
	ID         int    `gorm:"column:id;primary_key" json:"id"`
	Field      string    `gorm:"column:field;type:varchar(255)" json:"field"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of Checkout model
func (Checkout) TableName() string {
	return "checkouts"
}

