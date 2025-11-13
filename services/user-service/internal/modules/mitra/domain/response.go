package domain

import (
	"github.com/golangid/candi/candishared"
)

// ResponseMitraList model
type ResponseMitraList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseMitra  `json:"data"`
}

// ResponseMitra model
type ResponseMitra struct {
	ID        int    `json:"id"`
	Field     string `json:"field"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
