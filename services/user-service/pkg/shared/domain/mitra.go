package domain

import (
	"time"
)

// Mitra model
type Mitra struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	UserID    int       `gorm:"column:user_id;type:int;uniqueIndex" json:"user_id"`
	Type      string    `gorm:"column:type;type:varchar(255)" json:"type"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of Mitra model
func (Mitra) TableName() string {
	return "mitras"
}
