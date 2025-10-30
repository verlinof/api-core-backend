package domain

import "github.com/golangid/candi/candishared"

// FilterPayment model
type FilterPayment struct {
	candishared.Filter
	ID        *int `json:"id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Preloads  []string `json:"-"`
}
