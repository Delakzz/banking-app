package bank

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Repository manages bank data storage and retrieval using JSON files
type Repository struct {
	filePath string
	mutex    sync.RWMutex
	nextID   int64
	banks    []*Bank // Cache for in-memory operations
}

// NewRepository creates a new bank repository with JSON file storage
func NewRepository(dataDir string) *Repository {
	// Ensure data directory exists
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		panic(fmt.Sprintf("failed to create data directory: %v", err))
	}

	filePath := filepath.Join(dataDir, "banks.json")

	repo := &Repository{
		filePath: filePath,
		nextID:   1,
		banks:    []*Bank{},
	}

	// Initialize with existing data if file exists
	repo.loadData()

	return repo
}

// loadData loads existing bank data from JSON file
func (r *Repository) loadData() {
	// Check if file exists
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		// File doesn't exist, start with empty data
		return
	}

	// Read and parse existing data
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

	// Find the highest ID to set nextID correctly
	for _, bank := range banks {
		if bank.ID >= r.nextID {
			r.nextID = bank.ID + 1
		}
	}

	r.banks = banks
}

// saveData saves bank data to JSON file
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

// Create adds a new bank to the repository
func (r *Repository) Create(username, password, name string) (*Bank, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Create new bank
	bank := NewBank(r.nextID, username, password, name)
	r.banks = append(r.banks, bank)
	r.nextID++

	// Save updated data
	if err := r.saveData(); err != nil {
		return nil, fmt.Errorf("failed to save bank data: %w", err)
	}

	return bank, nil
}

// GetByID retrieves a bank by its ID
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

// GetAll retrieves all banks
func (r *Repository) GetAll() []*Bank {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// Return a copy to avoid external modification
	banks := make([]*Bank, len(r.banks))
	copy(banks, r.banks)

	return banks
}

// Update modifies an existing bank
func (r *Repository) Update(id int64, username, password, name string) (*Bank, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Find and update bank
	for _, bank := range r.banks {
		if bank.ID == id {
			bank.Name = name
			bank.Username = username
			bank.Password = password

			// Save updated data
			if err := r.saveData(); err != nil {
				return nil, fmt.Errorf("failed to save bank data: %w", err)
			}

			return bank, nil
		}
	}

	return nil, fmt.Errorf("bank with ID %d not found", id)
}

// Delete removes a bank by ID
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

// AddCustomer adds a customer to a bank
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

// RemoveCustomer removes a customer from a bank
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

// GetCustomers returns all customer IDs for a bank
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
