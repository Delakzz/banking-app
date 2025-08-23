package bank

import (
	"fmt"
	"strconv"
)

// Handler provides CLI interface functions for bank operations
type Handler struct {
	service *Service
}

// NewHandler creates a new bank handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// HandleCreate processes bank creation from CLI input
func (h *Handler) HandleCreate(name, code string) {
	bank, err := h.service.CreateBank(name, code)
	if err != nil {
		fmt.Printf("Error creating bank: %v\n", err)
		return
	}
	
	fmt.Printf("Bank created successfully!\n")
	fmt.Printf("ID: %d, Name: %s, Code: %s\n", bank.ID, bank.Name, bank.Code)
}

// HandleGet processes bank retrieval by ID from CLI input
func (h *Handler) HandleGet(idStr string) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Invalid ID format: %s\n", idStr)
		return
	}
	
	bank, err := h.service.GetBank(id)
	if err != nil {
		fmt.Printf("Error retrieving bank: %v\n", err)
		return
	}
	
	fmt.Printf("Bank found:\n")
	fmt.Printf("ID: %d, Name: %s, Code: %s\n", bank.ID, bank.Name, bank.Code)
}

// HandleList displays all banks
func (h *Handler) HandleList() {
	banks := h.service.GetAllBanks()
	
	if len(banks) == 0 {
		fmt.Println("No banks found.")
		return
	}
	
	fmt.Printf("Found %d bank(s):\n", len(banks))
	fmt.Println("ID\tName\t\tCode")
	fmt.Println("--\t----\t\t----")
	
	for _, bank := range banks {
		fmt.Printf("%d\t%s\t\t%s\n", bank.ID, bank.Name, bank.Code)
	}
}

// HandleUpdate processes bank updates from CLI input
func (h *Handler) HandleUpdate(idStr, name, code string) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Invalid ID format: %s\n", idStr)
		return
	}
	
	bank, err := h.service.UpdateBank(id, name, code)
	if err != nil {
		fmt.Printf("Error updating bank: %v\n", err)
		return
	}
	
	fmt.Printf("Bank updated successfully!\n")
	fmt.Printf("ID: %d, Name: %s, Code: %s\n", bank.ID, bank.Name, bank.Code)
}

// HandleDelete processes bank deletion from CLI input
func (h *Handler) HandleDelete(idStr string) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Invalid ID format: %s\n", idStr)
		return
	}
	
	err = h.service.DeleteBank(id)
	if err != nil {
		fmt.Printf("Error deleting bank: %v\n", err)
		return
	}
	
	fmt.Printf("Bank with ID %d deleted successfully!\n", id)
}

// ShowHelp displays available commands
func (h *Handler) ShowHelp() {
	fmt.Println("Bank Management Commands:")
	fmt.Println("  create <name> <code>     - Create a new bank")
	fmt.Println("  get <id>                 - Get bank by ID")
	fmt.Println("  list                     - List all banks")
	fmt.Println("  update <id> <name> <code> - Update bank")
	fmt.Println("  delete <id>              - Delete bank")
	fmt.Println("  help                     - Show this help")
	fmt.Println("  exit                     - Exit the application")
}
