package bank

type Bank struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"userid"`
	Name   string `json:"name,omitempty"`
}

func NewBank(id, userID int64, name string) *Bank {
	return &Bank{
		ID:     id,
		UserID: userID,
		Name:   name,
	}
}
