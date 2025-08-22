package models

type Bank struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Customer    []Customer `json:"customers"`