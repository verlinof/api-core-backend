package domain

import "github.com/golangid/candi/candishared"

// FilterUser model
type FilterUser struct {
	candishared.Filter
	ID        *int     `json:"id"`
	Email     *string  `json:"email"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Preloads  []string `json:"-"`
}
