package bank

// Bank represents a simple banking institution
type Bank struct {
	ID        int64   `json:"id"`
	Username  string  `json:"username,omitempty"`
	Password  string  `json:"password,omitempty"`
	Name      string  `json:"name,omitempty"`
	Customers []int64 `json:"customers,omitempty"` // Store customer IDs instead of full objects
}

// NewBank creates a new bank instance
func NewBank(id int64, username, password, name string) *Bank {
	return &Bank{
		ID:        id,
		Username:  username,
		Password:  password,
		Name:      name,
		Customers: []int64{},
	}
}
