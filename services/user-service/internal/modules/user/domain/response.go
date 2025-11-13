package domain

import (
	"github.com/golangid/candi/candishared"
)

// ResponseUserList model
type ResponseUserList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseUser   `json:"data"`
}

// ResponseUser model
type ResponseUser struct {
	ID        int    `json:"id"`
	Field     string `json:"field"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
