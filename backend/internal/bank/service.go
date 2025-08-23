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
func (s *Service) CreateBank(name, code string) (*Bank, error) {
	// Validate input
	if err := s.validateBankInput(name, code); err != nil {
		return nil, err
	}
	
	// Create bank through repository
	bank, err := s.repo.Create(name, code)
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
func (s *Service) UpdateBank(id int64, name, code string) (*Bank, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid bank ID: %d", id)
	}
	
	// Validate input
	if err := s.validateBankInput(name, code); err != nil {
		return nil, err
	}
	
	// Update bank through repository
	bank, err := s.repo.Update(id, name, code)
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

// validateBankInput validates bank name and code
func (s *Service) validateBankInput(name, code string) error {
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
	
	// Validate code
	if strings.TrimSpace(code) == "" {
		return fmt.Errorf("bank code cannot be empty")
	}
	if len(code) < 2 {
		return fmt.Errorf("bank code must be at least 2 characters long")
	}
	if len(code) > 10 {
		return fmt.Errorf("bank code cannot exceed 10 characters")
	}
	
	// Check for special characters in code (only allow alphanumeric)
	for _, char := range code {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return fmt.Errorf("bank code can only contain letters and numbers")
		}
	}
	
	return nil
}
