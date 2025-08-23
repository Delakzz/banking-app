package bank

// Bank represents a simple banking institution
type Bank struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// NewBank creates a new bank instance
func NewBank(id int64, name, code string) *Bank {
	return &Bank{
		ID:   id,
		Name: name,
		Code: code,
	}
}
