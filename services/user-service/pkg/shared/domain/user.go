package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID             int       `gorm:"column:id;primary_key" json:"id"`
	Name           string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Email          string    `gorm:"column:email;type:varchar(255);uniqueIndex" json:"email"`
	Password       string    `gorm:"column:password;type:varchar(255)" json:"-"`
	ProfilePicture string    `gorm:"column:profile_picture;type:text;default:null" json:"profile_picture"`
	PhoneNumber    string    `gorm:"column:phone_number;type:varchar(255);default:null" json:"phone_number"`
	Address        string    `gorm:"column:address;type:varchar(255);default:null" json:"address"`
	IsVerified     bool      `gorm:"column:is_verified;type:boolean;default:false" json:"is_verified"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of User model
func (User) TableName() string {
	return "users"
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
