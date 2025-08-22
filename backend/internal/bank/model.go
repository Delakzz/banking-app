package bank

import (
	"banking-app/backend/internal/customer"
)

type Bank struct {
	ID       int64               `json:"id"`
	Name     string              `json:"name"`
	Customer []customer.Customer `json:"customers"`
}
