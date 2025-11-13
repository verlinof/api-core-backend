package domain

import "github.com/golangid/candi/candishared"

// FilterDriver model
type FilterDriver struct {
	candishared.Filter
	ID        *int     `json:"id"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Preloads  []string `json:"-"`
}
