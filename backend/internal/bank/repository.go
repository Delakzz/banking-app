package bank

import (
	"fmt"
	"sync"
)

// Repository manages bank data storage and retrieval
type Repository struct {
	banks map[int64]*Bank
	mutex sync.RWMutex
	nextID int64
}

// NewRepository creates a new bank repository
func NewRepository() *Repository {
	return &Repository{
		banks: make(map[int64]*Bank),
		nextID: 1,
	}
}

// Create adds a new bank to the repository
func (r *Repository) Create(name, code string) (*Bank, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	// Check if code already exists
	for _, bank := range r.banks {
		if bank.Code == code {
			return nil, fmt.Errorf("bank with code %s already exists", code)
		}
	}
	
	bank := NewBank(r.nextID, name, code)
	r.banks[bank.ID] = bank
	r.nextID++
	
	return bank, nil
}

// GetByID retrieves a bank by its ID
func (r *Repository) GetByID(id int64) (*Bank, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	bank, exists := r.banks[id]
	if !exists {
		return nil, fmt.Errorf("bank with ID %d not found", id)
	}
	
	return bank, nil
}

// GetAll retrieves all banks
func (r *Repository) GetAll() []*Bank {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	banks := make([]*Bank, 0, len(r.banks))
	for _, bank := range r.banks {
		banks = append(banks, bank)
	}
	
	return banks
}

// Update modifies an existing bank
func (r *Repository) Update(id int64, name, code string) (*Bank, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	bank, exists := r.banks[id]
	if !exists {
		return nil, fmt.Errorf("bank with ID %d not found", id)
	}
	
	// Check if new code conflicts with existing banks
	for _, existingBank := range r.banks {
		if existingBank.ID != id && existingBank.Code == code {
			return nil, fmt.Errorf("bank with code %s already exists", code)
		}
	}
	
	bank.Name = name
	bank.Code = code
	
	return bank, nil
}

// Delete removes a bank by ID
func (r *Repository) Delete(id int64) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.banks[id]; !exists {
		return fmt.Errorf("bank with ID %d not found", id)
	}
	
	delete(r.banks, id)
	return nil
}
