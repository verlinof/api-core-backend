package domain

import "github.com/golangid/candi/candishared"

// FilterMitra model
type FilterMitra struct {
	candishared.Filter
	ID        *int     `json:"id"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Preloads  []string `json:"-"`
}
