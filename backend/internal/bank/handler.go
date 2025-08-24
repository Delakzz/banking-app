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
func (h *Handler) HandleCreate(name string) {
	bank, err := h.service.CreateBank(name)
	if err != nil {
		fmt.Printf("Error creating bank: %v\n", err)
		return
	}

	fmt.Printf("Bank created successfully!\n")
	fmt.Printf("ID: %d, Name: %s\n", bank.ID, bank.Name)
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
	fmt.Printf("ID: %d, Name: %s\n", bank.ID, bank.Name)
}

// HandleList displays all banks
func (h *Handler) HandleList() {
	banks := h.service.GetAllBanks()

	if len(banks) == 0 {
		fmt.Println("No banks found.")
		return
	}

	fmt.Printf("Found %d bank(s):\n", len(banks))
	fmt.Println("ID\tName")
	fmt.Println("--\t---------")

	for _, bank := range banks {
		// customerCount := len(bank.Customers)
		fmt.Printf("%d\t%s\n", bank.ID, bank.Name)
	}
}

// HandleUpdate processes bank updates from CLI input
func (h *Handler) HandleUpdate(idStr, name string) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Invalid ID format: %s\n", idStr)
		return
	}

	bank, err := h.service.UpdateBank(id, name)
	if err != nil {
		fmt.Printf("Error updating bank: %v\n", err)
		return
	}

	fmt.Printf("Bank updated successfully!\n")
	fmt.Printf("ID: %d, Name: %s\n", bank.ID, bank.Name)
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
// func (h *Handler) ShowHelp() {
// 	fmt.Println("Bank Management Commands:")
// 	fmt.Println("  create <name>            - Create a new bank")
// 	fmt.Println("  get <id>                 - Get bank by ID")
// 	fmt.Println("  list                     - List all banks")
// 	fmt.Println("  update <id> <name>       - Update bank")
// 	fmt.Println("  delete <id>              - Delete bank")
// 	fmt.Println("  add-customer <bank_id> <customer_id> - Add customer to bank")
// 	fmt.Println("  remove-customer <bank_id> <customer_id> - Remove customer from bank")
// 	fmt.Println("  get-customers <bank_id>  - List customers for a bank")
// 	fmt.Println("  help                     - Show this help")
// 	fmt.Println("  exit                     - Exit the application")
// }

// HandleAddCustomer processes adding a customer to a bank
func (h *Handler) HandleAddCustomer(bankIDStr, customerIDStr string) {
	bankID, err := strconv.ParseInt(bankIDStr, 10, 64)
	if err != nil {
		fmt.Printf("Invalid bank ID format: %s\n", bankIDStr)
		return
	}

	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		fmt.Printf("Invalid customer ID format: %s\n", customerIDStr)
		return
	}

	err = h.service.AddCustomer(bankID, customerID)
	if err != nil {
		fmt.Printf("Error adding customer to bank: %v\n", err)
		return
	}

	fmt.Printf("Customer %d successfully added to bank %d\n", customerID, bankID)
}

// HandleRemoveCustomer processes removing a customer from a bank
func (h *Handler) HandleRemoveCustomer(bankIDStr, customerIDStr string) {
	bankID, err := strconv.ParseInt(bankIDStr, 10, 64)
	if err != nil {
		fmt.Printf("Invalid bank ID format: %s\n", bankIDStr)
		return
	}

	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		fmt.Printf("Invalid customer ID format: %s\n", customerIDStr)
		return
	}

	err = h.service.RemoveCustomer(bankID, customerID)
	if err != nil {
		fmt.Printf("Error removing customer from bank: %v\n", err)
		return
	}

	fmt.Printf("Customer %d successfully removed from bank %d\n", customerID, bankID)
}

// HandleGetCustomers displays all customers for a specific bank
func (h *Handler) HandleGetCustomers(bankIDStr string) {
	bankID, err := strconv.ParseInt(bankIDStr, 10, 64)
	if err != nil {
		fmt.Printf("Invalid bank ID format: %s\n", bankIDStr)
		return
	}

	customers, err := h.service.GetCustomers(bankID)
	if err != nil {
		fmt.Printf("Error getting bank customers: %v\n", err)
		return
	}

	if len(customers) == 0 {
		fmt.Printf("No customers found for bank %d\n", bankID)
		return
	}

	fmt.Printf("Customers for bank %d:\n", bankID)
	for i, customerID := range customers {
		fmt.Printf("  %d. Customer ID: %d\n", i+1, customerID)
	}
}
