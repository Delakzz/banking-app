package bank

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HandleCreate(userID int64, name string) {
	bank, err := h.service.CreateBank(userID, name)
	if err != nil {
		fmt.Printf("Error creating bank: %v\n", err)
		return
	}

	fmt.Printf("Bank created successfully!\n")
	fmt.Printf("ID: %d, Name: %s\n", bank.ID, bank.Name)
}

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
		fmt.Printf("%d\t%s\n", bank.ID, bank.Name)
	}
}

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

func (h *Handler) NewBankLogin(userID int64) {
	scanner := bufio.NewReader(os.Stdin)
	bank, err := h.service.repo.GetBankByUserID(userID)
	if err != nil {
		fmt.Print("Enter bank name: ")
		bankName, _ := scanner.ReadString('\n')
		bankName = strings.TrimSpace(bankName) // Trim whitespace and newlines
		h.HandleCreate(userID, bankName)
	} else {
		fmt.Printf("Welcome back, %s!\n", bank.Name)
	}
}

// func (h *Handler) HandleAddCustomer(bankIDStr, customerIDStr string) {
// 	bankID, err := strconv.ParseInt(bankIDStr, 10, 64)
// 	if err != nil {
// 		fmt.Printf("Invalid bank ID format: %s\n", bankIDStr)
// 		return
// 	}

// 	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
// 	if err != nil {
// 		fmt.Printf("Invalid customer ID format: %s\n", customerIDStr)
// 		return
// 	}

// 	err = h.service.AddCustomer(bankID, customerID)
// 	if err != nil {
// 		fmt.Printf("Error adding customer to bank: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Customer %d successfully added to bank %d\n", customerID, bankID)
// }

// func (h *Handler) HandleRemoveCustomer(bankIDStr, customerIDStr string) {
// 	bankID, err := strconv.ParseInt(bankIDStr, 10, 64)
// 	if err != nil {
// 		fmt.Printf("Invalid bank ID format: %s\n", bankIDStr)
// 		return
// 	}

// 	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
// 	if err != nil {
// 		fmt.Printf("Invalid customer ID format: %s\n", customerIDStr)
// 		return
// 	}

// 	err = h.service.RemoveCustomer(bankID, customerID)
// 	if err != nil {
// 		fmt.Printf("Error removing customer from bank: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Customer %d successfully removed from bank %d\n", customerID, bankID)
// }

// func (h *Handler) HandleGetCustomers(bankIDStr string) {
// 	bankID, err := strconv.ParseInt(bankIDStr, 10, 64)
// 	if err != nil {
// 		fmt.Printf("Invalid bank ID format: %s\n", bankIDStr)
// 		return
// 	}

// 	customers, err := h.service.GetCustomers(bankID)
// 	if err != nil {
// 		fmt.Printf("Error getting bank customers: %v\n", err)
// 		return
// 	}

// 	if len(customers) == 0 {
// 		fmt.Printf("No customers found for bank %d\n", bankID)
// 		return
// 	}

// 	fmt.Printf("Customers for bank %d (Total: %d):\n", bankID, len(customers))
// 	for i, customerID := range customers {
// 		fmt.Printf("  %d. Customer ID: %d\n", i+1, customerID)
// 	}
// }
