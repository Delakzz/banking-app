package database

import (
	"banking-app/backend/internal/bank"
	"banking-app/backend/internal/customer"
)

// Database represents the overall database structure with collections
type Database struct {
	Banks     []bank.Bank         `json:"banks"`
	Customers []customer.Customer `json:"customers"`
}

// Note: Both banks and customers are stored in database.json
// Banks can have multiple customers (stored as customer IDs)
// Customers can only belong to one bank (stored as bank ID)
