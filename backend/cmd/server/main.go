package main

import (
	"banking-app/backend/internal/bank"
	"bufio"
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("Welcome to Banking App!")
	fmt.Println("==========================")
	fmt.Println()
	fmt.Println("1. Create Bank")
	fmt.Println("2. Update Bank")
	fmt.Println("3. Delete Bank")
	fmt.Println("4. Get all Banks")
	fmt.Println()
	fmt.Println("==========================")
}

func main() {
	// Initialize repository with data directory
	dataDir := "../../db"
	bankRepo := bank.NewRepository(dataDir)

	// Initialize service and handler
	bankService := bank.NewService(bankRepo)
	bankHandler := bank.NewHandler(bankService)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()
		fmt.Print("\nChoice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "0":
			fmt.Println("Exiting the application. Goodbye!")
			return
		case "1":
			fmt.Print("Enter bank name: ")
			scanner.Scan()
			bankName := scanner.Text()
			bankHandler.HandleCreate(bankName)
		case "2":
			bankHandler.HandleList()
			fmt.Print("\nEnter Bank ID to update: ")
			scanner.Scan()
			bankID := scanner.Text()
			fmt.Print("Enter new Bank name: ")
			scanner.Scan()
			newBankName := scanner.Text()
			bankHandler.HandleUpdate(bankID, newBankName)
		case "3":
			bankHandler.HandleList()
			fmt.Print("\nEnter Bank ID to delete: ")
			scanner.Scan()
			bankID := scanner.Text()
			bankHandler.HandleDelete(bankID)
		case "4":
			bankHandler.HandleList()
			fmt.Print("\nPress enter to continue ")
			scanner.Scan()
		default:
			fmt.Println("Whyyy???")
		}
	}
}
