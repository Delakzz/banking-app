package customer

import (
	"banking-app/backend/internal/bank_account"
)

type Customer struct {
	ID       int64                      `json:"id"`
	Name     string                     `json:"name"`
	Accounts []bank_account.BankAccount `json:"accounts"`
}
