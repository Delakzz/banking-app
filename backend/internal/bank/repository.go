package bank

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Repository struct {
	filePath string
	mutex    sync.RWMutex
	nextID   int64
	banks    []*Bank // Cache for in-memory operations
}

func NewRepository(dataDir string) *Repository {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		panic(fmt.Sprintf("failed to create data directory: %v", err))
	}

	filePath := filepath.Join(dataDir, "database.json")

	repo := &Repository{
		filePath: filePath,
		nextID:   1,
		banks:    []*Bank{},
	}

	repo.loadDB()

	return repo
}

func (r *Repository) loadDB() {
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		return
	}

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		fmt.Printf("Warning: failed to read banks file: %v\n", err)
		return
	}

	var banks []*Bank
	if err := json.Unmarshal(data, &banks); err != nil {
		fmt.Printf("Warning: failed to parse banks file: %v\n", err)
		return
	}

	// find the highest ID to set nextID correctly
	for _, bank := range banks {
		if bank.ID >= r.nextID {
			r.nextID = bank.ID + 1
		}
	}

	r.banks = banks
}

func (r *Repository) saveData() error {
	data, err := json.MarshalIndent(r.banks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal banks data: %w", err)
	}

	if err := os.WriteFile(r.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write banks file: %w", err)
	}

	return nil
}

func (r *Repository) Create(username, password, name string) (*Bank, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	bank := NewBank(r.nextID, username, password, name)
	r.banks = append(r.banks, bank)
	r.nextID++

	if err := r.saveData(); err != nil {
		return nil, fmt.Errorf("failed to save bank data: %w", err)
	}

	return bank, nil
}

func (r *Repository) GetByID(id int64) (*Bank, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, bank := range r.banks {
		if bank.ID == id {
			return bank, nil
		}
	}

	return nil, fmt.Errorf("bank with ID %d not found", id)
}

func (r *Repository) GetByName(name string) (*Bank, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	name = strings.ToLower(name)

	for _, bank := range r.banks {
		if bankName := strings.ToLower(bank.Name); bankName == name {
			return bank, nil
		}
	}

	return nil, fmt.Errorf("bank with name '%s' not found", name)
}

func (r *Repository) GetAll() []*Bank {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// Return a copy to avoid external modification
	banks := make([]*Bank, len(r.banks))
	copy(banks, r.banks)

	return banks
}

func (r *Repository) Update(id int64, username, password, name string) (*Bank, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Find and update bank
	for _, bank := range r.banks {
		if bank.ID == id {
			if username != "" {
				bank.Username = username
			}

			if password != "" {
				bank.Password = password
			}

			if name != "" {
				bank.Name = name
			}

			// Save updated data
			if err := r.saveData(); err != nil {
				return nil, fmt.Errorf("failed to save bank data: %w", err)
			}

			return bank, nil
		}
	}

	return nil, fmt.Errorf("bank with ID %d not found", id)
}

func (r *Repository) Delete(id int64) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Find and remove bank
	for i, bank := range r.banks {
		if bank.ID == id {
			// Remove bank by slicing
			r.banks = append(r.banks[:i], r.banks[i+1:]...)

			// Save updated data
			if err := r.saveData(); err != nil {
				return fmt.Errorf("failed to save bank data: %w", err)
			}

			return nil
		}
	}

	return fmt.Errorf("bank with ID %d not found", id)
}

func (r *Repository) AddCustomer(bankID, customerID int64) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Find bank and add customer
	for _, bank := range r.banks {
		if bank.ID == bankID {
			// Check if customer already exists
			for _, existingCustomerID := range bank.Customers {
				if existingCustomerID == customerID {
					return fmt.Errorf("customer %d already exists in bank %d", customerID, bankID)
				}
			}

			bank.Customers = append(bank.Customers, customerID)

			// Save updated data
			if err := r.saveData(); err != nil {
				return fmt.Errorf("failed to save bank data: %w", err)
			}

			return nil
		}
	}

	return fmt.Errorf("bank with ID %d not found", bankID)
}

func (r *Repository) RemoveCustomer(bankID, customerID int64) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Find bank and remove customer
	for _, bank := range r.banks {
		if bank.ID == bankID {
			// Find and remove customer
			for i, existingCustomerID := range bank.Customers {
				if existingCustomerID == customerID {
					// Remove customer by slicing
					bank.Customers = append(bank.Customers[:i], bank.Customers[i+1:]...)

					// Save updated data
					if err := r.saveData(); err != nil {
						return fmt.Errorf("failed to save bank data: %w", err)
					}

					return nil
				}
			}

			return fmt.Errorf("customer %d not found in bank %d", customerID, bankID)
		}
	}

	return fmt.Errorf("bank with ID %d not found", bankID)
}

func (r *Repository) GetCustomers(bankID int64) ([]int64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, bank := range r.banks {
		if bank.ID == bankID {
			// Return a copy to avoid external modification
			customers := make([]int64, len(bank.Customers))
			copy(customers, bank.Customers)

			return customers, nil
		}
	}

	return nil, fmt.Errorf("bank with ID %d not found", bankID)
}

func (r *Repository) GetCustomerCount(bankID int64) (int, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	bank, err := r.GetByID(bankID)

	if err != nil {
		return 0, fmt.Errorf("bank with ID %d not found", bankID)
	}

	return len(bank.Customers), nil
}
