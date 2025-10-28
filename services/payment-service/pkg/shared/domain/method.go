package domain

import (
	"time"
)

type PaymentMethod struct {
	ID          int              `gorm:"column:id;primary_key" json:"id"`
	Code        string           `gorm:"column:code;type:varchar(100)" json:"code"`
	Name        string           `gorm:"column:name;type:varchar(255)" json:"name"`
	Description string           `gorm:"column:description;type:text" json:"description"`
	Icon        string           `gorm:"column:icon;type:text" json:"icon"`
	CategoryID  int              `gorm:"column:category_id" json:"category_id"`
	BankID      int              `gorm:"column:bank_id" json:"bank_id"`
	ProviderID  int              `gorm:"column:provider_id" json:"provider_id"`
	IsActive    bool             `gorm:"column:is_active" json:"is_active"`
	CreatedAt   time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time        `gorm:"column:updated_at" json:"updated_at"`
	Category    *PaymentCategory `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`
	Bank        *PaymentBank     `gorm:"foreignKey:BankID;references:ID" json:"bank,omitempty"`
	Provider    *PaymentProvider `gorm:"foreignKey:ProviderID;references:ID" json:"provider,omitempty"`
}

func (PaymentMethod) TableName() string {
	return "payment_method"
}
