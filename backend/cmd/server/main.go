package main

import (
	"banking-app/backend/internal/bank"
	"banking-app/backend/internal/user"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Database Structure:
// - Both banks and customers are stored in db/database.json
// - Banks collection: stores bank information with customer IDs
// - Customers collection: stores customer information with bank references
// - Banks can have multiple customers (stored as customer IDs)
// - Customers can only belong to one bank (stored as bank ID)

// func showUpdateMenu() {
// 	fmt.Println("What would you like to update?")
// 	fmt.Println("==========================")
// 	fmt.Println()
// 	fmt.Println("1. Username")
// 	fmt.Println("2. Password")
// 	fmt.Println("3. Name")
// 	fmt.Println()
// 	fmt.Println("==========================")
// }

// func showMenu() {
// 	fmt.Println("Welcome to Banking App!")
// 	fmt.Println("==========================")
// 	fmt.Println()
// 	fmt.Println("1. Create Bank")
// 	fmt.Println("2. Update Bank")
// 	fmt.Println("3. Delete Bank")
// 	fmt.Println("4. Get all Banks")
// 	fmt.Println()
// 	fmt.Println("==========================")
// }

// func main() {
// 	dataDir := "../../db"
// 	bankRepo := bank.NewRepository(dataDir)

// 	bankService := bank.NewService(bankRepo)
// 	bankHandler := bank.NewHandler(bankService)

// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {
// 		showMenu()
// 		fmt.Print("\nChoice: ")
// 		scanner.Scan()
// 		choice := scanner.Text()

// 		switch choice {
// 		case "0":
// 			fmt.Println("Exiting the application. Goodbye!")
// 			return
// 		case "1":
// 			fmt.Print("Enter bank username: ")
// 			scanner.Scan()
// 			bankUsername := scanner.Text()
// 			fmt.Print("Enter bank password: ")
// 			scanner.Scan()
// 			bankPassword := scanner.Text()
// 			fmt.Print("Enter bank name: ")
// 			scanner.Scan()
// 			bankName := scanner.Text()
// 			bankHandler.HandleCreate(bankUsername, bankPassword, bankName)
// 		case "2":
// 			bankHandler.HandleList()
// 			fmt.Print("\nEnter Bank ID to update: ")
// 			scanner.Scan()
// 			bankID := scanner.Text()
// 			fmt.Println("Just hit enter if you do not wish to update that variable.")
// 			fmt.Print("Enter new username: ")
// 			scanner.Scan()
// 			newUsername := scanner.Text()
// 			fmt.Print("Enter new password: ")
// 			scanner.Scan()
// 			newPassword := scanner.Text()
// 			fmt.Print("Enter new name: ")
// 			scanner.Scan()
// 			newName := scanner.Text()

// 			bankHandler.HandleUpdate(bankID, newUsername, newPassword, newName)

// 		case "3":
// 			bankHandler.HandleList()
// 			fmt.Print("\nEnter Bank ID to delete: ")
// 			scanner.Scan()
// 			bankID := scanner.Text()
// 			bankHandler.HandleDelete(bankID)
// 		case "4":
// 			bankHandler.HandleList()
// 			fmt.Print("\nPress enter to continue ")
// 			scanner.Scan()
// 		default:
// 			fmt.Println("Whyyy???")
// 		}
// 	}
// }

func showLoginMenu() {
	fmt.Println("\n======= Login Menu =======")
	fmt.Println()
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println()
	fmt.Println("==========================")
}

func main() {
	dataDir := "../../db/database.json"
	userRepo, err := user.NewRepository(dataDir)
	if err != nil {
		panic(fmt.Errorf("failed to load database: %w", err))
	}

	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	bankRepo := bank.NewRepository(dataDir)

	bankService := bank.NewService(bankRepo)
	bankHandler := bank.NewHandler(bankService)

	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Banking App!")

	for {
		showLoginMenu()

		fmt.Print("Choose: ")
		choice, _ := scanner.ReadString('\n')
		choice = strings.TrimSpace(choice) // important: remove \n

		switch choice {
		case "0":
			fmt.Println("\n👋 Exiting the application. Goodbye!")
			return

		case "1":
			fmt.Println("\n🔓 Login")
			u := userHandler.Login()
			if u == nil {
				fmt.Println("❌ Login failed. Please try again.")
				continue
			}

			switch u.Role {
			case user.RoleBank:
				bankHandler.NewBankLogin(u.ID)
			case user.RoleCustomer:
				fmt.Println("🙋 You are logged in as a Customer!")
				// TODO: call customer menu handler here
			default:
				fmt.Println("⚠️ Unknown role. Please contact admin.")
			}

		case "2":
			fmt.Println("\n🔑 Register New User")
			userHandler.Register()

		default:
			fmt.Println("❌ Invalid choice. Please select a valid option.")
		}
	}
}
