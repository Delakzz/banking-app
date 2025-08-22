package models

import "time"

type AuditInfo struct {
	createdAt      time.Time
	lastModifiedAt time.Time
}

type BankAccount struct {
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
	AuditInfo
}
