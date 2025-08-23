package bank

import (
	"fmt"
	"strings"
)

// Service handles business logic for bank operations
type Service struct {
	repo *Repository
}

// NewService creates a new bank service
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// CreateBank creates a new bank with validation
func (s *Service) CreateBank(name string) (*Bank, error) {
	// Validate input
	if err := s.validateBankInput(name); err != nil {
		return nil, err
	}
	
	// Create bank through repository
	bank, err := s.repo.Create(name)
	if err != nil {
		return nil, fmt.Errorf("failed to create bank: %w", err)
	}
	
	return bank, nil
}

// GetBank retrieves a bank by ID
func (s *Service) GetBank(id int64) (*Bank, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid bank ID: %d", id)
	}
	
	bank, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get bank: %w", err)
	}
	
	return bank, nil
}

// GetAllBanks retrieves all banks
func (s *Service) GetAllBanks() []*Bank {
	return s.repo.GetAll()
}

// UpdateBank updates an existing bank
func (s *Service) UpdateBank(id int64, name string) (*Bank, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid bank ID: %d", id)
	}
	
	// Validate input
	if err := s.validateBankInput(name); err != nil {
		return nil, err
	}
	
	// Update bank through repository
	bank, err := s.repo.Update(id, name)
	if err != nil {
		return nil, fmt.Errorf("failed to update bank: %w", err)
	}
	
	return bank, nil
}

// DeleteBank removes a bank by ID
func (s *Service) DeleteBank(id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid bank ID: %d", id)
	}
	
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete bank: %w", err)
	}
	
	return nil
}

// validateBankInput validates bank name
func (s *Service) validateBankInput(name string) error {
	// Validate name
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("bank name cannot be empty")
	}
	if len(name) < 2 {
		return fmt.Errorf("bank name must be at least 2 characters long")
	}
	if len(name) > 50 {
		return fmt.Errorf("bank name cannot exceed 50 characters")
	}
	
	return nil
}

// AddCustomer adds a customer to a bank
func (s *Service) AddCustomer(bankID, customerID int64) error {
	if bankID <= 0 {
		return fmt.Errorf("invalid bank ID: %d", bankID)
	}
	if customerID <= 0 {
		return fmt.Errorf("invalid customer ID: %d", customerID)
	}
	
	err := s.repo.AddCustomer(bankID, customerID)
	if err != nil {
		return fmt.Errorf("failed to add customer to bank: %w", err)
	}
	
	return nil
}

// RemoveCustomer removes a customer from a bank
func (s *Service) RemoveCustomer(bankID, customerID int64) error {
	if bankID <= 0 {
		return fmt.Errorf("invalid bank ID: %d", bankID)
	}
	if customerID <= 0 {
		return fmt.Errorf("invalid customer ID: %d", customerID)
	}
	
	err := s.repo.RemoveCustomer(bankID, customerID)
	if err != nil {
		return fmt.Errorf("failed to remove customer from bank: %w", err)
	}
	
	return nil
}

// GetCustomers retrieves all customer IDs for a bank
func (s *Service) GetCustomers(bankID int64) ([]int64, error) {
	if bankID <= 0 {
		return nil, fmt.Errorf("invalid bank ID: %d", bankID)
	}
	
	customers, err := s.repo.GetCustomers(bankID)
	if err != nil {
		return nil, fmt.Errorf("failed to get bank customers: %w", err)
	}
	
	return customers, nil
}
