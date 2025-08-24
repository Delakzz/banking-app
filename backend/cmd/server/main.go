package main

import (
	"banking-app/backend/internal/bank"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Initialize repository with data directory
	dataDir := "../../db"
	bankRepo := bank.NewRepository(dataDir)

	// Initialize service and handler
	bankService := bank.NewService(bankRepo)
	bankHandler := bank.NewHandler(bankService)

	fmt.Println("Welcome to Banking App!")
	fmt.Println("==========================")
	fmt.Println()
	fmt.Println("1. Create Bank")
	fmt.Println("2. Get all Banks")
	fmt.Println()
	fmt.Println("==========================")

	scanner := bufio.NewScanner(os.Stdin)

	for {
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
		}

	}
}
