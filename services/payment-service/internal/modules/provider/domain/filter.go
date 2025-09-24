package domain

import "github.com/golangid/candi/candishared"

// FilterProvider model
type FilterProvider struct {
	candishared.Filter
	ID        *int `json:"id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Preloads  []string `json:"-"`
}
